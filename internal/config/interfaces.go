package config

import "time"

// APIConfig is grpc config interface
type APIConfig interface {
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
