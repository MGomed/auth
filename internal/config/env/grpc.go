package env_config

import (
	"fmt"
	"net"
	"os"
)

const (
	grpcHostName = "GRPC_HOST"
	grpcPortName = "GRPC_PORT"
)

type grpcConfig struct {
	host string
	port string
}

func NewGRPCConfig() (*grpcConfig, error) {
	host := os.Getenv(grpcHostName)
	if len(host) == 0 {
		return nil, fmt.Errorf("%w: %v", errEnvNotFound, grpcHostName)
	}

	port := os.Getenv(grpcPortName)
	if len(port) == 0 {
		return nil, fmt.Errorf("%w: %v", errEnvNotFound, grpcPortName)
	}

	return &grpcConfig{
		host: host,
		port: port,
	}, nil
}

func (c *grpcConfig) Address() string {
	return net.JoinHostPort(c.host, c.port)
}
