package app

import (
	"context"
	"os"
	"os/signal"

	"github.com/mmcloud/solder/pkg/interfaces"
)

func NewApp(
	commander interfaces.Commander,
	config interfaces.Config,
	logger interfaces.Logger,
	serviceFactory interfaces.ServiceFactory,
) interfaces.App {
	return &App{
		Commander:      commander,
		Config:         config,
		Logger:         logger,
		ServiceFactory: serviceFactory,
	}
}

var _ interfaces.App = (*App)(nil)

type App struct {
	Commander      interfaces.Commander
	Config         interfaces.Config
	Logger         interfaces.Logger
	ServiceFactory interfaces.ServiceFactory
	Tracer         interfaces.Tracer
}

func (a *App) Start(ctx context.Context) error {
	// handle SIGTERM
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	span, ctx := a.Tracer.StartSpan(ctx, "app.Start")
	defer span.End()
	service, err := a.ServiceFactory.GetService()
	if err != nil {
		return err
	}
	a.Commander.AddCommand(service.GetCommand())
	return a.Commander.ExecuteContext(ctx)
}

func (a *App) Stop(ctx context.Context) error {
	span, ctx := a.Tracer.StartSpan(ctx, "App.Stop")
	defer span.End()
	return nil
}
