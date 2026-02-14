package health

import "context"

type StaticChecker struct {
	ready bool
}

func NewStaticChecker(ready bool) StaticChecker {
	return StaticChecker{ready: ready}
}

func (s StaticChecker) Ready(context.Context) bool {
	return s.ready
}
