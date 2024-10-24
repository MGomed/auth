package app

import (
	"context"
	"log"

	pgxpool "github.com/jackc/pgx/v4/pgxpool"

	auth_api "github.com/MGomed/auth/internal/api/auth"
	config "github.com/MGomed/auth/internal/config"
	env_config "github.com/MGomed/auth/internal/config/env"
	repository "github.com/MGomed/auth/internal/repository"
	auth_repo "github.com/MGomed/auth/internal/repository/auth"
	service "github.com/MGomed/auth/internal/service"
	auth_service "github.com/MGomed/auth/internal/service/auth"
	logger "github.com/MGomed/auth/pkg/logger"
)

type serviceProvider struct {
	logger *log.Logger

	loggerConfig config.LoggerConfig
	pgConfig     config.PgConfig
	apiConfig    config.APIConfig

	repo repository.Repository

	service service.Service

	api *auth_api.API
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (p *serviceProvider) LoggerConfig() config.LoggerConfig {
	if p.loggerConfig == nil {
		cfg, err := env_config.NewLoggerConfig()
		if err != nil {
			log.Fatalf("failed to create logger config: %v", err)
		}

		p.loggerConfig = cfg
	}

	return p.loggerConfig
}

func (p *serviceProvider) PgConfig() config.PgConfig {
	if p.pgConfig == nil {
		cfg, err := env_config.NewPgConfig()
		if err != nil {
			log.Fatalf("failed to create pg config: %v", err)
		}

		p.pgConfig = cfg
	}

	return p.pgConfig
}

func (p *serviceProvider) ApiConfig() config.APIConfig {
	if p.apiConfig == nil {
		cfg, err := env_config.NewApiConfig()
		if err != nil {
			log.Fatalf("failed to create api config: %v", err)
		}

		p.apiConfig = cfg
	}

	return p.apiConfig
}

func (p *serviceProvider) Logger() *log.Logger {
	if p.logger == nil {
		logger, err := logger.InitLogger(p.LoggerConfig())
		if err != nil {
			log.Fatalf("failed to init logger: %v", err)
		}

		p.logger = logger
	}

	return p.logger
}

func (p *serviceProvider) Repository(ctx context.Context) repository.Repository {
	if p.repo == nil {
		pool, err := pgxpool.Connect(ctx, p.PgConfig().DSN())
		if err != nil {
			log.Fatalf("failed to connect to database: %v", err)
		}

		p.repo = auth_repo.NewRepository(pool)
	}

	return p.repo
}

func (p *serviceProvider) Service(ctx context.Context) service.Service {
	if p.service == nil {
		p.service = auth_service.NewService(p.Logger(), p.Repository(ctx))
	}

	return p.service
}

func (p *serviceProvider) API(ctx context.Context) *auth_api.API {
	if p.api == nil {
		p.api = auth_api.NewAPI(p.Logger(), p.Service(ctx))
	}

	return p.api
}
