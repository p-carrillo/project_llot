package ports

import (
	"context"
	"time"

	"github.com/diteria/project_llot/backend/internal/domain/traffic"
)

type NginxStructuredLogParser interface {
	ParseLine(line string) (traffic.ParsedLog, error)
}

type TrafficRepository interface {
	SaveEvents(ctx context.Context, events []traffic.RequestEvent) error
	QueryEvents(ctx context.Context, from, to time.Time, host string) ([]traffic.RequestEvent, error)
}
