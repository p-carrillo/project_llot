package traffic

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/diteria/project_llot/backend/internal/domain/traffic"
)

type fakeParser struct{}

func (fakeParser) ParseLine(line string) (traffic.ParsedLog, error) {
	if line == "bad" {
		return traffic.ParsedLog{}, fmt.Errorf("invalid")
	}
	return traffic.ParsedLog{
		OccurredAt: mustTime("2026-02-14T18:00:00Z"),
		Host:       "site.local",
		Path:       "/",
		Method:     "GET",
		StatusCode: 200,
		RemoteAddr: "1.2.3.4",
		UserAgent:  line,
	}, nil
}

type fakeRepo struct {
	events []traffic.RequestEvent
}

func (r *fakeRepo) SaveEvents(_ context.Context, events []traffic.RequestEvent) error {
	r.events = append(r.events, events...)
	return nil
}

func (r *fakeRepo) QueryEvents(_ context.Context, from, to time.Time, host string) ([]traffic.RequestEvent, error) {
	filtered := make([]traffic.RequestEvent, 0)
	for _, event := range r.events {
		if event.OccurredAt.Before(from) || !event.OccurredAt.Before(to) {
			continue
		}
		if host != "" && event.Host != host {
			continue
		}
		filtered = append(filtered, event)
	}
	return filtered, nil
}

func TestIngestOverviewAndWindows(t *testing.T) {
	repo := &fakeRepo{}
	svc := NewService(fakeParser{}, repo, 30*time.Minute)

	result, err := svc.IngestLines(context.Background(), []string{"Mozilla/5.0", "Googlebot", "bad"})
	if err != nil {
		t.Fatalf("unexpected ingest error: %v", err)
	}
	if result.Accepted != 2 || result.Rejected != 1 {
		t.Fatalf("unexpected ingest result: %+v", result)
	}

	overview, err := svc.Overview(context.Background(), OverviewQuery{
		From: mustTime("2026-02-14T17:59:00Z"),
		To:   mustTime("2026-02-14T18:30:00Z"),
	})
	if err != nil {
		t.Fatalf("unexpected overview error: %v", err)
	}
	if overview.Requests != 2 || overview.Human != 1 || overview.Bot != 1 {
		t.Fatalf("unexpected overview: %+v", overview)
	}

	windows, err := svc.Windows(context.Background(), WindowsQuery{
		From:  mustTime("2026-02-14T17:59:00Z"),
		To:    mustTime("2026-02-14T18:30:00Z"),
		Step:  time.Minute,
		Limit: 100,
	})
	if err != nil {
		t.Fatalf("unexpected windows error: %v", err)
	}
	if len(windows.Items) != 1 {
		t.Fatalf("expected one window, got %d", len(windows.Items))
	}
	if windows.Items[0].Requests != 2 {
		t.Fatalf("unexpected window metric: %+v", windows.Items[0])
	}
}

func mustTime(raw string) time.Time {
	parsed, err := time.Parse(time.RFC3339, raw)
	if err != nil {
		panic(err)
	}
	return parsed.UTC()
}
