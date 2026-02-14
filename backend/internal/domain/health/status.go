package health

import "time"

type Status string

const (
	StatusOK       Status = "ok"
	StatusDegraded Status = "degraded"
)

type Snapshot struct {
	Status Status    `json:"status"`
	At     time.Time `json:"at"`
}
