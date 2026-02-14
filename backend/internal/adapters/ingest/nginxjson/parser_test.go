package nginxjson

import (
	"testing"
	"time"
)

func TestParseLineOK(t *testing.T) {
	line := `{"time_iso8601":"2026-02-14T18:00:00Z","host":"site.local","request_method":"GET","uri":"/pricing","status":200,"remote_addr":"1.2.3.4","http_user_agent":"Mozilla/5.0"}`
	parser := NewParser()

	parsed, err := parser.ParseLine(line)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if parsed.Host != "site.local" {
		t.Fatalf("unexpected host: %s", parsed.Host)
	}
	if parsed.Path != "/pricing" {
		t.Fatalf("unexpected path: %s", parsed.Path)
	}
	if parsed.StatusCode != 200 {
		t.Fatalf("unexpected status code: %d", parsed.StatusCode)
	}
	if parsed.OccurredAt.Format(time.RFC3339) != "2026-02-14T18:00:00Z" {
		t.Fatalf("unexpected timestamp: %s", parsed.OccurredAt.Format(time.RFC3339))
	}
}

func TestParseLineInvalidJSON(t *testing.T) {
	parser := NewParser()
	_, err := parser.ParseLine("not-json")
	if err == nil {
		t.Fatalf("expected error for invalid json")
	}
}
