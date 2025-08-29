package calendar

import (
	"encoding/json"
	"fmt"
	"github.com/karaell/app/events"
	"github.com/karaell/app/logger"
	"github.com/karaell/app/storage"
)

func (c *Calendar) AddEvent(title string, startAt string, priority events.Priority) (*events.Event, error) {
	e, err := events.CreateEvent(title, startAt, priority)
	if err != nil {
		logger.Error("add event failed")
		return nil, fmt.Errorf("error adding event: %w", err)
	}

	c.calendarEvents[e.ID] = e

	logger.Info("add event success")
	return e, nil
}

func (c *Calendar) RemoveEvent(id string) error {
	_, ok := c.calendarEvents[id]
	if !ok {
		logger.Error("remove event failed, event not found")
		return fmt.Errorf("error removing event: %w Id: %s", events.ErrNotFoundEvent, id)
	}

	delete(c.calendarEvents, id)

	logger.Info("remove event success")
	return nil
}

func (c *Calendar) UpdateEvent(id string, title string, startAt string, priority events.Priority) error {
	e, ok := c.calendarEvents[id]
	if !ok {
		logger.Error("update event failed, event not found")
		return fmt.Errorf("error updating event: %w Id: %s", events.ErrNotFoundEvent, id)
	}

	err := e.Update(title, startAt, priority)
	if err != nil {
		logger.Error("update event failed")
		return fmt.Errorf("error updating event: %w", err)
	}

	logger.Info("update event success")
	return nil
}

func (c *Calendar) SetEventReminder(id string, message string, sendAt string) error {
	e, ok := c.calendarEvents[id]
	if !ok {
		logger.Error("set event reminder failed, event not found")
		return fmt.Errorf("error setting event reminder: %w Id: %s", events.ErrNotFoundEvent, id)
	}

	err := e.AddReminder(message, sendAt, c.Notify)
	if err != nil {
		logger.Error("set event reminder failed")
		return fmt.Errorf("error setting event reminder: %w", err)
	}
	logger.Info("set event reminder success")
	return nil
}

func (c *Calendar) CancelEventReminder(id string) error {
	e, ok := c.calendarEvents[id]
	if !ok {
		logger.Error("cancel event reminder failed, event not found")
		return fmt.Errorf("error canceling event reminder: %w Id: %s", events.ErrNotFoundEvent, id)
	}

	e.RemoveReminder()

	logger.Info("cancel event reminder success")
	return nil
}

func (c *Calendar) Notify(msg string) {
	c.Notification <- msg
}

func (c *Calendar) GetEvents() map[string]*events.Event {
	return c.calendarEvents
}

func (c *Calendar) Save() error {
	data, err := json.Marshal(c.calendarEvents)
	if err != nil {
		logger.Error("save calendar failed, marshal failed")
		return fmt.Errorf("error saving calendar: %w", err)
	}

	err = c.storage.Save(data)
	if err != nil {
		logger.Error("save calendar failed, saving to storage failed")
		return fmt.Errorf("error saving calendar: %w", err)
	}

	logger.Info("save calendar success")
	return nil
}

func (c *Calendar) Load() error {
	data, err := c.storage.Load()
	if err != nil {
		logger.Error("load calendar failed, loading from storage failed")
		return fmt.Errorf("error loading calendar: %w", err)
	}

	if len(data) == 0 {
		logger.Warning("load calendar failed, empty storage file")
		return fmt.Errorf("error loading calendar: %w", storage.ErrEmptyFile)
	}

	err = json.Unmarshal(data, &c.calendarEvents)
	if err != nil {
		logger.Error("load calendar failed, unmarshal failed")
		return fmt.Errorf("error loading calendar: %w", err)
	}

	logger.Info("load calendar success")
	return nil
}
