package cmd

import (
	"github.com/karaell/app/calendar"
	"github.com/karaell/app/storage"
)

type Cmd struct {
	calendar *calendar.Calendar
	storage  storage.Store
	log      []Log
}

func CreateCmd(c *calendar.Calendar, storage storage.Store) *Cmd {
	return &Cmd{
		calendar: c,
		storage:  storage,
	}
}
