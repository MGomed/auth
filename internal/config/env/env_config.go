package env_config

import (
	godotenv "github.com/joho/godotenv"
)

// Load setup envs from file
func Load(path string) error {
	return godotenv.Load(path)
}
