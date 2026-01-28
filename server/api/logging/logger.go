package logging

import (
	"os"

	"github.com/rs/zerolog"
)

// Logger is the global structured logger instance
var Logger zerolog.Logger

func init() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	Logger = zerolog.New(os.Stderr).With().Timestamp().Logger()
}

// Error logs an error with additional structured fields
func Error(err error, msg string, fields map[string]interface{}) {
	event := Logger.Error().Err(err)
	for k, v := range fields {
		event = event.Interface(k, v)
	}
	event.Msg(msg)
}

// Info logs an informational message with additional structured fields
func Info(msg string, fields map[string]interface{}) {
	event := Logger.Info()
	for k, v := range fields {
		event = event.Interface(k, v)
	}
	event.Msg(msg)
}

// Warn logs a warning message with additional structured fields
func Warn(msg string, fields map[string]interface{}) {
	event := Logger.Warn()
	for k, v := range fields {
		event = event.Interface(k, v)
	}
	event.Msg(msg)
}

// Debug logs a debug message with additional structured fields
func Debug(msg string, fields map[string]interface{}) {
	event := Logger.Debug()
	for k, v := range fields {
		event = event.Interface(k, v)
	}
	event.Msg(msg)
}
