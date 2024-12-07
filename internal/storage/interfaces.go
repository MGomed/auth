package repository

import (
	"context"

	service_model "github.com/MGomed/auth/internal/model"
	msg_bus_model "github.com/MGomed/auth/internal/storage/message_bus/model"
)

//go:generate mockgen -destination=./mocks/storage_mock.go -package=mocks -source=interfaces.go

// Repository declaired interface for database communication
type Repository interface {
	CreateUser(ctx context.Context, user *service_model.UserCreate) (int64, error)
	GetUser(ctx context.Context, id int64) (*service_model.UserInfo, error)
	GetUserByEmail(ctx context.Context, email string) (*service_model.UserInfo, error)
	UpdateUser(ctx context.Context, id int64, user *service_model.UserUpdate) (int64, error)
	DeleteUser(ctx context.Context, id int64) (int64, error)
}

// Cache declaired interface for redis cache communication
type Cache interface {
	CreateUser(ctx context.Context, id int64, user *service_model.UserInfo) error
	GetUser(ctx context.Context, id int64) (*service_model.UserInfo, error)
	DeleteUser(ctx context.Context, id int64) error
}

// MessageBus declaired interface for sending messages to message brokers
type MessageBus interface {
	SendMessage(ctx context.Context, msg *msg_bus_model.Message) error
}
