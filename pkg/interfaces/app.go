package interfaces

import "context"

// App represents the main application.
type App interface {
	// Start starts the application.
	Start(ctx context.Context) error
	// Stop stops the application.
	Stop(ctx context.Context) error
}
