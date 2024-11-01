package auth

import (
	"context"

	empty "github.com/golang/protobuf/ptypes/empty"

	converters "github.com/MGomed/auth/internal/converters"
	user_api "github.com/MGomed/auth/pkg/user_api"
)

// Update modifies user information
func (s *API) Update(ctx context.Context, req *user_api.UpdateRequest) (*empty.Empty, error) {
	if err := validateName(req.User.Name.GetValue()); err != nil {
		return nil, err
	}

	err := s.service.Update(ctx, converters.ToUserUpdateFromAPI(req.User))
	if err != nil {
		return nil, err
	}

	return &empty.Empty{}, nil
}
