package auth

import (
	"context"

	service_model "github.com/MGomed/auth/internal/model"
	user_api "github.com/MGomed/auth/pkg/user_api"
)

// Create creates new user
func (s *API) Create(ctx context.Context, req *user_api.CreateRequest) (*user_api.CreateResponse, error) {
	resp, err := s.service.Create(ctx, service_model.ToUserCreateFromAPI(req.User))
	if err != nil {
		return nil, err
	}

	return &user_api.CreateResponse{Id: resp}, nil
}
