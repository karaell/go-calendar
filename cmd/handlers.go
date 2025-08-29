package cmd

import (
	"errors"
	"fmt"
	"github.com/karaell/app/events"
	"github.com/karaell/app/logger"
	"os"
	"time"
)

func handleAdd(c *Cmd, parts []string) {
	if len(parts) < 4 {
		c.print("Can't add event because command format is incorrect, valid format: " + AddCmd)
		return
	}

	title := parts[1]
	date := parts[2]
	priority := events.Priority(parts[3])

	_, err := c.calendar.AddEvent(title, date, priority)

	if err != nil {
		switch {
		case errors.Is(err, events.ErrInvalidTitle):
			c.print("Can't add event because event title is invalid")
		case errors.Is(err, events.ErrInvalidPriority):
			c.print("Can't add event because event priority is invalid")
		case errors.Is(err, events.ErrParseDate):
			c.print("Can't add event because selected date is invalid")
		case errors.Is(err, events.ErrTimeInPast):
			c.print("Can't add event because selected date is invalid: time in past")
		default:
			c.print("Can't add event: " + err.Error())

		}
	}
}

func handleRemove(c *Cmd, parts []string) {
	if len(parts) < 2 {
		c.print("Can't remove event because command format is incorrect, valid format: " + RemoveCmd)
		return
	}

	id := parts[1]

	err := c.calendar.RemoveEvent(id)
	if err != nil {
		if errors.Is(err, events.ErrNotFoundEvent) {
			c.print("Can't remove event because event with target id not found. Id: " + id)
		} else {
			c.print("Can't remove event: " + err.Error())
		}
	}
}

func handleUpdate(c *Cmd, parts []string) {
	if len(parts) < 5 {
		c.print("Can't update event because command format is incorrect, valid format: " + UpdateCmd)
		return
	}

	id := parts[1]
	title := parts[2]
	date := parts[3]
	priority := events.Priority(parts[4])

	err := c.calendar.UpdateEvent(id, title, date, priority)
	if err != nil {
		switch {
		case errors.Is(err, events.ErrNotFoundEvent):
			c.print("Can't update event because event with target id not found. Id: " + id)
		case errors.Is(err, events.ErrInvalidTitle):
			c.print("Can't update event because event title is invalid")
		case errors.Is(err, events.ErrInvalidPriority):
			c.print("Can't update event because target event priority is invalid")
		case errors.Is(err, events.ErrParseDate):
			c.print("Can't update event because selected date is invalid")
		default:
			c.print("Can't update event: " + err.Error())
		}
	}

}

func handleSetEventReminder(c *Cmd, parts []string) {
	if len(parts) < 4 {
		c.print("Can't set event reminder because command format is incorrect, valid format: " + SetEventReminderCmd)
		return
	}

	id := parts[1]
	message := parts[2]
	date := parts[3]

	err := c.calendar.SetEventReminder(id, message, date)
	if err != nil {
		switch {
		case errors.Is(err, events.ErrNotFoundEvent):
			c.print("Can't set event reminder because event with target id not found. Id: " + id)
		case errors.Is(err, events.ErrEmptyMessage):
			c.print("Can't set event reminder because reminder message is empty")
		case errors.Is(err, events.ErrParseDate):
			c.print("Can't set event reminder because selected date is invalid")
		case errors.Is(err, events.ErrTooLateTime):
			c.print("Can't set event reminder because reminder time is later than event time")
		case errors.Is(err, events.ErrTimeInPast):
			c.print("Can't set event reminder because reminder time is earlier than current time")
		default:
			c.print("Can't set event reminder: " + err.Error())
		}
	}
}

func handleCancelEventReminder(c *Cmd, parts []string) {
	if len(parts) < 2 {
		c.print("Can't cancel event reminder because command format is incorrect, valid format: " + CancelEventReminderCmd)
		return
	}

	id := parts[1]

	err := c.calendar.CancelEventReminder(id)
	if err != nil {
		switch {
		case errors.Is(err, events.ErrNotFoundEvent):
			c.print("Can't cancel event reminder because event with target id not found. Id: " + id)
		default:
			c.print("Can't cancel event reminder: " + err.Error())
		}
	}
}

func handleList(c *Cmd) {
	calendarEvents := c.calendar.GetEvents()

	if len(calendarEvents) == 0 {
		c.print(events.NoEvents)
		return
	}

	for _, e := range calendarEvents {
		reminderMessage := "-"
		if e.Reminder != nil {
			reminderMessage = e.Reminder.Message
		}

		c.print(
			fmt.Sprintf("ID: %s\nTitle: %s\nStart At: %s\nPriority: %s\nReminder: %s\n",
				e.ID, e.Title, e.StartAt.Format(time.RFC3339), e.Priority.String(), reminderMessage))
	}
}

func handleHelp(c *Cmd) {
	c.print(fmt.Sprintf(
		"%s\n%s\n%s\n%s\n%s\n%s\n%s\n%s\n%s",
		AddCmd,
		UpdateCmd,
		RemoveCmd,
		SetEventReminderCmd,
		CancelEventReminderCmd,
		ListCmd,
		HelpCmd,
		ExitCmd,
		LogCmd,
	))
}

func handleExit(c *Cmd) {
	err := c.calendar.Save()
	if err != nil {
		logger.Error("saving calendar failed")
		c.print(fmt.Errorf("error exit command: %w", err).Error())
	}

	err = c.SaveLog()
	if err != nil {
		logger.Error("saving history log failed")
		c.print(fmt.Errorf("error exit command: %w", err).Error())
	}

	close(c.calendar.Notification)
	logger.Info("app exit successful")

	err = logger.Close()
	if err != nil {
		c.print(fmt.Errorf("error exit command: %w", err).Error())
	}

	os.Exit(0)
}

func handleLog(c *Cmd) {
	for _, l := range c.log {
		fmt.Println(l.Date.Format(time.RFC850) + ": " + l.Line)
	}
}
