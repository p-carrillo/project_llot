package traffic

import "errors"

var (
	ErrInvalidTimeRange = errors.New("invalid time range")
	ErrInvalidStep      = errors.New("invalid step")
)
