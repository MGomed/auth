package auth

import (
	"log"

	service "github.com/MGomed/auth/internal/service"
	user_api "github.com/MGomed/auth/pkg/user_api"
)

// API implements UserAPI grpc server
type API struct {
	user_api.UnimplementedUserAPIServer

	logger  *log.Logger
	service service.Service
}

// NewAPI is api struct constructor
func NewAPI(logger *log.Logger, service service.Service) *API {
	return &API{
		logger:  logger,
		service: service,
	}
}
