package authapiimpl

import (
	"log"
	"time"

	service "github.com/MGomed/auth/internal/service"
	auth_api "github.com/MGomed/auth/pkg/auth_api"
)

// AuthAPI implements AuthAPI grpc server
type AuthAPI struct {
	auth_api.UnimplementedAuthAPIServer

	refreshSecretKey []byte
	accessSecretKey  []byte

	refreshTokenExpiration time.Duration
	accessTokenExpiration  time.Duration

	logger  *log.Logger
	service service.AuthService
}

// NewAuthAPI is api struct constructor
func NewAuthAPI(
	logger *log.Logger,
	refreshTokenExpirationTime, accessTokenExpirationTime time.Duration,
	refreshSecretKey, accessSecretKey []byte,
	service service.AuthService,
) *AuthAPI {
	return &AuthAPI{
		logger:                 logger,
		service:                service,
		refreshTokenExpiration: refreshTokenExpirationTime,
		accessTokenExpiration:  accessTokenExpirationTime,
		refreshSecretKey:       refreshSecretKey,
		accessSecretKey:        accessSecretKey,
	}
}
