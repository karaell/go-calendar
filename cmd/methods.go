package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/c-bata/go-prompt"
	"github.com/google/shlex"
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
		handleAdd(c, parts)
	case "update":
		handleUpdate(c, parts)
	case "remove":
		handleRemove(c, parts)
	case "add_event_reminder":
		handleSetEventReminder(c, parts)
	case "remove_event_reminder":
		handleCancelEventReminder(c, parts)
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
		return fmt.Errorf("error save console history log: %w", err)
	}

	err = c.storage.Save(data)
	if err != nil {
		logger.Error("save console history log failed")
		return fmt.Errorf("error save console history log: %w", err)
	}

	logger.Info("save console history log success")
	return nil
}

func (c *Cmd) LoadLog() error {
	data, err := c.storage.Load()
	if err != nil {
		logger.Error("load console history log failed")
		return fmt.Errorf("error load console history log: %w", err)
	}

	if len(data) == 0 {
		logger.Warning("load console history log failed, empty storage file")
		return fmt.Errorf("error load console history log: %w", storage.ErrEmptyFile)
	}

	err = json.Unmarshal(data, &c.log)
	if err != nil {
		logger.Error("load console history log failed")
		return fmt.Errorf("error load console history log: %w", err)
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
