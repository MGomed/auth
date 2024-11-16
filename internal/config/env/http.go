package env_config

import (
	"fmt"
	"net"
	"os"

	consts "github.com/MGomed/auth/consts"
	config_errors "github.com/MGomed/auth/internal/config/errors"
)

type httpConfig struct {
	host string
	port string
}

// NewHTTPConfig is httpConfig struct constructor
func NewHTTPConfig() (*httpConfig, error) {
	host := os.Getenv(consts.HTTPServerHostEnv)
	if len(host) == 0 {
		return nil, fmt.Errorf("%w: %v", config_errors.ErrEnvNotFound, consts.HTTPServerHostEnv)
	}

	port := os.Getenv(consts.HTTPServerPortEnv)
	if len(port) == 0 {
		return nil, fmt.Errorf("%w: %v", config_errors.ErrEnvNotFound, consts.HTTPServerPortEnv)
	}

	return &httpConfig{
		host: host,
		port: port,
	}, nil
}

// Address returns http ip address
func (c *httpConfig) Address() string {
	return net.JoinHostPort(c.host, c.port)
}
