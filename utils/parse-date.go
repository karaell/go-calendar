package utils

import (
	"fmt"
	"github.com/araddon/dateparse"
	"time"
)

func ParseDate(date string) (time.Time, error) {
	parsedDate, err := dateparse.ParseLocal(date)
	if err != nil {
		return time.Time{}, fmt.Errorf("%w: %w", ErrParseDate, err)
	}

	return parsedDate, nil
}
