package env_config

import (
	"fmt"
	"os"
)

const (
	pgEnvHost     = "DB_HOST"
	pgEnvPort     = "DB_PORT"
	pgEnvDBName   = "POSTGRES_DB"
	pgEnvUser     = "POSTGRES_USER"
	pgEnvPassword = "POSTGRES_PASSWORD" //nolint: gosec
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
	host := os.Getenv(pgEnvHost)
	if len(host) == 0 {
		return nil, fmt.Errorf("%w: %v", errEnvNotFound, pgEnvHost)
	}

	port := os.Getenv(pgEnvPort)
	if len(port) == 0 {
		return nil, fmt.Errorf("%w: %v", errEnvNotFound, pgEnvPort)
	}

	dbName := os.Getenv(pgEnvDBName)
	if len(dbName) == 0 {
		return nil, fmt.Errorf("%w: %v", errEnvNotFound, pgEnvDBName)
	}

	user := os.Getenv(pgEnvUser)
	if len(user) == 0 {
		return nil, fmt.Errorf("%w: %v", errEnvNotFound, pgEnvUser)
	}

	password := os.Getenv(pgEnvPassword)
	if len(password) == 0 {
		return nil, fmt.Errorf("%w: %v", errEnvNotFound, pgEnvPassword)
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
