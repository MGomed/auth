package app

import (
	"context"
	"flag"
	"net"

	grpc "google.golang.org/grpc"
	reflection "google.golang.org/grpc/reflection"

	env_config "github.com/MGomed/auth/internal/config/env"
	closer "github.com/MGomed/auth/pkg/closer"
	user_api "github.com/MGomed/auth/pkg/user_api"
)

var configPath string

// App represents object for starting grpc server
type App struct {
	serviceProvider *serviceProvider
	server          *grpc.Server
}

// NewApp is App struct constructor
func NewApp(ctx context.Context) (*App, error) {
	flag.StringVar(&configPath, "config-path", "build/.env", "path to config file")
	flag.Parse()

	app := &App{}

	if err := app.initDeps(ctx); err != nil {
		return nil, err
	}

	return app, nil
}

// Run starts grpc server
func (a *App) Run() error {
	defer func() {
		closer.CloseAll()
		closer.Wait()
	}()

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
	closer.Add(func() error {
		a.server.Stop()

		return nil
	})

	reflection.Register(a.server)

	user_api.RegisterUserAPIServer(a.server, a.serviceProvider.API(ctx))

	return nil
}

func (a *App) runGRPCServer() error {
	lis, err := net.Listen("tcp", a.serviceProvider.APIConfig().Address())
	if err != nil {
		return err
	}

	return a.server.Serve(lis)
}
