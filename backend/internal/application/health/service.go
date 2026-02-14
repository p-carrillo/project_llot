package health

import (
	"context"
	"time"

	domain "github.com/diteria/project_llot/backend/internal/domain/health"
	"github.com/diteria/project_llot/backend/internal/ports"
)

type Service struct {
	checker ports.ReadinessChecker
}

func NewService(checker ports.ReadinessChecker) Service {
	return Service{checker: checker}
}

func (s Service) Live() domain.Snapshot {
	return domain.Snapshot{Status: domain.StatusOK, At: time.Now().UTC()}
}

func (s Service) Ready(ctx context.Context) domain.Snapshot {
	status := domain.StatusDegraded
	if s.checker.Ready(ctx) {
		status = domain.StatusOK
	}
	return domain.Snapshot{Status: status, At: time.Now().UTC()}
}
