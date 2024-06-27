package interfaces

import (
	"context"
)

// Commander represents the root command of the app
type Commander interface {
	ExecuteContext(ctx context.Context) error
	AddCommand(cmd Command)
}
