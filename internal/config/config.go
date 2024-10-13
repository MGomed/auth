package config

type GRPCConfig interface {
	Address() string
}

type PgConfig interface {
	DSN() string
}

type LoggerConfig interface {
	OutDir() string
}

type Config struct {
	GRPCConfig
	PgConfig
	LoggerConfig
}

func NewConfig(grpcConfig GRPCConfig, pgConfig PgConfig, loggerConfig LoggerConfig) *Config {
	return &Config{
		GRPCConfig:   grpcConfig,
		PgConfig:     pgConfig,
		LoggerConfig: loggerConfig,
	}
}
