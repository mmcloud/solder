package command

import (
	"github.com/mmcloud/solder/pkg/interfaces"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var _ interfaces.Serve = (*Serve)(nil)

type Serve struct {
	cmd          cobra.Command
	postRunHooks []func()
}

func NewServe() *Serve {
	return &Serve{cmd: cobra.Command{
		Use: "serve",
	}}
}

func (s *Serve) GetCobraCmd() *cobra.Command {
	return &s.cmd
}

func (s *Serve) WithShort(description string) interfaces.Serve {
	s.cmd.Short = description
	return s
}

func (s *Serve) WithLong(description string) interfaces.Serve {
	s.cmd.Long = description
	return s
}

func (s *Serve) WithExample(example string) interfaces.Serve {
	s.cmd.Example = example
	return s
}

func (s *Serve) PersistentPreRunE(hook func(*cobra.Command, []string) error) interfaces.Serve {
	s.cmd.PersistentPreRunE = hook
	return s
}

func (s *Serve) PersistentPreRun(hook func(*cobra.Command, []string)) interfaces.Serve {
	s.cmd.PersistentPreRun = hook
	return s
}

func (s *Serve) WithPostRunHook(hook func()) interfaces.Serve {
	s.postRunHooks = append(s.postRunHooks, hook)
	return s
}
func (s *Serve) Run(cmd *cobra.Command, args []string) {
	for _, hook := range s.postRunHooks {
		hook()
	}
	s.cmd.Run(cmd, args)
}

func (s *Serve) RunE(cmd *cobra.Command, args []string) error {
	for _, hook := range s.postRunHooks {
		hook()
	}
	return s.cmd.RunE(cmd, args)
}

func (s *Serve) WithRun(run func(cmd *cobra.Command, args []string)) interfaces.Serve {
	s.cmd.Run = run
	return s
}

func (s *Serve) WithRunE(run func(cmd *cobra.Command, args []string) error) interfaces.Serve {
	s.cmd.RunE = run
	return s
}

func (s *Serve) Execute() error {
	return s.cmd.Execute()
}

func (s *Serve) PersistentFlags() *pflag.FlagSet {
	return s.cmd.PersistentFlags()
}

func (s *Serve) Flags() *pflag.FlagSet {
	return s.cmd.Flags()
}
