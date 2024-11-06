package app

import (
	"context"
	"log"

	auth_api "github.com/MGomed/auth/internal/api/auth"
	config "github.com/MGomed/auth/internal/config"
	env_config "github.com/MGomed/auth/internal/config/env"
	repository "github.com/MGomed/auth/internal/repository"
	auth_repo "github.com/MGomed/auth/internal/repository/auth"
	service "github.com/MGomed/auth/internal/service"
	auth_service "github.com/MGomed/auth/internal/service/auth"
	db "github.com/MGomed/auth/pkg/client/db"
	pg "github.com/MGomed/auth/pkg/client/db/pg"
	transaction "github.com/MGomed/auth/pkg/client/db/transaction"
	closer "github.com/MGomed/auth/pkg/closer"
	logger "github.com/MGomed/auth/pkg/logger"
)

type serviceProvider struct {
	logger *log.Logger

	pgConfig  config.PgConfig
	apiConfig config.APIConfig

	dbc   db.Client
	txMgr db.TxManager

	repo repository.Repository

	service service.Service

	api *auth_api.API
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

// PgConfig init/get postgres config
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

// APIConfig init/get api(grpc) config
func (p *serviceProvider) APIConfig() config.APIConfig {
	if p.apiConfig == nil {
		cfg, err := env_config.NewAPIConfig()
		if err != nil {
			log.Fatalf("failed to create api config: %v", err)
		}

		p.apiConfig = cfg
	}

	return p.apiConfig
}

// Logger init/get logger
func (p *serviceProvider) Logger() *log.Logger {
	if p.logger == nil {
		p.logger = logger.InitLogger()
	}

	return p.logger
}

// DBClient init/get DBClient
func (p *serviceProvider) DBClient(ctx context.Context) db.Client {
	if p.dbc == nil {
		dbc, err := pg.New(ctx, p.Logger(), p.PgConfig().DSN())
		if err != nil {
			log.Fatalf("failed to create db client: %v", err)
		}

		err = dbc.DB().Ping(ctx)
		if err != nil {
			log.Fatalf("ping error: %s", err.Error())
		}
		closer.Add(dbc.Close)

		p.dbc = dbc
	}

	return p.dbc
}

func (p *serviceProvider) TxManager(ctx context.Context) db.TxManager {
	if p.txMgr == nil {
		p.txMgr = transaction.NewTransactionManager(p.DBClient(ctx).DB())
	}

	return p.txMgr
}

// Repository init/get Repository
func (p *serviceProvider) Repository(ctx context.Context) repository.Repository {
	if p.repo == nil {
		p.repo = auth_repo.NewRepository(p.DBClient(ctx))
	}

	return p.repo
}

// Service init/get Service(usecases)
func (p *serviceProvider) Service(ctx context.Context) service.Service {
	if p.service == nil {
		p.service = auth_service.NewService(p.Logger(), p.Repository(ctx), p.TxManager(ctx))
	}

	return p.service
}

// API init/get API(grpc implementation)
func (p *serviceProvider) API(ctx context.Context) *auth_api.API {
	if p.api == nil {
		p.api = auth_api.NewAPI(p.Logger(), p.Service(ctx))
	}

	return p.api
}
