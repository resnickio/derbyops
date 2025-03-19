// Package logger provides structured logging functionality for the application.
package logger

import (
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// Setup initializes the logger with the specified configuration.
// It configures log level, output format, and adds common fields.
func Setup(isDevelopment bool) {
	// Set logger time format
	zerolog.TimeFieldFormat = time.RFC3339

	// Set global log level based on environment
	if isDevelopment {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
		// Pretty print logs in development
		log.Logger = log.Output(zerolog.ConsoleWriter{
			Out:        os.Stdout,
			TimeFormat: time.RFC3339,
		})
	} else {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
		// JSON format for production
		log.Logger = zerolog.New(os.Stdout).With().Timestamp().Logger()
	}
}

// Debug logs a debug message with the given fields
func Debug(msg string, fields map[string]interface{}) {
	event := log.Debug()
	for k, v := range fields {
		event = event.Interface(k, v)
	}
	event.Msg(msg)
}

// Info logs an info message with the given fields
func Info(msg string, fields map[string]interface{}) {
	event := log.Info()
	for k, v := range fields {
		event = event.Interface(k, v)
	}
	event.Msg(msg)
}

// Warn logs a warning message with the given fields
func Warn(msg string, fields map[string]interface{}) {
	event := log.Warn()
	for k, v := range fields {
		event = event.Interface(k, v)
	}
	event.Msg(msg)
}

// Error logs an error message with the given fields
func Error(msg string, err error, fields map[string]interface{}) {
	event := log.Error()
	if err != nil {
		event = event.Err(err)
	}
	for k, v := range fields {
		event = event.Interface(k, v)
	}
	event.Msg(msg)
}

// Fatal logs a fatal message with the given fields and exits the program
func Fatal(msg string, err error, fields map[string]interface{}) {
	event := log.Fatal()
	if err != nil {
		event = event.Err(err)
	}
	for k, v := range fields {
		event = event.Interface(k, v)
	}
	event.Msg(msg)
} 