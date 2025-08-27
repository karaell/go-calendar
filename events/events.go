package events

import (
	"fmt"
	"github.com/karaell/app/logger"
	"github.com/karaell/app/reminder"
	"github.com/karaell/app/utils"
	"time"
)

type Event struct {
	ID       string             `json:"id"`
	Title    string             `json:"title"`
	StartAt  time.Time          `json:"start_at"`
	Priority Priority           `json:"priority"`
	Reminder *reminder.Reminder `json:"reminder"`
}

func CreateEvent(title string, startAt string, priority Priority) (*Event, error) {
	logger.Info(fmt.Sprintf("adding event"))

	if !isValidTitle(title) {
		logger.Error(fmt.Sprintf("invalid title for event: %s", title))
		return nil, fmt.Errorf("%w: %w - %s", ErrCreateEvent, ErrInvalidTitle, title)
	}

	date, err := utils.ParseDate(startAt)
	if err != nil {
		logger.Error(fmt.Sprintf("failed to parse date for event: %s", startAt))
		return nil, fmt.Errorf("%w: %w", ErrCreateEvent, err)
	}

	if date.Before(time.Now()) {
		logger.Error(fmt.Sprintf("start time in past for event: %s", startAt))
		return nil, fmt.Errorf("%w: %w", ErrCreateEvent, ErrTimeInPast)
	}

	err = priority.Validate()
	if err != nil {
		logger.Error(fmt.Sprintf("invalid priority for event: %s", priority))
		return nil, fmt.Errorf("%w: %w", ErrCreateEvent, err)
	}

	id := utils.GetUniqId()

	event := Event{
		ID:       id,
		Title:    title,
		StartAt:  date,
		Priority: priority,
	}
	logger.Info(fmt.Sprintf("event %s added successfully", id))
	return &event, nil
}
