package logging

import "github.com/google/wire"

var Set = wire.NewSet(
	NewLoggerFactory,
	ProvideLogger,
)
