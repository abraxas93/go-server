package logger

import (
	"errors"
	"log"
	"os"
)

// ANSI color escape codes
const (
	ColorRed    = "\033[31m"
	ColorGreen  = "\033[32m"
	ColorYellow = "\033[33m"
	ColorReset  = "\033[0m"
)

type Logger struct {
	infoLogger *log.Logger
	warnLogger *log.Logger
	errLogger  *log.Logger
}

var logger Logger

func (l *Logger) Info(format string, v ...interface{}) {
	l.infoLogger.Printf(ColorGreen+format+ColorReset, v...)
}

func (l *Logger) Warn(format string, v ...interface{}) {
	l.warnLogger.Printf(ColorYellow+format+ColorReset, v...)
}

func (l *Logger) Err(format string, v ...interface{}) {
	l.errLogger.Printf(ColorRed+format+ColorReset, v...)
}

func InitLogger(messages map[string]string, flags int) {
	logger = Logger{
		infoLogger: log.New(os.Stdout, ColorGreen+messages["Info"], flags),
		warnLogger: log.New(os.Stdout, ColorYellow+messages["Warn"], flags),
		errLogger:  log.New(os.Stdout, ColorRed+messages["Err"], flags),
	}
}

func GetLogger() (Logger, error) {
	if (logger != Logger{}) {
		return logger, nil
	}
	return Logger{}, errors.New("logger not initialised")
}
