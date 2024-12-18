package authapiimpl

import (
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

	service service.AuthService
}

// NewAuthAPI is api struct constructor
func NewAuthAPI(
	refreshTokenExpirationTime, accessTokenExpirationTime time.Duration,
	refreshSecretKey, accessSecretKey []byte,
	service service.AuthService,
) *AuthAPI {
	return &AuthAPI{
		service:                service,
		refreshTokenExpiration: refreshTokenExpirationTime,
		accessTokenExpiration:  accessTokenExpirationTime,
		refreshSecretKey:       refreshSecretKey,
		accessSecretKey:        accessSecretKey,
	}
}
