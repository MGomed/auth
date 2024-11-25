package env_config

import (
	"fmt"
	"os"
	"strconv"
	"time"

	consts "github.com/MGomed/auth/consts"
	config_errors "github.com/MGomed/auth/internal/config/errors"
)

type jwtConfig struct {
	refreshTokenExpirationMin time.Duration
	accessTokenExpirationMin  time.Duration
}

// NewJWTConfig is a jwtConfig struct constructor
func NewJWTConfig() (*jwtConfig, error) {
	refreshTokenExpirationTimeMinStr := os.Getenv(consts.JWTRefreshTokenExpirationTimeMinEnv)
	if len(refreshTokenExpirationTimeMinStr) == 0 {
		return nil, fmt.Errorf("%w: %v", config_errors.ErrEnvNotFound, consts.JWTRefreshTokenExpirationTimeMinEnv)
	}

	refreshTokenExpirationTimeMin, err := strconv.Atoi(refreshTokenExpirationTimeMinStr)
	if err != nil {
		return nil, err
	}

	accessTokenExpirationTimeMinStr := os.Getenv(consts.JWTAccessTokenExpirationTimeMinEnv)
	if len(accessTokenExpirationTimeMinStr) == 0 {
		return nil, fmt.Errorf("%w: %v", config_errors.ErrEnvNotFound, consts.JWTAccessTokenExpirationTimeMinEnv)
	}

	accessTokenExpirationTimeMin, err := strconv.Atoi(accessTokenExpirationTimeMinStr)
	if err != nil {
		return nil, err
	}

	return &jwtConfig{
		refreshTokenExpirationMin: time.Duration(refreshTokenExpirationTimeMin * int(time.Minute)),
		accessTokenExpirationMin:  time.Duration(accessTokenExpirationTimeMin * int(time.Minute)),
	}, nil
}

// GetRefreshTokenExpirationTimeMin returns a refresh token's expiration time in minutes
func (c *jwtConfig) GetRefreshTokenExpirationTimeMin() time.Duration {
	return c.refreshTokenExpirationMin
}

// GetAccessTokenExpirationTimeMin returns a access token's expiration time in minutes
func (c *jwtConfig) GetAccessTokenExpirationTimeMin() time.Duration {
	return c.accessTokenExpirationMin
}
