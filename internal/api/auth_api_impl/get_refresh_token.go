package authapiimpl

import (
	"context"

	auth_api "github.com/MGomed/auth/pkg/auth_api"
)

// GetRefreshToken creates new refresh JWT token
func (api *AuthAPI) GetRefreshToken(_ context.Context, req *auth_api.GetRefreshTokenRequest) (*auth_api.GetRefreshTokenResponse, error) {
	token, err := api.service.GetRefreshToken(req.RefreshToken, api.refreshSecretKey, api.refreshTokenExpiration)
	if err != nil {
		return nil, err
	}

	return &auth_api.GetRefreshTokenResponse{
		RefreshToken: token,
	}, nil
}
