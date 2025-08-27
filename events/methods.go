package events

import (
	"fmt"
	"github.com/karaell/app/logger"
	"github.com/karaell/app/reminder"
	"github.com/karaell/app/utils"
	"time"
)

func (e *Event) Update(title string, startAt string, priority Priority) error {
	logger.Info(fmt.Sprintf("updating event %s", e.ID))

	if title != "" {
		if !isValidTitle(title) {
			logger.Error(fmt.Sprintf("invalid title for event %s: %s", e.ID, title))
			return fmt.Errorf("%w: %s", ErrInvalidTitle, title)
		}

		e.Title = title
		logger.Info(fmt.Sprintf("event %s title updated to: %s", e.ID, title))
	}

	if startAt != "" {
		date, err := utils.ParseDate(startAt)
		if err != nil {
			logger.Error(fmt.Sprintf("failed to parse date for event %s: %s", e.ID, startAt))
			return err
		}

		if date.Before(time.Now()) {
			logger.Error(fmt.Sprintf("start time in past for event %s: %s", e.ID, startAt))
			return ErrTimeInPast
		}

		e.StartAt = date
		logger.Info(fmt.Sprintf("event %s start time updated to: %s", e.ID, date))
	}

	if priority != "" {
		err := priority.Validate()
		if err != nil {
			logger.Error(fmt.Sprintf("invalid priority for event %s: %s", e.ID, priority))
			return err
		}

		e.Priority = priority
		logger.Info(fmt.Sprintf("event %s priority updated to: %s", e.ID, priority))
	}

	logger.Info(fmt.Sprintf("event %s updated successfully", e.ID))
	return nil
}

func (e *Event) AddReminder(message string, sendAt string, notifier func(msg string)) error {
	logger.Info(fmt.Sprintf("adding reminder to event %s", e.ID))

	if len(message) == 0 {
		logger.Error(fmt.Sprintf("empty message for reminder in event %s", e.ID))
		return ErrEmptyMessage
	}

	date, err := utils.ParseDate(sendAt)
	if err != nil {
		logger.Error(fmt.Sprintf("failed to parse reminder date for event %s: %s", e.ID, sendAt))
		return err
	}

	if date.Before(time.Now()) {
		logger.Warning(fmt.Sprintf("reminder time in past for event %s: %s", e.ID, sendAt))
		return ErrTimeInPast
	}

	if date.After(e.StartAt) {
		logger.Warning(fmt.Sprintf("reminder time after event start time for event %s: %s", e.ID, sendAt))
		return ErrTooLateTime
	}

	if e.Reminder != nil {
		e.RemoveReminder()
	}

	e.Reminder = reminder.CreateReminder(message, date, notifier)
	e.Reminder.Start()

	logger.Info(fmt.Sprintf("reminder added successfully to event %s", e.ID))
	return nil
}

func (e *Event) RemoveReminder() {
	logger.Info(fmt.Sprintf("removing reminder for event %s", e.ID))

	if e.Reminder != nil {
		e.Reminder.Stop()
		e.Reminder = nil

		logger.Info(fmt.Sprintf("reminder removed successfully for event %s", e.ID))
	}
}
