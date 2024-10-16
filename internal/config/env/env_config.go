package env_config

import (
	"errors"

	godotenv "github.com/joho/godotenv"
)

var errEnvNotFound = errors.New("environment not found")

// Load setup envs from file
func Load(path string) error {
	return godotenv.Load(path)
}
