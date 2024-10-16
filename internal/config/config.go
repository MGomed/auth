package config

// GRPCConfig is grpc config interface
type GRPCConfig interface {
	Address() string
}

// PgConfig is postgres config interface
type PgConfig interface {
	DSN() string
}

// LoggerConfig is logger config interface
type LoggerConfig interface {
	OutDir() string
}

// Config is common config
type Config struct {
	GRPCConfig
	PgConfig
	LoggerConfig
}

// NewConfig is Config struct constructor
func NewConfig(grpcConfig GRPCConfig, pgConfig PgConfig, loggerConfig LoggerConfig) *Config {
	return &Config{
		GRPCConfig:   grpcConfig,
		PgConfig:     pgConfig,
		LoggerConfig: loggerConfig,
	}
}
