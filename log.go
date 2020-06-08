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

func init() {
	gLogger = newLogger()
}

func Get() *Logger {
	return gLogger
}

func newLogger() *Logger {
	l := Logger{}
	l.internal = nil
	l.level = EMERG
	return &l
}

func Set(l *log.Logger, lv Priority) {
	gLogger.internal = l
	gLogger.level = lv
}

func (l *Logger) write(lv Priority, msg string, v ...interface{}) {
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
	l.write(DEBUG, msg, v...)
}

func (l *Logger) Info(msg string, v ...interface{}) {
	l.write(INFO, msg, v...)
}

func (l *Logger) Notice(msg string, v ...interface{}) {
	l.write(NOTICE, msg, v...)
}

func (l *Logger) Warning(msg string, v ...interface{}) {
	l.write(WARNING, msg, v...)
}

func (l *Logger) Error(msg string, v ...interface{}) {
	l.write(ERR, msg, v...)
}

func (l *Logger) Critical(msg string, v ...interface{}) {
	l.write(CRIT, msg, v...)
}

func (l *Logger) Alert(msg string, v ...interface{}) {
	l.write(ALERT, msg, v...)
}

func (l *Logger) Emergency(msg string, v ...interface{}) {
	l.write(EMERG, msg, v...)
}
