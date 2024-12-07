package userapiimpl

import (
	"log"

	service "github.com/MGomed/auth/internal/service"
	user_api "github.com/MGomed/auth/pkg/user_api"
)

// UserAPI implements UserAPI grpc server
type UserAPI struct {
	user_api.UnimplementedUserAPIServer

	logger  *log.Logger
	service service.UserService
}

// NewUserAPI is api struct constructor
func NewUserAPI(logger *log.Logger, service service.UserService) *UserAPI {
	return &UserAPI{
		logger:  logger,
		service: service,
	}
}
