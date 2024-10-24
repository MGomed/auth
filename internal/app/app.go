package app

import (
	"context"
	"flag"
	"net"

	env_config "github.com/MGomed/auth/internal/config/env"
	"github.com/MGomed/auth/pkg/user_api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configPath string

type App struct {
	serviceProvider *serviceProvider
	server          *grpc.Server
}

func NewApp(ctx context.Context) (*App, error) {
	flag.StringVar(&configPath, "config-path", "build/.env", "path to config file")
	flag.Parse()

	app := &App{}

	if err := app.initDeps(ctx); err != nil {
		return nil, err
	}

	return app, nil
}

func (a *App) Run() error {
	return a.runGRPCServer()
}

func (a *App) initDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initConfig,
		a.initServiceProvider,
		a.initGRPCServer,
	}

	for _, f := range inits {
		if err := f(ctx); err != nil {
			return err
		}
	}

	return nil
}

func (a *App) initConfig(_ context.Context) error {
	if err := env_config.Load(configPath); err != nil {
		return err
	}

	return nil
}

func (a *App) initServiceProvider(_ context.Context) error {
	a.serviceProvider = newServiceProvider()

	return nil
}

func (a *App) initGRPCServer(ctx context.Context) error {
	a.server = grpc.NewServer()

	reflection.Register(a.server)

	user_api.RegisterUserAPIServer(a.server, a.serviceProvider.API(ctx))

	return nil
}

func (a *App) runGRPCServer() error {
	lis, err := net.Listen("tcp", a.serviceProvider.ApiConfig().Address())
	if err != nil {
		return err
	}

	return a.server.Serve(lis)
}
