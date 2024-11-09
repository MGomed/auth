package auth

import (
	"context"

	converters "github.com/MGomed/auth/internal/converters"
	user_api "github.com/MGomed/auth/pkg/user_api"
)

// Get gets user by id
func (s *UserAPI) Get(ctx context.Context, req *user_api.GetRequest) (*user_api.GetResponse, error) {
	resp, err := s.service.Get(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &user_api.GetResponse{
		User: converters.ToUserInfoFromService(resp),
	}, nil
}
