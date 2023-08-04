package logger

import (
	"log"
	"os"
)

type Logger struct {
	infoLogger *log.Logger
	warnLogger *log.Logger
	errLogger  *log.Logger
}

func CreateLogger(messages map[string]string, flags int) *Logger {
	// flags := log.LstdFlags | log.Lshortfile

	return &Logger{
		infoLogger: log.New(os.Stdout, messages["Info"], flags),
		warnLogger: log.New(os.Stdout, messages["Warn"], flags),
		errLogger:  log.New(os.Stdout, messages["Err"], flags),
	}
}
