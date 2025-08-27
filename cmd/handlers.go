package cmd

import (
	"fmt"
	"github.com/karaell/app/events"
	"github.com/karaell/app/logger"
	"os"
	"time"
)

func handleAdd(c *Cmd, parts []string) error {
	if len(parts) < 4 {
		return fmt.Errorf("%w: %w", ErrAddCmd, ErrCmdFormat)
	}

	title := parts[1]
	date := parts[2]
	priority := events.Priority(parts[3])

	_, err := c.calendar.AddEvent(title, date, priority)
	if err != nil {
		return fmt.Errorf("%w: %w", ErrAddCmd, err)
	}
	return nil
}

func handleRemove(c *Cmd, parts []string) error {
	if len(parts) < 2 {
		return fmt.Errorf("%w: %w", ErrRemoveCmd, ErrCmdFormat)
	}

	id := parts[1]

	err := c.calendar.RemoveEvent(id)
	if err != nil {
		return fmt.Errorf("%w: %w", ErrRemoveCmd, err)
	}
	return nil
}

func handleUpdate(c *Cmd, parts []string) error {
	if len(parts) < 5 {
		return fmt.Errorf("%w: %w", ErrUpdateCmd, ErrCmdFormat)
	}

	id := parts[1]
	title := parts[2]
	date := parts[3]
	priority := events.Priority(parts[4])

	err := c.calendar.UpdateEvent(id, title, date, priority)
	if err != nil {
		return fmt.Errorf("%w: %w", ErrUpdateCmd, err)
	}
	return nil
}

func handleSetEventReminder(c *Cmd, parts []string) error {
	if len(parts) < 4 {
		return fmt.Errorf("%w: %w", ErrSetEventReminderCmd, ErrCmdFormat)
	}

	id := parts[1]
	message := parts[2]
	date := parts[3]

	err := c.calendar.SetEventReminder(id, message, date)
	if err != nil {
		return fmt.Errorf("%w: %w", ErrSetEventReminderCmd, err)
	}
	return nil
}

func handleCancelEventReminder(c *Cmd, parts []string) error {
	if len(parts) < 2 {
		return fmt.Errorf("%w: %w", ErrCancelEventReminderCmd, ErrCmdFormat)
	}

	id := parts[1]

	err := c.calendar.CancelEventReminder(id)
	if err != nil {
		return fmt.Errorf("%w: %w", ErrCancelEventReminderCmd, err)
	}
	return nil
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
		c.print(fmt.Errorf("%w: %w", ErrExitCmd, err).Error())
	}

	err = c.SaveLog()
	if err != nil {
		logger.Error("saving history log failed")
		c.print(fmt.Errorf("%w: %w", ErrExitCmd, err).Error())
	}

	close(c.calendar.Notification)
	logger.Info("app exit successful")

	err = logger.Close()
	if err != nil {
		c.print(fmt.Errorf("%w: %w", ErrExitCmd, err).Error())
	}

	os.Exit(0)
}

func handleLog(c *Cmd) {
	for _, l := range c.log {
		fmt.Println(l.Date.Format(time.RFC850) + ": " + l.Line)
	}
}
