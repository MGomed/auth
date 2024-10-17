package env_config

import (
	"fmt"
	"os"
)

const loggerEnvOutDirName = "LOG_OUT_DIR"

type loggerConfig struct {
	logOutDir string
}

// NewLoggerConfig is loggerConfig struct constructor
func NewLoggerConfig() (*loggerConfig, error) {
	dir := os.Getenv(loggerEnvOutDirName)
	if len(dir) == 0 {
		return nil, fmt.Errorf("%w: %v", errEnvNotFound, loggerEnvOutDirName)
	}

	return &loggerConfig{
		logOutDir: dir,
	}, nil
}

// OutDir return logs out director
func (c *loggerConfig) OutDir() string {
	return c.logOutDir
}
