package log

import (
	"log"
)

type Priority int

const (
	EMERG Priority = iota
	ALERT
	CRIT
	ERR
	WARNING
	NOTICE
	INFO
	DEBUG
)

type Logger struct {
	internal *log.Logger
	level    Priority
}

var gLogger *Logger
var (
	gLogger *Logger = nil
)

func init() {
	gLogger = newLogger()
}

func getLogger() *Logger {
	return gLogger
}

func newLogger() *Logger {
	l := Logger()
	l.internal = nil
	l.level = EMERG
	return &l
}

func Set(l *log.Logger, lv Priority) {
	gLogger.internal = l
	gLogger.level = lv
}

func (l *Logger) write(Priority lv, msg string, v ...interface{}) {
	if l.level > lv {
		return
	}

	if v == nil || len(v) == 0 {
		l.internal.Print(msg)
	} else {
		l.internal.Printf(msg, v...)
	}
}

func (l *Logger) Debug(msg string, v ...interface{}) {
	write(DEBUG, msg, v...)
}

func (l *Logger) Info(msg string, v ...interface{}) {
	write(INFO, msg, v...)
}

func (l *Logger) Notice(msg string, v ...interface{}) {
	write(NOTICE, msg, v...)
}

func (l *Logger) Warning(msg string, v ...interface{}) {
	write(WARNING, msg, v...)
}

func (l *Logger) Error(msg string, v ...interface{}) {
	write(ERR, msg, v...)
}

func (l *Logger) Critical(msg string, v ...interface{}) {
	write(CRIT, msg, v...)
}

func (l *Logger) Alert(msg string, v ...interface{}) {
	write(ALERT, msg, v...)
}

func (l *Logger) Emergency(msg string, v ...interface{}) {
	write(EMERG, msg, v...)
}
