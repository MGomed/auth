package app

import (
	"context"
	"log"

	sarama "github.com/IBM/sarama"

	auth_api "github.com/MGomed/auth/internal/api/auth"
	config "github.com/MGomed/auth/internal/config"
	env_config "github.com/MGomed/auth/internal/config/env"
	service "github.com/MGomed/auth/internal/service"
	auth_service "github.com/MGomed/auth/internal/service/auth"
	storage "github.com/MGomed/auth/internal/storage"
	auth_cache "github.com/MGomed/auth/internal/storage/cache/auth"
	msg_bus_auth "github.com/MGomed/auth/internal/storage/message_bus/auth"
	auth_repo "github.com/MGomed/auth/internal/storage/repository/auth"
	kafka "github.com/MGomed/auth/pkg/kafka"
	kafka_consumer "github.com/MGomed/auth/pkg/kafka/consumer"
	kafka_producer "github.com/MGomed/auth/pkg/kafka/producer"
	cache "github.com/MGomed/common/pkg/client/cache"
	redis "github.com/MGomed/common/pkg/client/cache/redis"
	db "github.com/MGomed/common/pkg/client/db"
	pg "github.com/MGomed/common/pkg/client/db/pg"
	transaction "github.com/MGomed/common/pkg/client/db/transaction"
	closer "github.com/MGomed/common/pkg/closer"
	logger "github.com/MGomed/common/pkg/logger"
	go_redis "github.com/gomodule/redigo/redis"
)

type serviceProvider struct {
	logger *log.Logger

	pgConfig      config.PgConfig
	redisConfig   config.RedisConfig
	grpcConfig    config.GRPCConfig
	httpConfig    config.HTTPConfig
	swaggerConfig config.SwaggerConfig
	kafkaConfig   config.KafkaConfig

	dbc         db.Client
	redisClient cache.RedisClient
	txMgr       db.TxManager
	producer    kafka.Producer

	cache  storage.Cache
	repo   storage.Repository
	msgBus storage.MessageBus

	service service.Service

	userAPI *auth_api.UserAPI
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

// GRPCConfig init/get grpc config
func (p *serviceProvider) GRPCConfig() config.GRPCConfig {
	if p.grpcConfig == nil {
		cfg, err := env_config.NewGRPCConfig()
		if err != nil {
			log.Fatalf("failed to create grpc config: %v", err)
		}

		p.grpcConfig = cfg
	}

	return p.grpcConfig
}

// HTTPConfig init/get http config
func (p *serviceProvider) HTTPConfig() config.HTTPConfig {
	if p.httpConfig == nil {
		cfg, err := env_config.NewHTTPConfig()
		if err != nil {
			log.Fatalf("failed to create http config: %v", err)
		}

		p.httpConfig = cfg
	}

	return p.httpConfig
}

// SwaggerConfig init/get swagger config
func (p *serviceProvider) SwaggerConfig() config.SwaggerConfig {
	if p.swaggerConfig == nil {
		cfg, err := env_config.NewSwaggerConfig()
		if err != nil {
			log.Fatalf("failed to create swagger config: %v", err)
		}

		p.swaggerConfig = cfg
	}

	return p.swaggerConfig
}

// KafkaConfig init/get kafka config
func (p *serviceProvider) KafkaConfig() config.KafkaConfig {
	if p.kafkaConfig == nil {
		cfg, err := env_config.NewKafkaConfig()
		if err != nil {
			log.Fatalf("failed to create kafka config: %v", err)
		}

		p.kafkaConfig = cfg
	}

	return p.kafkaConfig
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

// Producer init/get kafka Producer
func (p *serviceProvider) Producer(_ context.Context) kafka.Producer {
	if p.producer == nil {
		producer, err := kafka_producer.NewProducer(p.Logger(), p.KafkaConfig().Brokers())
		if err != nil {
			log.Fatalf("failed to create kafka producer: %e", err)
		}

		p.producer = producer

		closer.Add(p.producer.Close)
	}

	return p.producer
}

// Consumer init/get kafka Consumer
func (p *serviceProvider) Consumer() kafka.Consumer {
	group, err := sarama.NewConsumerGroup(
		p.KafkaConfig().Brokers(),
		p.KafkaConfig().GroupID(),
		p.KafkaConfig().Config(),
	)
	if err != nil {
		log.Fatalf("failed to create consumer group: %v", err)
	}

	handler := kafka_consumer.NewGroupHandler(p.Logger())

	return kafka_consumer.NewConsumer(p.Logger(), group, handler)
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

// MessageBus init/get MessageBus
func (p *serviceProvider) MessageBus(ctx context.Context) storage.MessageBus {
	if p.msgBus == nil {
		p.msgBus = msg_bus_auth.NewMessageBus(p.Producer(ctx))
	}

	return p.msgBus
}

// Service init/get Service(usecases)
func (p *serviceProvider) Service(ctx context.Context) service.Service {
	if p.service == nil {
		p.service = auth_service.NewService(
			p.Logger(),
			p.Repository(ctx),
			p.Cache(ctx),
			p.TxManager(ctx),
			p.MessageBus(ctx),
		)
	}

	return p.service
}

// UserAPI init/get UserAPI(grpc implementation)
func (p *serviceProvider) UserAPI(ctx context.Context) *auth_api.UserAPI {
	if p.userAPI == nil {
		p.userAPI = auth_api.NewUserAPI(p.Logger(), p.Service(ctx))
	}

	return p.userAPI
}
