package events

import "errors"

var ErrInvalidTitle = errors.New("invalid title")
var ErrInvalidPriority = errors.New("invalid priority")
var ErrParseDate = errors.New("could not parse date")

var ErrEmptyMessage = errors.New("message is empty")
var ErrTooLateTime = errors.New("time is later than event")
var ErrTimeInPast = errors.New("time in past")

var ErrCreateEvent = errors.New("error create event")
