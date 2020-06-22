package main

import (
	"fmt"
	"os"
	"runtime"
	"strings"
	"time"

	"golang.org/x/crypto/ssh/terminal"
)

// LogLevel represents the logger's level.
type LogLevel int

// Logger level constants.
const (
	Trace LogLevel = iota
	Debug
	Info
	Warn
	Error
	Fatal
)

// LogLevels contains available log levels as strings.
var LogLevels []string = []string{"trace", "debug", "info", "warn", "error", "fatal"}

// String returns the log level as a string.
func (l LogLevel) String() string {
	return LogLevels[l]
}

// Type return the type of LogLevel.
func (l LogLevel) Type() string {
	return "string"
}

// Set the log level.
func (l *LogLevel) Set(level string) error {
	switch strings.ToLower(level) {
	case "trace":
		*l = Trace
	case "debug":
		*l = Debug
	case "info":
		*l = Info
	case "warn", "warning":
		*l = Warn
	case "error":
		*l = Error
	case "fatal":
		*l = Fatal
	default:
		return fmt.Errorf("invalid log level, use one of %v (case insensitive)", LogLevels)
	}

	return nil
}

// UnmarshalYAML unmarshals the logger LogLevel type.
func (l LogLevel) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var s string

	if err := unmarshal(&s); err != nil {
		return err
	}

	if err := l.Set(s); err != nil {
		return err
	}

	return nil
}

// LogMessage holds a log message.
type LogMessage struct {
	Date    time.Time
	Level   LogLevel
	File    string
	Line    int
	Message string
}

// Log represents the Log object.
type Log struct {
	Level         LogLevel
	Size          int
	Color         bool
	ShutdownFuncs []func()
	Logs          []LogMessage
}

// LogInit creates a new logger instance.
func LogInit(level LogLevel, size int, color bool) Log {
	return Log{
		Level: level,
		Size:  size,
		Color: color,
	}
}

// RegisterShutdownFunc registers a shutdown function.
func (l *Log) RegisterShutdownFunc(fn func()) {
	l.ShutdownFuncs = append(l.ShutdownFuncs, fn)
}

// GetLogs returns n logs from the log buffer.
func (l *Log) GetLogs(n int) []LogMessage {
	return l.Logs[len(l.Logs)-n:]
}

// Log a message of the given level.
func (l *Log) Log(callerSkip int, level LogLevel, format string, args ...interface{}) {
	if level < l.Level {
		return
	}

	now := time.Now()
	message := strings.ReplaceAll(fmt.Sprintf(format, args...), "\n", " ")
	_, file, line, ok := runtime.Caller(callerSkip)

	if !ok {
		file = "???"
		line = 0
	}

	l.Logs = append(l.Logs, LogMessage{
		Date:    now,
		Level:   level,
		File:    file,
		Line:    line,
		Message: message,
	})

	if len(l.Logs) > l.Size {
		l.Logs = l.Logs[len(l.Logs)-l.Size:]
	}

	if terminal.IsTerminal(int(os.Stdout.Fd())) {
		if l.Color {
			color := "0"

			switch level {
			case Debug:
				color = "37"
			case Info:
				color = "34"
			case Warn:
				color = "33"
			case Error, Fatal:
				color = "91"
			}

			fmt.Printf("%s \033[%sm%-5s\033[0m %s:%d: %s\n", now.Format("2006-01-02 15:04:05"), color, strings.ToUpper(level.String()), file, line, message)
		} else {
			fmt.Printf("%s %-5s %s:%d: %s\n", now.Format("2006-01-02 15:04:05"), strings.ToUpper(level.String()), file, line, message)
		}
	}
}

// IsTrace returns true if the log level includes trace messages, otherwise false.
func (l *Log) IsTrace() bool {
	return l.Level <= Trace
}

// IsDebug returns true if the log level includes debug messages, otherwise false.
func (l *Log) IsDebug() bool {
	return l.Level <= Debug
}

// IsInfo returns true if the log level includes informational messages, otherwise false.
func (l *Log) IsInfo() bool {
	return l.Level <= Info
}

// IsWarn returns true if the log level includes warning messages, otherwise false.
func (l *Log) IsWarn() bool {
	return l.Level <= Warn
}

// IsError returns true if the log level includes error messages, otherwise false.
func (l *Log) IsError() bool {
	return l.Level <= Error
}

// Trace logs a message useful for tracing logic-level program flow.
func (l *Log) Trace(format string, args ...interface{}) {
	l.Log(2, Trace, format, args...)
}

// Debug logs a message useful for debugging issues.
func (l *Log) Debug(format string, args ...interface{}) {
	l.Log(2, Debug, format, args...)
}

// Info logs an informational message.
func (l *Log) Info(format string, args ...interface{}) {
	l.Log(2, Info, format, args...)
}

// Warn logs a warning message.
func (l *Log) Warn(format string, args ...interface{}) {
	l.Log(2, Warn, format, args...)
}

// Error logs an error message.
func (l *Log) Error(format string, args ...interface{}) {
	l.Log(2, Error, format, args...)
}

// Fatal logs a fatal message and exits the program.
func (l *Log) Fatal(format string, args ...interface{}) {
	l.Log(2, Fatal, format, args...)

	for _, fn := range l.ShutdownFuncs {
		fn()
	}

	os.Exit(1)
}
