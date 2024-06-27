package interfaces

import (
	"github.com/mmcloud/solder/pkg/constants"
)

// LoggerFactory is an interface for the LoggerFactory.
type LoggerFactory interface {
	// GetLogger returns a Logger instance.
	GetLogger() (Logger, error)
	// GetLoggerConfig returns a LoggingConfig instance.
	GetLoggerConfig() LoggingConfig
}

// LoggingConfig is an interface for the LoggingConfig.
type LoggingConfig interface {
	// GetLevel returns the logging level.
	GetLevel() constants.Level
	// GetLoggerKind returns the logger Kind.
	GetLoggerKind() constants.LoggerKind
}

// Logger is an interface for the Logger.
type Logger interface {
	// Debug logs a message at the debug level.
	Debug(msg string)
	// Info logs a message at the info level.
	Info(msg string)
	// Warn logs a message at the warn level.
	Warn(msg string)
	// Error logs a message at the error level.
	Error(msg string)
	// Fatal logs a message at the fatal level.
	Fatal(msg string)
	// SetLevel sets the logging level.
	SetLevel(level constants.Level)
	// GetLevel returns the logging level.
	GetLevel() constants.Level
	// GetLoggerKind returns the logger Kind.
	GetLoggerKind() constants.LoggerKind
	// Debugf logs a formatted message at the debug level.
	Debugf(format string, args ...any)
	// Infof logs a formatted message at the info level.
	Infof(format string, args ...any)
	// Warnf logs a formatted message at the warn level.
	Warnf(format string, args ...any)
	// Errorf logs a formatted message at the error level.
	Errorf(format string, args ...any)
	// Fatalf logs a formatted message at the fatal level.
	Fatalf(format string, args ...any)
}
