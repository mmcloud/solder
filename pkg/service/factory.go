// file: pkg/service/service.go
package service

import (
	"github.com/mmcloud/solder/pkg/constants"
	"github.com/mmcloud/solder/pkg/interfaces"
)

type ServiceFactory struct {
	config    interfaces.ServiceConfig
	logger    interfaces.Logger
	commander interfaces.Commander
}

var _ interfaces.ServiceFactory = (*ServiceFactory)(nil)

// NewServiceFactory returns a new service factory.
func NewServiceFactory(
	logger interfaces.Logger,
	commander interfaces.Commander,
	config interfaces.Config,
) (*ServiceFactory, error) {
	serviceConfig := config.GetServiceConfig()
	if serviceConfig == nil {
		return nil, ErrServiceConfigNotFound
	}
	return &ServiceFactory{config: serviceConfig, logger: logger, commander: commander}, nil
}

func (sf *ServiceFactory) GetServiceConfig() (interfaces.ServiceConfig, error) {
	if sf.config == nil {
		return nil, ErrServiceConfigNotFound
	}
	return sf.config, nil
}

func (sf *ServiceFactory) build() (interfaces.Service, error) {
	var s interfaces.Service
	switch sf.config.GetKind() {
	case constants.ServeKind:
		s = NewServe(sf.config, sf.logger, sf.commander)
	default:
		return nil, ErrUnsupportedServiceKind
	}
	sf.logger.Infof("Service %s built", sf.config.GetName())

	return s, nil
}

func (sf *ServiceFactory) GetService() (interfaces.Service, error) {
	return sf.build()
}

var _ interfaces.ServiceConfig = (*ServiceConfig)(nil)

// ServiceConfig is a struct that holds the configuration for a service.
type ServiceConfig struct {
	// Kind is the kind of the service.
	Kind constants.ServiceKind
	// Name is the name of the service.
	Name string `mapstructure:"name"`
	// Description is the description of the service.
	Description string `mapstructure:"description"`
	// Short is a short description of the service.
	Short string `mapstructure:"short"`
	// Long is a long description of the service.
	Long string `mapstructure:"long"`
	// Example is the example execution
	Example string `mapstructure:"example"`
}

// GetName returns the name of the service.
func (sc *ServiceConfig) GetName() string {
	return sc.Name
}

// GetShort returns a short description of the service.
func (sc *ServiceConfig) GetShort() string {
	return sc.Short
}

// GetLong returns a long description of the service.
func (sc *ServiceConfig) GetLong() string {
	return sc.Long
}

// GetServiceKind returns the kind of the service.
func (sc *ServiceConfig) GetKind() constants.ServiceKind {
	return sc.Kind
}

func (sc *ServiceConfig) GetExample() string {
	return sc.Example
}
