package config

import (
	"github.com/google/wire"
)

// Set provides a Wire provider for the Config struct.
var Set = wire.NewSet(
	ProvideDefaultConfig,
	NewViper,
	LoadConfig,
)
