package log

import (
	"context"
	stdlog "log"
)

type Logger struct {
	StdLog *stdlog.Logger
}

var defaultLogger *Logger

func init() {
	defaultLogger = &Logger{StdLog: stdlog.Default()}
}

func (l *Logger) Info(msg string, keysAndValues ...interface{}) {
	l.StdLog.Println("[INFO]", msg, keysAndValues)
}

func (l *Logger) Error(err error, msg string, keysAndValues ...interface{}) {
	l.StdLog.Println("[ERROR]", msg, err, keysAndValues)
}

func (l *Logger) Set() *Logger {
	defaultLogger = l
	return l
}

func FromCtx(ctx context.Context) *Logger {
	if defaultLogger == nil {
		defaultLogger = &Logger{StdLog: stdlog.Default()}
	}
	return defaultLogger
}