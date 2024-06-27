// file: pkg/command/set.go
package command

import (
	"github.com/google/wire"
)

var Set = wire.NewSet(
	NewCommander,
)
