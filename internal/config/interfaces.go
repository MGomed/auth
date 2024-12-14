package config

import (
	"time"

	sarama "github.com/IBM/sarama"
)

// GRPCConfig is grpc config interface
type GRPCConfig interface {
	Address() string
}

// HTTPConfig is http config interface
type HTTPConfig interface {
	Address() string
}

// SwaggerConfig is swagger config interface
type SwaggerConfig interface {
	Address() string
}

// PrometheusConfig is http config interface
type PrometheusConfig interface {
	Address() string
}

// PgConfig is postgres config interface
type PgConfig interface {
	DSN() string
}

// RedisConfig is redis config interface
type RedisConfig interface {
	Address() string
	ConnectionTimeout() time.Duration
	MaxIdle() int
	IdleTimeout() time.Duration
}

// KafkaConfig is kafka config interface
type KafkaConfig interface {
	Brokers() []string
	GroupID() string
	Config() *sarama.Config
}

// JWTConfig is jwt config interface
type JWTConfig interface {
	GetRefreshTokenExpirationTimeMin() time.Duration
	GetAccessTokenExpirationTimeMin() time.Duration
}
