package app

import (
	"context"
	"log"

	auth_api "github.com/MGomed/auth/internal/api/auth"
	config "github.com/MGomed/auth/internal/config"
	env_config "github.com/MGomed/auth/internal/config/env"
	service "github.com/MGomed/auth/internal/service"
	auth_service "github.com/MGomed/auth/internal/service/auth"
	storage "github.com/MGomed/auth/internal/storage"
	auth_cache "github.com/MGomed/auth/internal/storage/cache/auth"
	auth_repo "github.com/MGomed/auth/internal/storage/repository/auth"
	cache "github.com/MGomed/auth/pkg/client/cache"
	redis "github.com/MGomed/auth/pkg/client/cache/redis"
	db "github.com/MGomed/auth/pkg/client/db"
	pg "github.com/MGomed/auth/pkg/client/db/pg"
	transaction "github.com/MGomed/auth/pkg/client/db/transaction"
	closer "github.com/MGomed/auth/pkg/closer"
	logger "github.com/MGomed/auth/pkg/logger"
	go_redis "github.com/gomodule/redigo/redis"
)

type serviceProvider struct {
	logger *log.Logger

	pgConfig    config.PgConfig
	redisConfig config.RedisConfig
	apiConfig   config.APIConfig

	dbc         db.Client
	redisClient cache.RedisClient
	txMgr       db.TxManager

	cache storage.Cache
	repo  storage.Repository

	service service.Service

	api *auth_api.UserAPI
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

// RedisConfig init/get redis config
func (p *serviceProvider) RedisConfig() config.RedisConfig {
	if p.redisConfig == nil {
		cfg, err := env_config.NewRedisConfig()
		if err != nil {
			log.Fatalf("failed to create pg config: %v", err)
		}

		p.redisConfig = cfg
	}

	return p.redisConfig
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

		if err := dbc.DB().Ping(ctx); err != nil {
			log.Fatalf("ping error: %v", err)
		}
		closer.Add(dbc.Close)

		p.dbc = dbc
	}

	return p.dbc
}

// RedisClient init/get RedisClient
func (p *serviceProvider) RedisClient(ctx context.Context) cache.RedisClient {
	if p.redisClient == nil {
		pool := &go_redis.Pool{
			MaxIdle:     p.RedisConfig().MaxIdle(),
			IdleTimeout: p.RedisConfig().IdleTimeout(),
			DialContext: func(ctx context.Context) (go_redis.Conn, error) {
				return go_redis.DialContext(ctx, "tcp", p.RedisConfig().Address())
			},
		}

		client := redis.NewClient(p.Logger(), pool, p.RedisConfig().ConnectionTimeout())

		if err := client.Ping(ctx); err != nil {
			log.Fatalf("ping error: %v", err.Error())
		}
		closer.Add(pool.Close)

		p.redisClient = client
	}

	return p.redisClient
}

// TxManager init/get TxManager
func (p *serviceProvider) TxManager(ctx context.Context) db.TxManager {
	if p.txMgr == nil {
		p.txMgr = transaction.NewTransactionManager(p.DBClient(ctx).DB())
	}

	return p.txMgr
}

// Repository init/get Repository
func (p *serviceProvider) Repository(ctx context.Context) storage.Repository {
	if p.repo == nil {
		p.repo = auth_repo.NewRepository(p.DBClient(ctx))
	}

	return p.repo
}

// Cache init/get Cache
func (p *serviceProvider) Cache(ctx context.Context) storage.Cache {
	if p.cache == nil {
		p.cache = auth_cache.NewCacher(p.RedisClient(ctx))
	}

	return p.cache
}

// Service init/get Service(usecases)
func (p *serviceProvider) Service(ctx context.Context) service.Service {
	if p.service == nil {
		p.service = auth_service.NewService(p.Logger(), p.Repository(ctx), p.Cache(ctx), p.TxManager(ctx))
	}

	return p.service
}

// API init/get API(grpc implementation)
func (p *serviceProvider) API(ctx context.Context) *auth_api.UserAPI {
	if p.api == nil {
		p.api = auth_api.NewAPI(p.Logger(), p.Service(ctx))
	}

	return p.api
}
