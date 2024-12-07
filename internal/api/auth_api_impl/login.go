package authapiimpl

import (
	"context"

	auth_api "github.com/MGomed/auth/pkg/auth_api"
)

// Login logs in user
func (api *AuthAPI) Login(ctx context.Context, req *auth_api.LoginRequest) (*auth_api.LoginResponse, error) {
	token, err := api.service.Login(ctx, req.Email, req.Password, api.refreshSecretKey, api.refreshTokenExpiration)
	if err != nil {
		return nil, err
	}

	return &auth_api.LoginResponse{
		RefreshToken: token,
	}, nil
}
