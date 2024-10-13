package env_config

import (
	"fmt"
	"os"
)

const loggerEnvOutDirName = "LOG_OUT_DIR"

type loggerConfig struct {
	logOutDir string
}

func NewLoggerConfig() (*loggerConfig, error) {
	dir := os.Getenv(loggerEnvOutDirName)
	if len(dir) == 0 {
		return nil, fmt.Errorf("%w: %v", errEnvNotFound, loggerEnvOutDirName)
	}

	return &loggerConfig{
		logOutDir: dir,
	}, nil
}

func (c *loggerConfig) OutDir() string {
	return c.logOutDir
}
