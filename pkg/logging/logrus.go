package logging

import (
	"github.com/mmcloud/solder/pkg/constants"
	"github.com/mmcloud/solder/pkg/interfaces"
	"github.com/sirupsen/logrus"
)

var _ interfaces.Logger = (*LogrusLogger)(nil)

// LogrusLogger is a logger implementation using logrus.
type LogrusLogger struct {
	Logger *logrus.Logger
}

// NewLogrusLogger creates a new instance of LogrusLogger.
func NewLogrusLogger(config interfaces.LoggingConfig) *LogrusLogger {
	ll := &LogrusLogger{
		Logger: logrus.New(),
	}
	ll.SetLevel(config.GetLevel())
	return ll
}

// SetLevel sets the logging level.
func (l *LogrusLogger) SetLevel(level constants.Level) {
	parsedLevel := l.parseLogLevel(level)
	l.Logger.SetLevel(parsedLevel)
}

// Debug logs a message at the debug level.
func (l *LogrusLogger) Debug(msg string) {
	l.Logger.Debug(msg)
}

// Info logs a message at the info level.
func (l *LogrusLogger) Info(msg string) {
	l.Logger.Info(msg)
}

// Warn logs a message at the warn level.
func (l *LogrusLogger) Warn(msg string) {
	l.Logger.Warn(msg)
}

// Error logs a message at the error level.
func (l *LogrusLogger) Error(msg string) {
	l.Logger.Error(msg)
}

// Fatal logs a message at the fatal level.
func (l *LogrusLogger) Fatal(msg string) {
	// For testing purposes, we will log Fatal as an error with a "FATAL: " prefix
	l.Logger.Error("FATAL: " + msg)
}

// GetLevel returns the logging level.
func (l *LogrusLogger) GetLevel() constants.Level {
	return l.convertLogrusLevel(l.Logger.GetLevel())
}

// GetLoggerKind returns the logger Kind.
func (l *LogrusLogger) GetLoggerKind() constants.LoggerKind {
	return constants.LogrusLoggerKind
}

// Debugf logs a formatted message at the debug level.
func (l *LogrusLogger) Debugf(format string, args ...interface{}) {
	l.Logger.Debugf(format, args...)
}

// Infof logs a formatted message at the info level.
func (l *LogrusLogger) Infof(format string, args ...interface{}) {
	l.Logger.Infof(format, args...)
}

// Warnf logs a formatted message at the warn level.
func (l *LogrusLogger) Warnf(format string, args ...interface{}) {
	l.Logger.Warnf(format, args...)
}

// Errorf logs a formatted message at the error level.
func (l *LogrusLogger) Errorf(format string, args ...interface{}) {
	l.Logger.Errorf(format, args...)
}

// Fatalf logs a formatted message at the fatal level.
func (l *LogrusLogger) Fatalf(format string, args ...interface{}) {
	l.Logger.Fatalf(format, args...)
}

// parseLogLevel converts the logging level to logrus level.
func (l *LogrusLogger) parseLogLevel(level constants.Level) logrus.Level {
	switch level {
	case constants.DebugLevel:
		return logrus.DebugLevel
	case constants.InfoLevel:
		return logrus.InfoLevel
	case constants.WarnLevel:
		return logrus.WarnLevel
	case constants.ErrorLevel:
		return logrus.ErrorLevel
	case constants.FatalLevel:
		return logrus.FatalLevel
	default:
		return l.parseLogLevel(constants.DefaultLevel)
	}
}

// convertLogrusLevel converts the logrus level to logging level.
func (l *LogrusLogger) convertLogrusLevel(level logrus.Level) constants.Level {
	switch level {
	case logrus.DebugLevel:
		return constants.DebugLevel
	case logrus.InfoLevel:
		return constants.InfoLevel
	case logrus.WarnLevel:
		return constants.WarnLevel
	case logrus.ErrorLevel:
		return constants.ErrorLevel
	case logrus.FatalLevel:
		return constants.FatalLevel
	default:
		return constants.DefaultLevel
	}
}
