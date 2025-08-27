package cmd

const (
	AddCmd                 = "add \"title\" \"date and time\" \"priority\""
	RemoveCmd              = "remove \"id\""
	UpdateCmd              = "update \"id\" \"title\" \"date and time\" \"priority\""
	SetEventReminderCmd    = "add_event_reminder \"id\" \"message\" \"date and time\""
	CancelEventReminderCmd = "remove_event_reminder \"id\""
	ListCmd                = "list"
	HelpCmd                = "help"
	ExitCmd                = "exit"
	LogCmd                 = "log"
)
