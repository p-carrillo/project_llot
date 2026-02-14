package traffic

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"sync"
	"time"
)

type sessionState struct {
	lastSeen time.Time
	counter  int
}

type SessionEstimator struct {
	mu      sync.Mutex
	timeout time.Duration
	state   map[string]sessionState
}

func NewSessionEstimator(timeout time.Duration) *SessionEstimator {
	return &SessionEstimator{
		timeout: timeout,
		state:   make(map[string]sessionState),
	}
}

func (s *SessionEstimator) Estimate(visitorKey string, occurredAt time.Time) string {
	s.mu.Lock()
	defer s.mu.Unlock()

	current := s.state[visitorKey]
	if current.lastSeen.IsZero() || occurredAt.Sub(current.lastSeen) > s.timeout {
		current.counter++
	}
	current.lastSeen = occurredAt
	s.state[visitorKey] = current

	raw := fmt.Sprintf("%s|%d", visitorKey, current.counter)
	sum := sha256.Sum256([]byte(raw))
	return hex.EncodeToString(sum[:8])
}
