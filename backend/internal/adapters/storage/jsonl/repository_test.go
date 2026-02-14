package jsonl

import (
	"context"
	"path/filepath"
	"testing"
	"time"

	"github.com/diteria/project_llot/backend/internal/domain/traffic"
)

func TestSaveAndQueryEvents(t *testing.T) {
	t.Parallel()

	tempDir := t.TempDir()
	repo, err := NewRepository(filepath.Join(tempDir, "events.jsonl"))
	if err != nil {
		t.Fatalf("new repository: %v", err)
	}

	events := []traffic.RequestEvent{
		{ID: "1", Host: "a.local", OccurredAt: mustTime("2026-02-14T10:00:00Z")},
		{ID: "2", Host: "b.local", OccurredAt: mustTime("2026-02-14T10:05:00Z")},
	}
	if err := repo.SaveEvents(context.Background(), events); err != nil {
		t.Fatalf("save events: %v", err)
	}

	result, err := repo.QueryEvents(context.Background(), mustTime("2026-02-14T09:59:00Z"), mustTime("2026-02-14T10:10:00Z"), "a.local")
	if err != nil {
		t.Fatalf("query events: %v", err)
	}
	if len(result) != 1 {
		t.Fatalf("expected 1 event, got %d", len(result))
	}
	if result[0].ID != "1" {
		t.Fatalf("unexpected event ID: %s", result[0].ID)
	}
}

func mustTime(raw string) time.Time {
	parsed, err := time.Parse(time.RFC3339, raw)
	if err != nil {
		panic(err)
	}
	return parsed.UTC()
}
