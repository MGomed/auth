package app

import (
	"context"
	"log"
	"os"

	sarama "github.com/IBM/sarama"

	consts "github.com/MGomed/auth/consts"
	access_api_impl "github.com/MGomed/auth/internal/api/access_api_impl"
	auth_api_impl "github.com/MGomed/auth/internal/api/auth_api_impl"
	user_api_impl "github.com/MGomed/auth/internal/api/user_api_impl"
	config "github.com/MGomed/auth/internal/config"
	env_config "github.com/MGomed/auth/internal/config/env"
	service "github.com/MGomed/auth/internal/service"
	accessservice "github.com/MGomed/auth/internal/service/access_service"
	auth_service "github.com/MGomed/auth/internal/service/auth_service"
	user_service "github.com/MGomed/auth/internal/service/user_service"
	storage "github.com/MGomed/auth/internal/storage"
	auth_cache "github.com/MGomed/auth/internal/storage/cache/auth"
	msg_bus_auth "github.com/MGomed/auth/internal/storage/message_bus/auth"
	auth_repo "github.com/MGomed/auth/internal/storage/repository/auth"
	kafka "github.com/MGomed/auth/pkg/kafka"
	kafka_consumer "github.com/MGomed/auth/pkg/kafka/consumer"
	kafka_producer "github.com/MGomed/auth/pkg/kafka/producer"
	cache "github.com/MGomed/common/client/cache"
	redis "github.com/MGomed/common/client/cache/redis"
	db "github.com/MGomed/common/client/db"
	pg "github.com/MGomed/common/client/db/pg"
	transaction "github.com/MGomed/common/client/db/transaction"
	closer "github.com/MGomed/common/closer"
	logger "github.com/MGomed/common/logger"
	go_redis "github.com/gomodule/redigo/redis"
)

type serviceProvider struct {
	refreshSecretKey []byte
	accessSecretKey  []byte

	logger *log.Logger

	pgConfig      config.PgConfig
	redisConfig   config.RedisConfig
	grpcConfig    config.GRPCConfig
	httpConfig    config.HTTPConfig
	swaggerConfig config.SwaggerConfig
	kafkaConfig   config.KafkaConfig
	jwtConfig     config.JWTConfig

	dbc         db.Client
	redisClient cache.RedisClient
	txMgr       db.TxManager
	producer    kafka.Producer

	cache  storage.Cache
	repo   storage.Repository
	msgBus storage.MessageBus

	userService   service.UserService
	authService   service.AuthService
	accessService service.AccessService

	userAPI   *user_api_impl.UserAPI
	authAPI   *auth_api_impl.AuthAPI
	accessAPI *access_api_impl.AccessAPI
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

// RefreshSecretKey init/get secret key refresh token
func (p *serviceProvider) RefreshSecretKey() []byte {
	if len(p.refreshSecretKey) == 0 {
		key, err := os.ReadFile(consts.RefreshSecretKeyPath)
		if err != nil {
			log.Fatalf("couldn't read refresh secret key: %v", err)
		}

		p.refreshSecretKey = key
	}

	return p.refreshSecretKey
}

// AccessSecretKey init/get secret key access token
func (p *serviceProvider) AccessSecretKey() []byte {
	if len(p.accessSecretKey) == 0 {
		key, err := os.ReadFile(consts.AccessSecretKeyPath)
		if err != nil {
			log.Fatalf("couldn't read access secret key: %v", err)
		}

		p.accessSecretKey = key
	}

	return p.accessSecretKey
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

// JWTConfig init/get jwt config
func (p *serviceProvider) JWTConfig() config.JWTConfig {
	if p.jwtConfig == nil {
		cfg, err := env_config.NewJWTConfig()
		if err != nil {
			log.Fatalf("failed to create jwt config: %v", err)
		}

		p.jwtConfig = cfg
	}

	return p.jwtConfig
}

// Logger init/get logger
func (p *serviceProvider) Logger() *log.Logger {
	if p.logger == nil {
		p.logger = logger.InitLogger(consts.ServiceName)
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
			log.Fatalf("failed to create kafka producer: %v", err)
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

// UserService init/get UserService(usecases)
func (p *serviceProvider) UserService(ctx context.Context) service.UserService {
	if p.userService == nil {
		p.userService = user_service.NewUserService(
			p.Logger(),
			p.Repository(ctx),
			p.Cache(ctx),
			p.TxManager(ctx),
			p.MessageBus(ctx),
		)
	}

	return p.userService
}

// AuthService init/get AuthService(usecases)
func (p *serviceProvider) AuthService(ctx context.Context) service.AuthService {
	if p.authService == nil {
		p.authService = auth_service.NewAuthService(
			p.Logger(),
			p.Repository(ctx),
		)
	}

	return p.authService
}

// AccessService init/get AccessService(usecases)
func (p *serviceProvider) AccessService() service.AccessService {
	if p.accessService == nil {
		p.accessService = accessservice.NewAccessService(p.Logger())
	}

	return p.accessService
}

// UserAPI init/get UserAPI(grpc implementation)
func (p *serviceProvider) UserAPI(ctx context.Context) *user_api_impl.UserAPI {
	if p.userAPI == nil {
		p.userAPI = user_api_impl.NewUserAPI(p.Logger(), p.UserService(ctx))
	}

	return p.userAPI
}

// AuthAPI init/get AuthAPI(grpc implementation)
func (p *serviceProvider) AuthAPI(ctx context.Context) *auth_api_impl.AuthAPI {
	if p.authAPI == nil {
		p.authAPI = auth_api_impl.NewAuthAPI(
			p.Logger(),
			p.JWTConfig().GetRefreshTokenExpirationTimeMin(),
			p.JWTConfig().GetAccessTokenExpirationTimeMin(),
			p.RefreshSecretKey(),
			p.AccessSecretKey(),
			p.AuthService(ctx),
		)
	}

	return p.authAPI
}

// AccessAPI init/get AccessAPI(grpc implementation)
func (p *serviceProvider) AccessAPI() *access_api_impl.AccessAPI {
	if p.accessAPI == nil {
		p.accessAPI = access_api_impl.NewAccessAPI(
			p.Logger(),
			p.AccessSecretKey(),
			p.AccessService(),
		)
	}

	return p.accessAPI
}
