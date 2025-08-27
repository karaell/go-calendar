package main

import (
	"fmt"
	"github.com/karaell/app/calendar"
	"github.com/karaell/app/cmd"
	"github.com/karaell/app/logger"
	"github.com/karaell/app/storage"
)

func main() {
	err := logger.Init("app.log")
	if err != nil {
		fmt.Println("Can't initialize logger: ", err)
	} else {
		logger.Info("init logger")
	}

	s := storage.CreateJsonStorage("calendar.json")
	//s := storage.CreateZipStorage("calendar.zip")
	logger.Info("create storage")

	c := calendar.CreateCalendar(s)
	logger.Info("create calendar")

	err = c.Load()
	if err != nil {
		fmt.Println("Load calendar file failed, but you can still work with calendar: ", err)
		logger.Error("calendar file load failed")
	}

	ls := storage.CreateJsonStorage("calendar_log.json")
	logger.Info("create calendar log history storage")

	cli := cmd.CreateCmd(c, ls)
	logger.Info("create cmd")

	err = cli.LoadLog()
	if err != nil {
		fmt.Println("Load history log file failed, but you can still work with calendar: ", err)
		logger.Error("calendar history log file load failed")
	}

	cli.Run()
}
