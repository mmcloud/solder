package constants

import (
	"encoding/json"
	"errors"
	"strings"
)

var DefaultLevel = DebugLevel

// Level represents the logging level.
type Level string

// Logging levels.
const (
	DebugLevel Level = "DEBUG"
	InfoLevel  Level = "INFO"
	WarnLevel  Level = "WARN"
	ErrorLevel Level = "ERROR"
	FatalLevel Level = "FATAL"
)

func (l Level) String() string {
	return strings.ToUpper(string(l))
}

func (l *Level) UnmarshalJSON(text []byte) error {
	var s string
	if err := json.Unmarshal(text, &s); err != nil {
		return err
	}
	return l.UnmarshalText([]byte(s))
}

// UnmarshalText implements the encoding.TextUnmarshaler interface for Level.
func (l *Level) UnmarshalText(text []byte) error {
	switch strings.ToUpper(string(text)) {
	case "DEBUG":
		*l = DebugLevel
	case "INFO":
		*l = InfoLevel
	case "WARN":
		*l = WarnLevel
	case "ERROR":
		*l = ErrorLevel
	case "FATAL":
		*l = FatalLevel
	default:
		return errors.New("invalid log level")
	}
	return nil
}

// LoggerKind represents the logger type.
type LoggerKind string

// String returns the string representation of the Loggerkind.
func (lt LoggerKind) String() string {
	return string(lt)
}

const (
	// LogrusLoggerKind represents the logrus logger type.
	LogrusLoggerKind LoggerKind = "logrus"
	// SlogLoggerKind represents the slog logger type.
	SlogLoggerKind LoggerKind = "slog"
)
