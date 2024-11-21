package env_config

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"time"

	consts "github.com/MGomed/auth/consts"
	config_errors "github.com/MGomed/auth/internal/config/errors"
)

type redisConfig struct {
	host string
	port string

	connectionTimeoutSecond time.Duration
	maxIdle                 int
	idleTimeoutSecond       time.Duration
}

// NewRedisConfig is redisConfig struct constructor
func NewRedisConfig() (*redisConfig, error) {
	host := os.Getenv(consts.RedisHostEnv)
	if len(host) == 0 {
		return nil, fmt.Errorf("%w: %v", config_errors.ErrEnvNotFound, consts.RedisHostEnv)
	}

	port := os.Getenv(consts.RedisPortEnv)
	if len(port) == 0 {
		return nil, fmt.Errorf("%w: %v", config_errors.ErrEnvNotFound, consts.RedisPortEnv)
	}

	connectionTimeoutSecStr := os.Getenv(consts.RedisConnectionTimeoutSecEnv)
	if len(connectionTimeoutSecStr) == 0 {
		return nil, fmt.Errorf("%w: %v", config_errors.ErrEnvNotFound, consts.RedisConnectionTimeoutSecEnv)
	}

	connectionTimeoutSec, err := strconv.ParseInt(connectionTimeoutSecStr, 10, 64)
	if err != nil {
		return nil, err
	}

	maxIdleStr := os.Getenv(consts.RedisMaxIdleEnv)
	if len(maxIdleStr) == 0 {
		return nil, fmt.Errorf("%w: %v", config_errors.ErrEnvNotFound, consts.RedisMaxIdleEnv)
	}

	maxIdle, err := strconv.Atoi(maxIdleStr)
	if err != nil {
		return nil, err
	}

	idleTimeoutSecStr := os.Getenv(consts.RedisIdleTimeoutSecEnv)
	if len(idleTimeoutSecStr) == 0 {
		return nil, fmt.Errorf("%w: %v", config_errors.ErrEnvNotFound, consts.RedisIdleTimeoutSecEnv)
	}

	idleTimeoutSec, err := strconv.ParseInt(idleTimeoutSecStr, 10, 64)
	if err != nil {
		return nil, err
	}

	return &redisConfig{
		host:                    host,
		port:                    port,
		connectionTimeoutSecond: time.Duration(connectionTimeoutSec) * time.Second,
		maxIdle:                 maxIdle,
		idleTimeoutSecond:       time.Duration(idleTimeoutSec) * time.Second,
	}, nil
}

func (cfg *redisConfig) Address() string {
	return net.JoinHostPort(cfg.host, cfg.port)
}

func (cfg *redisConfig) ConnectionTimeout() time.Duration {
	return cfg.connectionTimeoutSecond
}

func (cfg *redisConfig) MaxIdle() int {
	return cfg.maxIdle
}

func (cfg *redisConfig) IdleTimeout() time.Duration {
	return cfg.idleTimeoutSecond
}
