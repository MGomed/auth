package accessapiimpl

import (
	"log"

	"github.com/MGomed/auth/internal/service"
	access_api "github.com/MGomed/auth/pkg/access_api"
)

// AccessAPI implements AuthAPI grpc server
type AccessAPI struct {
	access_api.UnimplementedAccessAPIServer

	logger *log.Logger

	accessSecretKey []byte

	service service.AccessService
}

// NewAccessAPI is api struct constructor
func NewAccessAPI(logger *log.Logger, secretKey []byte, service service.AccessService) *AccessAPI {
	return &AccessAPI{
		logger:          logger,
		accessSecretKey: secretKey,
		service:         service,
	}
}
