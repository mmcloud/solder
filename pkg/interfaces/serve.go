package interfaces

import "github.com/spf13/cobra"

type Command interface {
	// GetCobraCmd returns the underlying cobra.Command instance.
	GetCobraCmd() *cobra.Command
}

type Serve interface {
	Command
	WithShort(description string) Serve
	WithLong(description string) Serve
	WithExample(example string) Serve
	Execute() error
	WithRunE(runE func(cmd *cobra.Command, args []string) error) Serve
	WithRun(run func(cmd *cobra.Command, args []string)) Serve
}
