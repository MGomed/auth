package env_config

import (
	"errors"

	godotenv "github.com/joho/godotenv"
)

var errEnvNotFound = errors.New("environment not found")

func Load(path string) error {
	return godotenv.Load(path)
}
