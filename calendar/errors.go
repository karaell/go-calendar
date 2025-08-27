package calendar

import "errors"

var ErrNotFoundEvent = errors.New("event with target id not found")
var ErrSetEventReminder = errors.New("error setting event reminder")
var ErrCancelEventReminder = errors.New("error canceling event reminder")
var ErrAddEvent = errors.New("error adding event")
var ErrRemoveEvent = errors.New("error removing event")
var ErrUpdateEvent = errors.New("error updating event")
var ErrSaveCalendar = errors.New("error saving calendar")
var ErrLoadCalendar = errors.New("error loading calendar")
