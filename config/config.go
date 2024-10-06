package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

const configPath = "config/auth.conf"

var errReadConf = errors.New("couldn't read config")

type ConfigGRPC struct {
	Host string `json:"host"`
	Port uint32 `json:"port"`
}

type Config struct {
	GRPC      ConfigGRPC `json:"grpc"`
	OutLogDir string     `json:"outLogDir"`
}

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
