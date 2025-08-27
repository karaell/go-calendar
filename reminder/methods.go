package reminder

import (
	"time"
)

func (r *Reminder) Start() {
	duration := time.Until(r.SendAt)
	if duration < 0 {
		return
	}

	timer := time.AfterFunc(duration, r.Send)
	r.timer = timer
}

func (r *Reminder) Send() {
	if r.Sent {
		return
	}

	r.notifier(r.Message)
	r.Sent = true
}

func (r *Reminder) Stop() {
	if r.timer == nil {
		return
	}

	r.timer.Stop()
	r.timer = nil
}
