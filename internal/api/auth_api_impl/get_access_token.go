package authapiimpl

import (
	"context"

	auth_api "github.com/MGomed/auth/pkg/auth_api"
)

// GetAccessToken creates new access JWT token
func (api *AuthAPI) GetAccessToken(_ context.Context, req *auth_api.GetAccessTokenRequest) (*auth_api.GetAccessTokenResponse, error) {
	token, err := api.service.GetAccessToken(
		req.RefreshToken,
		api.refreshSecretKey,
		api.accessSecretKey,
		api.accessTokenExpiration,
	)
	if err != nil {
		return nil, err
	}

	return &auth_api.GetAccessTokenResponse{
		AccessToken: token,
	}, nil
}
