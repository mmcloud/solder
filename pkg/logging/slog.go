package logging

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/mmcloud/solder/pkg/constants"
	"github.com/mmcloud/solder/pkg/interfaces"
)

var _ interfaces.Logger = (*SlogLogger)(nil)

// SlogLogger is a logger that logs to slog.
type SlogLogger struct {
	Logger *slog.Logger
	Level  slog.Level
}

// NewSlogLogger creates a new SlogLogger.
func NewSlogLogger(config interfaces.LoggingConfig) *SlogLogger {
	sl := &SlogLogger{}

	l := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	sl.Logger = l
	sl.SetLevel(config.GetLevel()) // Ensure level is set during initialization

	return sl
}

// SetLevel sets the logging level.
func (l *SlogLogger) SetLevel(level constants.Level) {
	lev := &slog.LevelVar{}
	parsedLevel := l.parseLogLevel(level)
	lev.Set(parsedLevel)
	opts := &slog.HandlerOptions{
		Level: lev,
	}
	l.Logger = slog.New(slog.NewJSONHandler(os.Stdout, opts))
	l.Level = parsedLevel
}

// Debug logs a message at the debug level.
func (l *SlogLogger) Debug(msg string) {
	l.Logger.Debug(msg)
}

// Info logs a message at the info level.
func (l *SlogLogger) Info(msg string) {
	l.Logger.Info(msg)
}

// Warn logs a message at the warn level.
func (l *SlogLogger) Warn(msg string) {
	l.Logger.Warn(msg)
}

// Error logs a message at the error level.
func (l *SlogLogger) Error(msg string) {
	l.Logger.Error(msg)
}

// Fatal logs a message at the fatal level.
func (l *SlogLogger) Fatal(msg string) {
	l.Logger.Error("FATAL: " + msg)
	os.Exit(1)
}

// GetLevel returns the logging level.
func (l *SlogLogger) GetLevel() constants.Level {
	return l.convertSlogLevel(l.Level)
}

// GetLoggerKind returns the logger Kind.
func (l *SlogLogger) GetLoggerKind() constants.LoggerKind {
	return constants.SlogLoggerKind
}

// Debugf logs a formatted message at the debug level.
func (l *SlogLogger) Debugf(format string, args ...any) {
	msg := fmt.Sprintf(format, args...)
	l.Logger.Debug(msg)
}

// Infof logs a formatted message at the info level.
func (l *SlogLogger) Infof(format string, args ...any) {
	msg := fmt.Sprintf(format, args...)
	l.Logger.Info(msg)
}

// Warnf logs a formatted message at the warn level.
func (l *SlogLogger) Warnf(format string, args ...any) {
	msg := fmt.Sprintf(format, args...)
	l.Logger.Warn(msg)
}

// Errorf logs a formatted message at the error level.
func (l *SlogLogger) Errorf(format string, args ...any) {
	msg := fmt.Sprintf(format, args...)
	l.Logger.Error(msg)
}

// Fatalf logs a formatted message at the fatal level.
func (l *SlogLogger) Fatalf(format string, args ...any) {
	msg := fmt.Sprintf(format, args...)
	l.Logger.Error("FATAL: " + msg)
	os.Exit(1)
}

// parseLogLevel parses a string into a slog.Level.
func (l *SlogLogger) parseLogLevel(level constants.Level) slog.Level {
	switch level {
	case constants.DebugLevel:
		return slog.LevelDebug
	case constants.InfoLevel:
		return slog.LevelInfo
	case constants.WarnLevel:
		return slog.LevelWarn
	case constants.ErrorLevel:
		return slog.LevelError
	case constants.FatalLevel:
		return slog.LevelError
	default:
		return l.parseLogLevel(constants.DefaultLevel) // Default to DefaultLevel
	}
}

// convertSlogLevel converts a slog.Level to a Level.
func (l *SlogLogger) convertSlogLevel(level slog.Level) constants.Level {
	switch level {
	case slog.LevelDebug:
		return constants.DebugLevel
	case slog.LevelInfo:
		return constants.InfoLevel
	case slog.LevelWarn:
		return constants.WarnLevel
	case slog.LevelError:
		return constants.ErrorLevel
	default:
		return constants.DebugLevel // Default to debug level if the provided level is invalid
	}
}
