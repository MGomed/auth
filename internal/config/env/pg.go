package env_config

import (
	"fmt"
	"os"
)

const (
	pgEnvHost     = "PG_HOST"
	pgEnvPort     = "PG_PORT"
	pgEnvDBName   = "PG_DATABASE_NAME"
	pgEnvUser     = "PG_USER"
	pgEnvPassword = "PG_PASSWORD"
)

type pgConfig struct {
	host     string
	port     string
	dbName   string
	user     string
	password string
}

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

func (c *pgConfig) DSN() string {
	return fmt.Sprintf("host=%v port=%v dbname=%v user=%v password=%v sslmode=disable",
		c.host, c.port, c.dbName, c.user, c.password)
}
