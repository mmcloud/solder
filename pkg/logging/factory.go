package logging

import (
	"github.com/mmcloud/solder/pkg/constants"
	"github.com/mmcloud/solder/pkg/interfaces"
)

type LoggerFactory struct {
	config interfaces.LoggingConfig
}

var _ interfaces.LoggerFactory = (*LoggerFactory)(nil)

// NewLoggerFactory creates a new instance of LoggerFactory.
func NewLoggerFactory(config interfaces.Config) (*LoggerFactory, error) {
	loggingConfig := config.GetLoggingConfig()
	if loggingConfig == nil {
		return nil, ErrLoggingConfigNotFound
	}
	return &LoggerFactory{config: loggingConfig}, nil
}

func (b *LoggerFactory) GetLoggerConfig() interfaces.LoggingConfig {
	return b.config
}

// Build creates a new logger based on the configuration.
func (b *LoggerFactory) build() (interfaces.Logger, error) {

	var l interfaces.Logger
	switch b.config.GetLoggerKind() {
	case constants.LogrusLoggerKind:
		l = NewLogrusLogger(b.config)
	case constants.SlogLoggerKind:
		l = NewSlogLogger(b.config)
	default:
		return nil, ErrUnsupportedLoggerKind
	}
	l.Infof("Using Logger: %s", b.config.GetLoggerKind().String())
	l.SetLevel(b.config.GetLevel())
	l.Infof("Logger Level: %s", l.GetLevel().String())
	return l, nil
}

// GetLogger provides a logger based on the configuration.
func (b *LoggerFactory) GetLogger() (interfaces.Logger, error) {
	return b.build()
}

func ProvideLogger(config interfaces.Config) (interfaces.Logger, error) {
	f, err := NewLoggerFactory(config)
	if err != nil {
		return nil, err
	}
	return f.GetLogger()
}

// LoggingConfig represents the logging configuration.
type LoggingConfig struct {
	// Level is the logging level.
	Level constants.Level `mapstructure:"level" json:"level"`
	// Kind is the logger Kind.
	LoggerKind constants.LoggerKind `mapstructure:"kind" json:"kind"`
}

// GetLoggerKind returns the logger Kind.
func (l *LoggingConfig) GetLoggerKind() constants.LoggerKind {
	return l.LoggerKind
}

// GetLevel returns the logging level.
func (l *LoggingConfig) GetLevel() constants.Level {
	return l.Level
}
