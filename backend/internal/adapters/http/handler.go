package http

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	apphealth "github.com/diteria/project_llot/backend/internal/application/health"
	apptraffic "github.com/diteria/project_llot/backend/internal/application/traffic"
)

type Handler struct {
	health       apphealth.Service
	traffic      apptraffic.Service
	maxBodyBytes int64
}

type ingestRequest struct {
	Lines []string `json:"lines"`
}

type errorEnvelope struct {
	Error errorBody `json:"error"`
}

type errorBody struct {
	Code      string `json:"code"`
	Message   string `json:"message"`
	RequestID string `json:"request_id"`
}

func NewHandler(health apphealth.Service, traffic apptraffic.Service, maxBodyBytes int64) Handler {
	return Handler{health: health, traffic: traffic, maxBodyBytes: maxBodyBytes}
}

func (h Handler) Routes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/health/live", h.live)
	mux.HandleFunc("/health/ready", h.ready)
	mux.HandleFunc("/api/v1/health", h.ready)
	mux.HandleFunc("/api/v1/ingest/logs", h.ingestLogs)
	mux.HandleFunc("/api/v1/metrics/overview", h.metricsOverview)
	mux.HandleFunc("/api/v1/metrics/windows", h.metricsWindows)
	return mux
}

func (h Handler) live(w http.ResponseWriter, _ *http.Request) {
	writeJSON(w, http.StatusOK, h.health.Live())
}

func (h Handler) ready(w http.ResponseWriter, r *http.Request) {
	snapshot := h.health.Ready(r.Context())
	status := http.StatusOK
	if snapshot.Status != "ok" {
		status = http.StatusServiceUnavailable
	}
	writeJSON(w, status, snapshot)
}

func (h Handler) ingestLogs(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		h.writeError(w, r, http.StatusMethodNotAllowed, "method_not_allowed", "method not allowed")
		return
	}

	body := io.LimitReader(r.Body, h.maxBodyBytes)
	defer r.Body.Close()

	var req ingestRequest
	if err := json.NewDecoder(body).Decode(&req); err != nil {
		h.writeError(w, r, http.StatusBadRequest, "invalid_json", "request body must be valid JSON")
		return
	}
	if len(req.Lines) == 0 {
		h.writeError(w, r, http.StatusBadRequest, "invalid_payload", "lines must not be empty")
		return
	}
	if len(req.Lines) > 10000 {
		h.writeError(w, r, http.StatusBadRequest, "payload_too_large", "maximum 10000 lines per request")
		return
	}

	result, err := h.traffic.IngestLines(r.Context(), req.Lines)
	if err != nil {
		h.writeError(w, r, http.StatusInternalServerError, "ingest_failed", "failed to ingest lines")
		return
	}
	writeJSON(w, http.StatusAccepted, result)
}

func (h Handler) metricsOverview(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		h.writeError(w, r, http.StatusMethodNotAllowed, "method_not_allowed", "method not allowed")
		return
	}

	from, to, err := parseTimeRange(r)
	if err != nil {
		h.writeError(w, r, http.StatusBadRequest, "invalid_time_range", err.Error())
		return
	}

	query := apptraffic.OverviewQuery{From: from, To: to, Host: r.URL.Query().Get("host")}
	result, err := h.traffic.Overview(r.Context(), query)
	if err != nil {
		if errors.Is(err, apptraffic.ErrInvalidTimeRange) {
			h.writeError(w, r, http.StatusBadRequest, "invalid_time_range", err.Error())
			return
		}
		h.writeError(w, r, http.StatusInternalServerError, "overview_failed", "failed to build overview")
		return
	}
	writeJSON(w, http.StatusOK, result)
}

func (h Handler) metricsWindows(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		h.writeError(w, r, http.StatusMethodNotAllowed, "method_not_allowed", "method not allowed")
		return
	}

	from, to, err := parseTimeRange(r)
	if err != nil {
		h.writeError(w, r, http.StatusBadRequest, "invalid_time_range", err.Error())
		return
	}

	step := time.Minute
	if raw := r.URL.Query().Get("step"); raw != "" {
		parsed, parseErr := time.ParseDuration(raw)
		if parseErr != nil {
			h.writeError(w, r, http.StatusBadRequest, "invalid_step", "step must be a valid duration (e.g. 1m, 5m)")
			return
		}
		step = parsed
	}

	limit := 100
	if raw := r.URL.Query().Get("limit"); raw != "" {
		parsed, parseErr := strconv.Atoi(raw)
		if parseErr != nil || parsed <= 0 || parsed > 1000 {
			h.writeError(w, r, http.StatusBadRequest, "invalid_limit", "limit must be between 1 and 1000")
			return
		}
		limit = parsed
	}

	cursor := 0
	if raw := r.URL.Query().Get("cursor"); raw != "" {
		parsed, parseErr := strconv.Atoi(raw)
		if parseErr != nil || parsed < 0 {
			h.writeError(w, r, http.StatusBadRequest, "invalid_cursor", "cursor must be a non-negative integer")
			return
		}
		cursor = parsed
	}

	query := apptraffic.WindowsQuery{
		From:   from,
		To:     to,
		Host:   r.URL.Query().Get("host"),
		Step:   step,
		Cursor: cursor,
		Limit:  limit,
	}

	result, err := h.traffic.Windows(r.Context(), query)
	if err != nil {
		switch {
		case errors.Is(err, apptraffic.ErrInvalidTimeRange):
			h.writeError(w, r, http.StatusBadRequest, "invalid_time_range", err.Error())
		case errors.Is(err, apptraffic.ErrInvalidStep):
			h.writeError(w, r, http.StatusBadRequest, "invalid_step", err.Error())
		default:
			h.writeError(w, r, http.StatusInternalServerError, "windows_failed", "failed to build window metrics")
		}
		return
	}

	writeJSON(w, http.StatusOK, result)
}

func parseTimeRange(r *http.Request) (time.Time, time.Time, error) {
	fromRaw := r.URL.Query().Get("from")
	toRaw := r.URL.Query().Get("to")
	if fromRaw == "" || toRaw == "" {
		return time.Time{}, time.Time{}, errors.New("from and to query params are required")
	}

	from, err := time.Parse(time.RFC3339, fromRaw)
	if err != nil {
		return time.Time{}, time.Time{}, fmt.Errorf("invalid from timestamp: %w", err)
	}
	to, err := time.Parse(time.RFC3339, toRaw)
	if err != nil {
		return time.Time{}, time.Time{}, fmt.Errorf("invalid to timestamp: %w", err)
	}
	return from.UTC(), to.UTC(), nil
}

func (h Handler) writeError(w http.ResponseWriter, r *http.Request, status int, code string, message string) {
	requestID := r.Header.Get("X-Request-ID")
	if requestID == "" {
		requestID = "local"
	}
	writeJSON(w, status, errorEnvelope{
		Error: errorBody{Code: code, Message: message, RequestID: requestID},
	})
}

func writeJSON(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(payload)
}
