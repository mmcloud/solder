package service

import (
	"github.com/mmcloud/solder/pkg/interfaces"
)

// ProvideService provides a Service instance.
func ProvideServe(
	config interfaces.ServiceConfig,
	logger interfaces.Logger,
	commander interfaces.Commander,
) interfaces.Service { // Return the Service interface
	return ProvideServe(config, logger, commander)
}

func ProvideServiceFactory(logger interfaces.Logger,
	commander interfaces.Commander,
	config interfaces.Config) interfaces.ServiceFactory {
	sf, _ := NewServiceFactory(logger, commander, config)
	return sf

}
