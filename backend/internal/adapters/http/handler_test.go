package http

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/diteria/project_llot/backend/internal/adapters/health"
	"github.com/diteria/project_llot/backend/internal/adapters/ingest/nginxjson"
	"github.com/diteria/project_llot/backend/internal/adapters/storage/memory"
	apphealth "github.com/diteria/project_llot/backend/internal/application/health"
	apptraffic "github.com/diteria/project_llot/backend/internal/application/traffic"
)

func TestIngestAndOverview(t *testing.T) {
	t.Parallel()

	checker := health.NewStaticChecker(true)
	healthService := apphealth.NewService(checker)
	parser := nginxjson.NewParser()
	repo := memory.NewRepository()
	trafficService := apptraffic.NewService(parser, repo, 30*time.Minute)

	h := NewHandler(healthService, trafficService, 1024*1024)
	routes := h.Routes()

	ingestBody := map[string]any{
		"lines": []string{
			`{"time_iso8601":"2026-02-14T18:00:00Z","host":"site.local","request_method":"GET","uri":"/","status":200,"remote_addr":"1.2.3.4","http_user_agent":"Mozilla/5.0"}`,
			`{"time_iso8601":"2026-02-14T18:01:00Z","host":"site.local","request_method":"GET","uri":"/pricing","status":200,"remote_addr":"5.6.7.8","http_user_agent":"Googlebot"}`,
		},
	}
	payload, _ := json.Marshal(ingestBody)

	ingestReq := httptest.NewRequest(http.MethodPost, "/api/v1/ingest/logs", bytes.NewReader(payload))
	ingestReq.Header.Set("Content-Type", "application/json")
	ingestRec := httptest.NewRecorder()
	routes.ServeHTTP(ingestRec, ingestReq)
	if ingestRec.Code != http.StatusAccepted {
		t.Fatalf("expected ingest status %d, got %d", http.StatusAccepted, ingestRec.Code)
	}

	overviewReq := httptest.NewRequest(
		http.MethodGet,
		"/api/v1/metrics/overview?from=2026-02-14T17:00:00Z&to=2026-02-14T19:00:00Z",
		nil,
	)
	overviewRec := httptest.NewRecorder()
	routes.ServeHTTP(overviewRec, overviewReq)
	if overviewRec.Code != http.StatusOK {
		t.Fatalf("expected overview status %d, got %d", http.StatusOK, overviewRec.Code)
	}

	var body struct {
		Requests int `json:"requests"`
		Human    int `json:"human"`
		Bot      int `json:"bot"`
	}
	if err := json.Unmarshal(overviewRec.Body.Bytes(), &body); err != nil {
		t.Fatalf("decode overview body: %v", err)
	}
	if body.Requests != 2 || body.Human != 1 || body.Bot != 1 {
		t.Fatalf("unexpected overview body: %+v", body)
	}
}

func TestIngestRejectsEmptyLines(t *testing.T) {
	t.Parallel()

	checker := health.NewStaticChecker(true)
	healthService := apphealth.NewService(checker)
	parser := nginxjson.NewParser()
	repo := memory.NewRepository()
	trafficService := apptraffic.NewService(parser, repo, 30*time.Minute)

	h := NewHandler(healthService, trafficService, 1024*1024)
	routes := h.Routes()

	payload := []byte(`{"lines":[]}`)
	req := httptest.NewRequest(http.MethodPost, "/api/v1/ingest/logs", bytes.NewReader(payload))
	rec := httptest.NewRecorder()
	routes.ServeHTTP(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("expected status %d, got %d", http.StatusBadRequest, rec.Code)
	}
}
