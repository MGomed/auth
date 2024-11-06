package service

import (
	"context"

	service_model "github.com/MGomed/auth/internal/model"
)

//go:generate mockgen -destination=./mocks/service_mock.go -package=mocks -source=interfaces.go

// Service interface of user_api usecase
type Service interface {
	Create(ctx context.Context, user *service_model.UserCreate) (int64, error)
	Get(ctx context.Context, id int64) (*service_model.UserInfo, error)
	Update(ctx context.Context, user *service_model.UserUpdate) error
	Delete(ctx context.Context, id int64) error
}
