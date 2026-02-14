package nginxjson

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/diteria/project_llot/backend/internal/domain/traffic"
)

type Parser struct{}

func NewParser() Parser {
	return Parser{}
}

func (Parser) ParseLine(line string) (traffic.ParsedLog, error) {
	var payload map[string]any
	if err := json.Unmarshal([]byte(line), &payload); err != nil {
		return traffic.ParsedLog{}, fmt.Errorf("invalid json: %w", err)
	}

	occurredAt, err := parseTime(payload)
	if err != nil {
		return traffic.ParsedLog{}, err
	}

	host := firstString(payload, "host", "server_name")
	if host == "" {
		return traffic.ParsedLog{}, errors.New("missing host")
	}

	path := firstString(payload, "uri", "request_uri", "path")
	if path == "" {
		path = "/"
	}

	method := strings.ToUpper(firstString(payload, "request_method", "method"))
	if method == "" {
		method = "GET"
	}

	status, err := parseInt(payload, "status")
	if err != nil {
		return traffic.ParsedLog{}, errors.New("missing or invalid status")
	}

	remoteAddr := firstString(payload, "remote_addr", "client_ip")
	if remoteAddr == "" {
		remoteAddr = "unknown"
	}

	userAgent := firstString(payload, "http_user_agent", "user_agent")

	return traffic.ParsedLog{
		OccurredAt: occurredAt,
		Host:       host,
		Path:       path,
		Method:     method,
		StatusCode: status,
		RemoteAddr: remoteAddr,
		UserAgent:  userAgent,
	}, nil
}

func parseTime(payload map[string]any) (time.Time, error) {
	raw := firstString(payload, "time_iso8601", "time", "ts", "timestamp")
	if raw == "" {
		return time.Time{}, errors.New("missing timestamp")
	}

	layouts := []string{
		time.RFC3339Nano,
		time.RFC3339,
		"02/Jan/2006:15:04:05 -0700",
	}

	for _, layout := range layouts {
		parsed, err := time.Parse(layout, raw)
		if err == nil {
			return parsed.UTC(), nil
		}
	}
	return time.Time{}, fmt.Errorf("invalid timestamp format: %s", raw)
}

func firstString(payload map[string]any, keys ...string) string {
	for _, key := range keys {
		value, ok := payload[key]
		if !ok || value == nil {
			continue
		}
		switch typed := value.(type) {
		case string:
			trimmed := strings.TrimSpace(typed)
			if trimmed != "" {
				return trimmed
			}
		case float64:
			return fmt.Sprintf("%.0f", typed)
		}
	}
	return ""
}

func parseInt(payload map[string]any, key string) (int, error) {
	value, ok := payload[key]
	if !ok || value == nil {
		return 0, errors.New("missing key")
	}

	switch typed := value.(type) {
	case float64:
		return int(typed), nil
	case int:
		return typed, nil
	case string:
		var parsed int
		_, err := fmt.Sscanf(strings.TrimSpace(typed), "%d", &parsed)
		if err != nil {
			return 0, err
		}
		return parsed, nil
	default:
		return 0, errors.New("unsupported type")
	}
}
