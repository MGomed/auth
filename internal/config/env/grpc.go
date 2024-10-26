package env_config

import (
	"fmt"
	"net"
	"os"
)

const (
	grpcHostName = "SERVER_HOST"
	grpcPortName = "SERVER_PORT"
)

type grpcConfig struct {
	host string
	port string
}

// NewAPIConfig is grpcConfig struct constructor
func NewAPIConfig() (*grpcConfig, error) {
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

// Address returns grpc ip address
func (c *grpcConfig) Address() string {
	return net.JoinHostPort(c.host, c.port)
}
