package env_config

import (
	"fmt"
	"net"
	"os"

	consts "github.com/MGomed/auth/consts"
	config_errors "github.com/MGomed/auth/internal/config/errors"
)

type grpcConfig struct {
	host string
	port string
}

// NewGRPCConfig is grpcConfig struct constructor
func NewGRPCConfig() (*grpcConfig, error) {
	host := os.Getenv(consts.GRPCServerHostEnv)
	if len(host) == 0 {
		return nil, fmt.Errorf("%w: %v", config_errors.ErrEnvNotFound, consts.GRPCServerHostEnv)
	}

	port := os.Getenv(consts.GRPCServerPortEnv)
	if len(port) == 0 {
		return nil, fmt.Errorf("%w: %v", config_errors.ErrEnvNotFound, consts.GRPCServerPortEnv)
	}

	return &grpcConfig{
		host: host,
		port: port,
	}, nil
}

// Address returns grpc ip address
func (c *grpcConfig) Address() string {
	return net.JoinHostPort(c.host, c.port)
}
