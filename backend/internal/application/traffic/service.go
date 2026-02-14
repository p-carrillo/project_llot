package traffic

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/diteria/project_llot/backend/internal/domain/traffic"
	"github.com/diteria/project_llot/backend/internal/ports"
)

type Service struct {
	parser   ports.NginxStructuredLogParser
	repo     ports.TrafficRepository
	sessions *SessionEstimator
}

type IngestResult struct {
	Received      int      `json:"received"`
	Accepted      int      `json:"accepted"`
	Rejected      int      `json:"rejected"`
	RejectedLines []string `json:"rejected_lines,omitempty"`
}

type OverviewQuery struct {
	From time.Time
	To   time.Time
	Host string
}

type WindowsQuery struct {
	From   time.Time
	To     time.Time
	Host   string
	Step   time.Duration
	Cursor int
	Limit  int
}

type WindowsResult struct {
	Items      []traffic.WindowMetric `json:"items"`
	NextCursor *string                `json:"next_cursor,omitempty"`
}

func NewService(
	parser ports.NginxStructuredLogParser,
	repo ports.TrafficRepository,
	sessionTimeout time.Duration,
) Service {
	return Service{
		parser:   parser,
		repo:     repo,
		sessions: NewSessionEstimator(sessionTimeout),
	}
}

func (s Service) IngestLines(ctx context.Context, lines []string) (IngestResult, error) {
	result := IngestResult{Received: len(lines)}
	events := make([]traffic.RequestEvent, 0, len(lines))

	for idx, line := range lines {
		trimmed := strings.TrimSpace(line)
		if trimmed == "" {
			result.Rejected++
			result.RejectedLines = append(result.RejectedLines, fmt.Sprintf("line %d: empty line", idx+1))
			continue
		}

		parsed, err := s.parser.ParseLine(trimmed)
		if err != nil {
			result.Rejected++
			if len(result.RejectedLines) < 10 {
				result.RejectedLines = append(result.RejectedLines, fmt.Sprintf("line %d: %v", idx+1, err))
			}
			continue
		}

		classification, score := classifyUserAgent(parsed.UserAgent)
		remoteHash := hashString(parsed.RemoteAddr)
		visitorKey := fmt.Sprintf("%s|%s|%s", parsed.Host, remoteHash, strings.ToLower(parsed.UserAgent))
		sessionID := s.sessions.Estimate(visitorKey, parsed.OccurredAt)

		event := traffic.RequestEvent{
			ID:             eventID(parsed, sessionID),
			OccurredAt:     parsed.OccurredAt,
			Host:           parsed.Host,
			Path:           parsed.Path,
			Method:         parsed.Method,
			StatusCode:     parsed.StatusCode,
			RemoteAddrHash: remoteHash,
			UserAgent:      parsed.UserAgent,
			BotScore:       score,
			Class:          classification,
			SessionID:      sessionID,
		}
		events = append(events, event)
	}

	result.Accepted = len(events)
	if len(events) == 0 {
		return result, nil
	}

	if err := s.repo.SaveEvents(ctx, events); err != nil {
		return result, err
	}
	return result, nil
}

func (s Service) Overview(ctx context.Context, query OverviewQuery) (traffic.Overview, error) {
	if !query.From.Before(query.To) {
		return traffic.Overview{}, ErrInvalidTimeRange
	}
	events, err := s.repo.QueryEvents(ctx, query.From, query.To, query.Host)
	if err != nil {
		return traffic.Overview{}, err
	}

	hosts := make(map[string]struct{})
	sessions := make(map[string]struct{})
	result := traffic.Overview{From: query.From, To: query.To, Host: query.Host}
	for _, event := range events {
		result.Requests++
		hosts[event.Host] = struct{}{}
		sessions[event.SessionID] = struct{}{}
		switch event.Class {
		case traffic.ClassificationBot:
			result.Bot++
		case traffic.ClassificationHuman:
			result.Human++
		default:
			result.Unknown++
		}
	}
	result.UniqueHosts = len(hosts)
	result.Sessions = len(sessions)
	return result, nil
}

func (s Service) Windows(ctx context.Context, query WindowsQuery) (WindowsResult, error) {
	if !query.From.Before(query.To) {
		return WindowsResult{}, ErrInvalidTimeRange
	}
	if query.Step < time.Minute {
		return WindowsResult{}, ErrInvalidStep
	}
	if query.Limit <= 0 {
		query.Limit = 100
	}

	events, err := s.repo.QueryEvents(ctx, query.From, query.To, query.Host)
	if err != nil {
		return WindowsResult{}, err
	}

	type sessionSet map[string]struct{}
	type bucket struct {
		metric   traffic.WindowMetric
		sessions sessionSet
	}

	buckets := make(map[string]*bucket)
	for _, event := range events {
		windowStart := event.OccurredAt.UTC().Truncate(query.Step)
		windowEnd := windowStart.Add(query.Step)
		bucketKey := fmt.Sprintf("%s|%s", windowStart.Format(time.RFC3339Nano), event.Host)
		current, ok := buckets[bucketKey]
		if !ok {
			current = &bucket{
				metric: traffic.WindowMetric{
					WindowStart: windowStart,
					WindowEnd:   windowEnd,
					Host:        event.Host,
				},
				sessions: make(sessionSet),
			}
			buckets[bucketKey] = current
		}
		current.metric.Requests++
		current.sessions[event.SessionID] = struct{}{}
		switch event.Class {
		case traffic.ClassificationBot:
			current.metric.Bot++
		case traffic.ClassificationHuman:
			current.metric.Human++
		default:
			current.metric.Unknown++
		}
	}

	items := make([]traffic.WindowMetric, 0, len(buckets))
	for _, current := range buckets {
		current.metric.Sessions = len(current.sessions)
		items = append(items, current.metric)
	}

	sort.Slice(items, func(i, j int) bool {
		if items[i].WindowStart.Equal(items[j].WindowStart) {
			return items[i].Host < items[j].Host
		}
		return items[i].WindowStart.Before(items[j].WindowStart)
	})

	if query.Cursor >= len(items) {
		return WindowsResult{Items: []traffic.WindowMetric{}}, nil
	}

	end := query.Cursor + query.Limit
	if end > len(items) {
		end = len(items)
	}

	result := WindowsResult{Items: items[query.Cursor:end]}
	if end < len(items) {
		next := fmt.Sprintf("%d", end)
		result.NextCursor = &next
	}
	return result, nil
}

func eventID(parsed traffic.ParsedLog, sessionID string) string {
	raw := fmt.Sprintf(
		"%s|%s|%s|%s|%d|%s|%s",
		parsed.OccurredAt.UTC().Format(time.RFC3339Nano),
		parsed.Host,
		parsed.Path,
		parsed.Method,
		parsed.StatusCode,
		parsed.RemoteAddr,
		sessionID,
	)
	return hashString(raw)
}

func hashString(value string) string {
	sum := sha256.Sum256([]byte(value))
	return hex.EncodeToString(sum[:16])
}
