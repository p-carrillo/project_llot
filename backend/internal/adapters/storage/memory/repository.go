package memory

import (
	"context"
	"sort"
	"sync"
	"time"

	"github.com/diteria/project_llot/backend/internal/domain/traffic"
)

type Repository struct {
	mu     sync.RWMutex
	events []traffic.RequestEvent
}

func NewRepository() *Repository {
	return &Repository{events: make([]traffic.RequestEvent, 0, 1024)}
}

func (r *Repository) SaveEvents(_ context.Context, events []traffic.RequestEvent) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.events = append(r.events, events...)
	return nil
}

func (r *Repository) QueryEvents(_ context.Context, from, to time.Time, host string) ([]traffic.RequestEvent, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

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

	sort.Slice(filtered, func(i, j int) bool {
		return filtered[i].OccurredAt.Before(filtered[j].OccurredAt)
	})

	return filtered, nil
}
