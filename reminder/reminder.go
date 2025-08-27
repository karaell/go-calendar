package reminder

import "time"

type Reminder struct {
	Message  string    `json:"message"`
	SendAt   time.Time `json:"send_at"`
	Sent     bool      `json:"sent"`
	timer    *time.Timer
	notifier func(msg string)
}

func CreateReminder(message string, sendAt time.Time, notifier func(msg string)) *Reminder {
	return &Reminder{
		Message:  message,
		SendAt:   sendAt,
		Sent:     false,
		notifier: notifier,
	}
}
