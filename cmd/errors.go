package cmd

import "errors"

var ErrCmdFormat = errors.New("error command format")
var ErrSaveHistoryLog = errors.New("error save console history log")
var ErrLoadHistoryLog = errors.New("error load console history log")
var ErrAddCmd = errors.New("error add command")
var ErrUpdateCmd = errors.New("error update command")
var ErrRemoveCmd = errors.New("error remove command")
var ErrSetEventReminderCmd = errors.New("error set event reminder command")
var ErrCancelEventReminderCmd = errors.New("error cancel event reminder command")
var ErrExitCmd = errors.New("error exit command")
