package config

import (
	"io"
	"log"
	"os"
)

// Logger is a custom logger with different levels of logging.
type Logger struct {
	debug   *log.Logger
	info    *log.Logger
	warning *log.Logger
	err     *log.Logger
	writer  io.Writer
}

// NewLogger creates a new Logger instance with the specified prefix.
func NewLogger(p string) *Logger {
	writer := io.Writer(os.Stdout)
	logger := log.New(writer, p, log.Ldate|log.Ltime)

	return &Logger{
		debug:   log.New(writer, ">> DEBUG: ", logger.Flags()),
		info:    log.New(writer, ">> INFO: ", logger.Flags()),
		warning: log.New(writer, ">> WARN: ", logger.Flags()),
		err:     log.New(writer, ">> ERROR: ", logger.Flags()),
		writer:  writer,
	}
}

// Debug logs a debug message.
func (l *Logger) Debug(v ...interface{}) {
	l.debug.Println(v...)
}

// Info logs an info message.
func (l *Logger) Info(v ...interface{}) {
	l.info.Println(v...)
}

// Warn logs a warning message.
func (l *Logger) Warn(v ...interface{}) {
	l.warning.Println(v...)
}

// Error logs an error message.
func (l *Logger) Error(v ...interface{}) {
	l.err.Println(v...)
}

// Debugf logs a formatted debug message.
func (l *Logger) Debugf(format string, v ...interface{}) {
	l.debug.Printf(format, v...)
}

// Infof logs a formatted info message.
func (l *Logger) Infof(format string, v ...interface{}) {
	l.info.Printf(format, v...)
}

// Warnf logs a formatted warning message.
func (l *Logger) Warnf(format string, v ...interface{}) {
	l.warning.Printf(format, v...)
}

// Errorf logs a formatted error message.
func (l *Logger) Errorf(format string, v ...interface{}) {
	l.err.Printf(format, v...)
}
