package env_config

import (
	"fmt"
	"net"
	"os"

	consts "github.com/MGomed/auth/consts"
	config_errors "github.com/MGomed/auth/internal/config/errors"
)

type prometheusConfig struct {
	host string
	port string
}

// NewPrometheusConfig is httpConfig struct constructor
func NewPrometheusConfig() (*prometheusConfig, error) {
	host := os.Getenv(consts.PrometheusServerHostEnv)
	if len(host) == 0 {
		return nil, fmt.Errorf("%w: %v", config_errors.ErrEnvNotFound, consts.PrometheusServerHostEnv)
	}

	port := os.Getenv(consts.PrometheusServerPortEnv)
	if len(port) == 0 {
		return nil, fmt.Errorf("%w: %v", config_errors.ErrEnvNotFound, consts.PrometheusServerPortEnv)
	}

	return &prometheusConfig{
		host: host,
		port: port,
	}, nil
}

// Address returns http ip address
func (c *prometheusConfig) Address() string {
	return net.JoinHostPort(c.host, c.port)
}
