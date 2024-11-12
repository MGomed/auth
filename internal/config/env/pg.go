package env_config

import (
	"fmt"
	"os"

	consts "github.com/MGomed/auth/consts"
	config_errors "github.com/MGomed/auth/internal/config/errors"
)

type pgConfig struct {
	host     string
	port     string
	dbName   string
	user     string
	password string
}

// NewPgConfig is pgConfig struct constructor
func NewPgConfig() (*pgConfig, error) {
	host := os.Getenv(consts.DBHostEnv)
	if len(host) == 0 {
		return nil, fmt.Errorf("%w: %v", config_errors.ErrEnvNotFound, consts.DBHostEnv)
	}

	port := os.Getenv(consts.DBPortEnv)
	if len(port) == 0 {
		return nil, fmt.Errorf("%w: %v", config_errors.ErrEnvNotFound, consts.DBPortEnv)
	}

	dbName := os.Getenv(consts.DBNameEnv)
	if len(dbName) == 0 {
		return nil, fmt.Errorf("%w: %v", config_errors.ErrEnvNotFound, consts.DBNameEnv)
	}

	user := os.Getenv(consts.DBUserEnv)
	if len(user) == 0 {
		return nil, fmt.Errorf("%w: %v", config_errors.ErrEnvNotFound, consts.DBUserEnv)
	}

	password := os.Getenv(consts.DBPasswordEnv)
	if len(password) == 0 {
		return nil, fmt.Errorf("%w: %v", config_errors.ErrEnvNotFound, consts.DBPasswordEnv)
	}

	return &pgConfig{
		host:     host,
		port:     port,
		dbName:   dbName,
		user:     user,
		password: password,
	}, nil
}

// DSN returns postgres connection dsn
func (c *pgConfig) DSN() string {
	return fmt.Sprintf("host=%v port=%v dbname=%v user=%v password=%v sslmode=disable",
		c.host, c.port, c.dbName, c.user, c.password)
}
