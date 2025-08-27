package events

import (
	"fmt"
)

type Priority string

const (
	PriorityLow    Priority = "low"
	PriorityMedium Priority = "medium"
	PriorityHigh   Priority = "high"
)

func (p Priority) Validate() error {
	switch p {
	case PriorityLow, PriorityMedium, PriorityHigh:
		return nil
	default:
		return fmt.Errorf("%w: %s", ErrInvalidPriority, p)
	}
}

func (p Priority) String() string {
	return string(p)
}
