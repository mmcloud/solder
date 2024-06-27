// file: pkg/command/root.go
package command

import (
	"context"

	"github.com/mmcloud/solder/pkg/interfaces"
	"github.com/spf13/cobra"
)

var _ interfaces.Commander = (*Commander)(nil)

type Commander struct {
	cmd *cobra.Command
}

func NewCommander() *Commander {
	c := &Commander{}
	cmd := &cobra.Command{
		Use:   "solder",
		Short: "A DI framework for microservices",
		Long:  "A DI framework for microservices",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}
	c.cmd = cmd
	return c
}

func ProviderCommander() interfaces.Commander {
	return NewCommander()
}

func (c *Commander) ExecuteContext(ctx context.Context) error {
	return c.cmd.ExecuteContext(ctx)
}

func (c *Commander) AddCommand(cmd interfaces.Command) {
	command := cmd.GetCobraCmd()
	c.cmd.AddCommand(command)
}
