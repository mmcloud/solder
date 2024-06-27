// file: pkg/service/recommendation.go
package service

import (
	"context"

	"github.com/mmcloud/solder/pkg/command"
	"github.com/mmcloud/solder/pkg/interfaces"
	"github.com/spf13/cobra"
)

var _ interfaces.Service = (*Serve)(nil)

type Serve struct {
	config    interfaces.ServiceConfig
	logger    interfaces.Logger
	commander interfaces.Commander
}

func NewServe(config interfaces.ServiceConfig, logger interfaces.Logger, commander interfaces.Commander) *Serve {
	return &Serve{
		config:    config,
		logger:    logger,
		commander: commander,
	}
}

func (s *Serve) GetCommand() interfaces.Command {
	serve := command.NewServe()
	service := serve.WithShort(s.config.GetShort()).WithLong((s.config.GetLong())).WithExample(s.config.GetExample())
	return service
}

func (s *Serve) RunE(cmd *cobra.Command, args []string) error {
	s.Run(cmd.Context())
	return nil
}

// Name returns the name of the service.
func (s *Serve) Name() string {
	return s.config.GetName()
}

// Short returns the short description of the service.
func (s *Serve) Short() string {
	return s.config.GetShort()
}

// Long returns the long description of the service.
func (s *Serve) Long() string {
	return s.config.GetLong()
}

// Run runs the service.
func (s *Serve) Run(ctx context.Context) error {
	s.logger.Debug("start recommendation service")
	return nil
}

// Stop stops the service.
func (s *Serve) Stop(ctx context.Context) error {
	s.logger.Debug("stop recommendation service")
	ctx.Done()
	return nil
}

func (s *Serve) Config() interfaces.ServiceConfig {
	return s.config
}
