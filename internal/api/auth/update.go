package auth

import (
	"context"

	empty "github.com/golang/protobuf/ptypes/empty"

	service_model "github.com/MGomed/auth/internal/model"
	user_api "github.com/MGomed/auth/pkg/user_api"
)

// Update modifies user information
func (s *API) Update(ctx context.Context, req *user_api.UpdateRequest) (*empty.Empty, error) {
	err := s.service.Update(ctx, service_model.ToUserUpdateFromAPI(req.User))
	if err != nil {
		return nil, err
	}

	return &empty.Empty{}, nil
}
