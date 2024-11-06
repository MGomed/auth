package auth

import (
	"fmt"
	"log"
	"strings"

	errors "github.com/MGomed/auth/internal/api/errors"
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

func validateName(name string) error {
	if n := len([]rune(name)); n < 2 || n > 32 {
		return errors.ErrNameLenInvalid
	}

	return nil
}

func validateEmail(email string) error {
	atInd := strings.Index(email, "@")
	dotInd := strings.LastIndex(email, ".")

	if atInd == -1 || dotInd == -1 {
		return fmt.Errorf("%w: %v", errors.ErrEmailInvalid, email)
	}

	if atInd > dotInd {
		return fmt.Errorf("%w: %v", errors.ErrEmailInvalid, email)
	}

	return nil
}
