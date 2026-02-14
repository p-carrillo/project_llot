package jsonl

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"sync"
	"time"

	"github.com/diteria/project_llot/backend/internal/domain/traffic"
)

type Repository struct {
	mu       sync.RWMutex
	filePath string
}

func NewRepository(filePath string) (*Repository, error) {
	if filePath == "" {
		return nil, fmt.Errorf("file path is required")
	}

	dir := filepath.Dir(filePath)
	if err := os.MkdirAll(dir, 0o750); err != nil {
		return nil, fmt.Errorf("create data directory: %w", err)
	}

	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o640)
	if err != nil {
		return nil, fmt.Errorf("open data file: %w", err)
	}
	if err := file.Close(); err != nil {
		return nil, fmt.Errorf("close data file: %w", err)
	}

	return &Repository{filePath: filePath}, nil
}

func (r *Repository) SaveEvents(ctx context.Context, events []traffic.RequestEvent) error {
	if len(events) == 0 {
		return nil
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	file, err := os.OpenFile(r.filePath, os.O_APPEND|os.O_WRONLY, 0o640)
	if err != nil {
		return fmt.Errorf("open data file for append: %w", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	for _, event := range events {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}
		if err := encoder.Encode(event); err != nil {
			return fmt.Errorf("encode event: %w", err)
		}
	}

	return nil
}

func (r *Repository) QueryEvents(ctx context.Context, from, to time.Time, host string) ([]traffic.RequestEvent, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	file, err := os.OpenFile(r.filePath, os.O_RDONLY, 0o640)
	if err != nil {
		return nil, fmt.Errorf("open data file for read: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Buffer(make([]byte, 0, 64*1024), 2*1024*1024)

	filtered := make([]traffic.RequestEvent, 0)
	line := 0
	for scanner.Scan() {
		line++
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}

		var event traffic.RequestEvent
		if err := json.Unmarshal(scanner.Bytes(), &event); err != nil {
			return nil, fmt.Errorf("decode event line %d: %w", line, err)
		}
		if event.OccurredAt.Before(from) || !event.OccurredAt.Before(to) {
			continue
		}
		if host != "" && event.Host != host {
			continue
		}
		filtered = append(filtered, event)
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("scan data file: %w", err)
	}

	sort.Slice(filtered, func(i, j int) bool {
		return filtered[i].OccurredAt.Before(filtered[j].OccurredAt)
	})

	return filtered, nil
}
