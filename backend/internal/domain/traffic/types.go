package traffic

import "time"

type Classification string

const (
	ClassificationHuman   Classification = "human"
	ClassificationBot     Classification = "bot"
	ClassificationUnknown Classification = "unknown"
)

type RequestEvent struct {
	ID             string         `json:"id"`
	OccurredAt     time.Time      `json:"occurred_at"`
	Host           string         `json:"host"`
	Path           string         `json:"path"`
	Method         string         `json:"method"`
	StatusCode     int            `json:"status_code"`
	RemoteAddrHash string         `json:"remote_addr_hash"`
	UserAgent      string         `json:"user_agent"`
	BotScore       float64        `json:"bot_score"`
	Class          Classification `json:"class"`
	SessionID      string         `json:"session_id"`
}

type ParsedLog struct {
	OccurredAt time.Time
	Host       string
	Path       string
	Method     string
	StatusCode int
	RemoteAddr string
	UserAgent  string
}

type Overview struct {
	From        time.Time `json:"from"`
	To          time.Time `json:"to"`
	Host        string    `json:"host,omitempty"`
	Requests    int       `json:"requests"`
	Human       int       `json:"human"`
	Bot         int       `json:"bot"`
	Unknown     int       `json:"unknown"`
	UniqueHosts int       `json:"unique_hosts"`
	Sessions    int       `json:"sessions"`
}

type WindowMetric struct {
	WindowStart time.Time `json:"window_start"`
	WindowEnd   time.Time `json:"window_end"`
	Host        string    `json:"host,omitempty"`
	Requests    int       `json:"requests"`
	Human       int       `json:"human"`
	Bot         int       `json:"bot"`
	Unknown     int       `json:"unknown"`
	Sessions    int       `json:"sessions"`
}
