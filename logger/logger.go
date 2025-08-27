package logger

import (
	"fmt"
	"log"
	"os"
)

type Logger struct {
	filename      string
	file          *os.File
	infoLogger    *log.Logger
	errorLogger   *log.Logger
	warningLogger *log.Logger
}

var infoLogger *log.Logger
var errorLogger *log.Logger
var warningLogger *log.Logger
var loggerFile *os.File

func Init(filename string) error {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)

	if err != nil {
		return fmt.Errorf("%w: %w", ErrInitLogger, err)

	}

	infoLogger = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	errorLogger = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	warningLogger = log.New(file, "WARN: ", log.Ldate|log.Ltime|log.Lshortfile)
	loggerFile = file

	return nil
}

func Info(msg string) {
	err := infoLogger.Output(2, msg)

	if err != nil {
		fmt.Printf("%s: %s", ErrSaveInfoLog, err)
	}
}

func Error(msg string) {
	err := errorLogger.Output(2, msg)

	if err != nil {
		fmt.Printf("%s: %s", ErrSaveErrorLog, err)
	}
}

func Warning(msg string) {
	err := warningLogger.Output(2, msg)

	if err != nil {
		fmt.Printf("%s: %s", ErrSaveWarningLog, err)
	}
}

func Close() error {
	if loggerFile == nil {
		return nil
	}

	err := loggerFile.Close()

	if err != nil {
		return fmt.Errorf("%w: %w", ErrCloseFile, err)
	}

	return nil
}
