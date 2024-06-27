package interfaces

type Config interface {
	GetLoggingConfig() LoggingConfig
	GetServiceConfig() ServiceConfig
}
