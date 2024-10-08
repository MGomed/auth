package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

const configPath = "config/auth.conf"

var errReadConf = errors.New("couldn't read config")

// GRPC defines struct with grpc host and port
type GRPC struct {
	Host string `json:"host"`
	Port uint32 `json:"port"`
}

// Config defines auth service configuration
type Config struct {
	GRPC      GRPC   `json:"grpc"`
	OutLogDir string `json:"outLogDir"`
}

// GetConfig reads config from configPath and unmarshal it to Config
func GetConfig() (*Config, error) {
	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("%w from %v: %w", errReadConf, configPath, err)
	}

	var config = &Config{}
	if err := json.Unmarshal(data, config); err != nil {
		return nil, fmt.Errorf("%w: %w", errReadConf, err)
	}

	return config, nil
}
