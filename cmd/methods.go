package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/c-bata/go-prompt"
	"github.com/google/shlex"
	"github.com/karaell/app/calendar"
	"github.com/karaell/app/events"
	"github.com/karaell/app/logger"
	"github.com/karaell/app/storage"
	"strings"
	"sync"
	"time"
)

var mu sync.Mutex

func (c *Cmd) executor(input string) {
	input = strings.TrimSpace(input)
	if input == "" {
		return
	}

	parts, err := shlex.Split(input)
	if err != nil {
		c.print("Error: " + err.Error())
	}

	c.addToLog(input)

	cmd := strings.ToLower(parts[0])

	switch cmd {
	case "add":
		err = handleAdd(c, parts)

		if errors.Is(err, ErrCmdFormat) {
			c.print("Can't add event because of invalid command format, format is: " + AddCmd)
			return
		}

		if err != nil {
			c.print("Can't add event: " + err.Error())
		}
	case "update":
		err = handleUpdate(c, parts)

		switch {
		case errors.Is(err, ErrCmdFormat):
			c.print("Can't update event because of invalid command format, format is: " + UpdateCmd)
		case errors.Is(err, calendar.ErrNotFoundEvent):
			c.print("Can't update event because target event is not found")
		case errors.Is(err, events.ErrInvalidTitle):
			c.print("Can't update event because event title is invalid")
		case errors.Is(err, events.ErrInvalidPriority):
			c.print("Can't update event because target event priority is invalid")
		case errors.Is(err, events.ErrParseDate):
			c.print("Can't update event because selected date is invalid")
		default:
			if err != nil {
				c.print("Can't update event: " + err.Error())
			}
		}
	case "remove":
		err = handleRemove(c, parts)

		if errors.Is(err, ErrCmdFormat) {
			c.print("Can't remove event because of invalid command format, format is: " + RemoveCmd)
			return
		}

		if err != nil {
			c.print("Can't remove event: " + err.Error())
		}
	case "add_event_reminder":
		err = handleSetEventReminder(c, parts)

		switch {
		case errors.Is(err, ErrCmdFormat):
			c.print("Can't set event reminder because of invalid command format, format is: " + SetEventReminderCmd)
		case errors.Is(err, calendar.ErrNotFoundEvent):
			c.print("Can't set event reminder because target event is not found")
		case errors.Is(err, events.ErrEmptyMessage):
			c.print("Can't set event reminder with empty message")
		case errors.Is(err, events.ErrTooLateTime):
			c.print("Can't set event reminder because reminder time is later than event time")
		case errors.Is(err, events.ErrTimeInPast):
			c.print("Can't set event reminder because reminder time is earlier than current time")
		default:
			if err != nil {
				c.print("Can't set event reminder: " + err.Error())
			}
		}
	case "remove_event_reminder":
		err = handleCancelEventReminder(c, parts)

		if errors.Is(err, ErrCmdFormat) {
			c.print("Can't cancel event reminder because of invalid command format, format is: " + CancelEventReminderCmd)
			return
		}

		if err != nil {
			c.print("Can't remove event reminder: " + err.Error())
		}
	case "help":
		handleHelp(c)
	case "list":
		handleList(c)
	case "exit":
		handleExit(c)
	case "log":
		handleLog(c)
	default:
		c.print("Unknown command: " + cmd)
	}
}

func (c *Cmd) completer(d prompt.Document) []prompt.Suggest {
	suggestions := []prompt.Suggest{
		{Text: "add", Description: "Add event"},
		{Text: "update ", Description: "Update event"},
		{Text: "remove", Description: "Remove event"},
		{Text: "add_event_reminder", Description: "Add reminder for event"},
		{Text: "remove_event_reminder", Description: "Cancel reminder for event"},
		{Text: "list", Description: "Show event list"},
		{Text: "help", Description: "Show help"},
		{Text: "log", Description: "Show console history log"},
		{Text: "exit", Description: "Exit from program"},
	}

	return prompt.FilterHasPrefix(suggestions, d.GetWordBeforeCursor(), true)
}

func (c *Cmd) addToLog(line string) {
	mu.Lock()

	log := Log{
		Date: time.Now(),
		Line: line,
	}

	c.log = append(c.log, log)
	logger.Info("add console output to history log")

	mu.Unlock()
}

func (c *Cmd) print(line string) {
	fmt.Println(line)
	c.addToLog(line)
}

func (c *Cmd) SaveLog() error {
	data, err := json.Marshal(c.log)

	if err != nil {
		logger.Error("save console history log failed")
		return fmt.Errorf("%w: %w", ErrSaveHistoryLog, err)
	}

	err = c.storage.Save(data)
	if err != nil {
		logger.Error("save console history log failed")
		return fmt.Errorf("%w: %w", ErrSaveHistoryLog, err)
	}

	logger.Info("save console history log success")
	return nil
}

func (c *Cmd) LoadLog() error {
	data, err := c.storage.Load()
	if err != nil {
		logger.Error("load console history log failed")
		return fmt.Errorf("%w: %w", ErrLoadHistoryLog, err)
	}

	if len(data) == 0 {
		logger.Warning("load console history log failed, empty storage file")
		return fmt.Errorf("%w: %w", ErrLoadHistoryLog, storage.ErrEmptyFile)
	}

	err = json.Unmarshal(data, &c.log)
	if err != nil {
		logger.Error("load console history log failed")
		return fmt.Errorf("%w: %w", ErrLoadHistoryLog, err)
	}

	logger.Info("load console history log success")
	return nil
}

func (c *Cmd) Run() {
	go func() {
		for msg := range c.calendar.Notification {
			c.print("Reminder: " + msg)
		}
	}()

	p := prompt.New(
		c.executor,
		c.completer,
		prompt.OptionPrefix("> "),
	)

	logger.Info("cli run")
	p.Run()
}
