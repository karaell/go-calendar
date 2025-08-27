package cmd

import "time"

type Log struct {
	Date time.Time `json:"date"`
	Line string    `json:"line"`
}
