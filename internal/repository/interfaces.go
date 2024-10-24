package repository

import (
	"context"

	service_model "github.com/MGomed/auth/internal/model"
)

// Repository declaired interface for database communication
type Repository interface {
	CreateUser(ctx context.Context, user *service_model.User) (int64, error)
	GetUser(ctx context.Context, id int64) (*service_model.User, error)
	UpdateUser(ctx context.Context, user *service_model.User) (int64, error)
	DeleteUser(ctx context.Context, id int64) (int64, error)
}
