package logging

import "errors"

var (
	// ErrUnsupportedLoggerKind is returned when an unsupported logger Kind is specified.
	ErrUnsupportedLoggerKind = errors.New("unsupported logger kind")
	// ErrInvalidLogLevel is returned when an invalid log level is specified.
	ErrInvalidLogLevel = errors.New("invalid log level")
	// ErrInvalidLogger is returned when an invalid logger is specified.
	ErrInvalidLogger = errors.New("invalid logger")
	// ErrLoggingConfigNotFound is returned when a logging config is not found.
	ErrLoggingConfigNotFound = errors.New("logging config not found")
)
