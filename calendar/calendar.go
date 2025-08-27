package calendar

import (
	"github.com/karaell/app/events"
	"github.com/karaell/app/storage"
)

type Calendar struct {
	calendarEvents map[string]*events.Event
	storage        storage.Store
	Notification   chan string
}

func CreateCalendar(storage storage.Store) *Calendar {
	return &Calendar{
		calendarEvents: make(map[string]*events.Event),
		storage:        storage,
		Notification:   make(chan string),
	}
}
