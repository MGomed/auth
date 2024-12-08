package accessapiimpl

import (
	service "github.com/MGomed/auth/internal/service"
	access_api "github.com/MGomed/auth/pkg/access_api"
)

// AccessAPI implements AuthAPI grpc server
type AccessAPI struct {
	access_api.UnimplementedAccessAPIServer

	accessSecretKey []byte

	service service.AccessService
}

// NewAccessAPI is api struct constructor
func NewAccessAPI(secretKey []byte, service service.AccessService) *AccessAPI {
	return &AccessAPI{
		accessSecretKey: secretKey,
		service:         service,
	}
}
