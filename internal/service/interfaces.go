package service

import (
	"context"
	"time"

	service_model "github.com/MGomed/auth/internal/model"
)

//go:generate mockgen -destination=./mocks/service_mock.go -package=mocks -source=interfaces.go

// UserService interface of user_api usecase
type UserService interface {
	Create(ctx context.Context, user *service_model.UserCreate) (int64, error)
	Get(ctx context.Context, id int64) (*service_model.UserInfo, error)
	Update(ctx context.Context, id int64, user *service_model.UserUpdate) error
	Delete(ctx context.Context, id int64) error
}

// AuthService interface of auth_api usecase
type AuthService interface {
	Login(ctx context.Context, email, password string, secretKey []byte, duration time.Duration) (string, error)
	GetRefreshToken(refreshToken string, secretKey []byte, duration time.Duration) (string, error)
	GetAccessToken(refreshToken string, refreshSecretKey, accessSecretKey []byte, duration time.Duration) (string, error)
}

// AccessService interface of access_api usecase
type AccessService interface {
	Check(route string, accessToken string, secretKey []byte) error
}
