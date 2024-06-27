package pkg

import (
	"github.com/google/wire"
	"github.com/mmcloud/solder/pkg/app"
	"github.com/mmcloud/solder/pkg/command"
	"github.com/mmcloud/solder/pkg/config"
	"github.com/mmcloud/solder/pkg/logging"
	"github.com/mmcloud/solder/pkg/service"
)

// SharedSet defines the base common dependencies for all microservices
var CommonSet = wire.NewSet(
	config.Set, // provide config
	service.ProvideServiceFactory,
	command.ProviderCommander,
	logging.Set, // provide logging
	app.Set,
)
