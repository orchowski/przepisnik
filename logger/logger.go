package logger

import (
	"fmt"
	"os"
	"sync"
)

type LogLevel int

const (
	INFO LogLevel = iota
	WARNING
	ERROR
)

const (
	colorRed    = "\033[31m"
	colorYellow = "\033[33m"
	colorReset  = "\033[0m"
)

type llogger struct {
	level LogLevel
}

var (
	loggers = make(map[LogLevel]*llogger)
	mu      sync.Mutex
)

func NewLogger(level LogLevel) {
	mu.Lock()
	defer mu.Unlock()

	if _, ok := loggers[level]; ok {
		panic("Logger exists!!")
	}

	logger := &llogger{level}
	loggers[level] = logger
}

func Log(level LogLevel, format string, args ...interface{}) {
	mu.Lock()
	defer mu.Unlock()

	if logger, ok := loggers[level]; ok {
		logger.log(level, format, args...)
	}
}

func (l *llogger) log(level LogLevel, format string, args ...interface{}) {
	var out *os.File

	switch level {
	case INFO:
		out = os.Stdout
	case WARNING:
		out = os.Stdout
	case ERROR:
		out = os.Stderr
	}
	msg := fmt.Sprintf(format, args...)
	if level == WARNING {
		msg = colorYellow + msg + colorReset
	}
	if l.level <= level {
		fmt.Fprintf(out, "[%s] %s\n", levelString(level), fmt.Sprintf(format, args...))
	}
}

func levelString(level LogLevel) string {
	switch level {
	case INFO:
		return "INFO"
	case WARNING:
		return "WARNING"
	case ERROR:
		return "ERROR"
	default:
		return "UNKNOWN"
	}
}
