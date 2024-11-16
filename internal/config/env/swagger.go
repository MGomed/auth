package env_config

import (
	"fmt"
	"net"
	"os"

	consts "github.com/MGomed/auth/consts"
	config_errors "github.com/MGomed/auth/internal/config/errors"
)

type swaggerConfig struct {
	host string
	port string
}

// NewSwaggerConfig is swaggerConfig struct constructor
func NewSwaggerConfig() (*swaggerConfig, error) {
	host := os.Getenv(consts.SwaggerServerHostEnv)
	if len(host) == 0 {
		return nil, fmt.Errorf("%w: %v", config_errors.ErrEnvNotFound, consts.SwaggerServerHostEnv)
	}

	port := os.Getenv(consts.SwaggerServerPortEnv)
	if len(port) == 0 {
		return nil, fmt.Errorf("%w: %v", config_errors.ErrEnvNotFound, consts.SwaggerServerPortEnv)
	}

	return &swaggerConfig{
		host: host,
		port: port,
	}, nil
}

// Address returns http ip address
func (c *swaggerConfig) Address() string {
	return net.JoinHostPort(c.host, c.port)
}
