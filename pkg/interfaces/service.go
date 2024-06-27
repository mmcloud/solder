package interfaces

import (
	"context"

	"github.com/mmcloud/solder/pkg/constants"
)

type ServiceFactory interface {
	// GetService returns the service.
	GetService() (Service, error)
	// GetServiceConfig returns the service config.
	GetServiceConfig() (ServiceConfig, error)
}
type ServiceConfig interface {
	GetKind() constants.ServiceKind
	GetName() string
	GetShort() string
	GetLong() string
	GetExample() string
}

// Service is an interface that defines the methods that a service must implement.
type Service interface {
	// Name returns the name of the service.
	Name() string
	// Short returns a short description of the service.
	Short() string
	// Long returns a long description of the service.
	Long() string
	// Run runs the service.
	Run(ctx context.Context) error
	// Stop stops the service.
	Stop(ctx context.Context) error
	// GetCommand returns the command.
	GetCommand() Command
}
