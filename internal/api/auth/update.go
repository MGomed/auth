package auth

import (
	"context"

	empty "github.com/golang/protobuf/ptypes/empty"

	consts "github.com/MGomed/auth/consts"
	converters "github.com/MGomed/auth/internal/converters"
	user_api "github.com/MGomed/auth/pkg/user_api"
)

// Update modifies user information
func (s *UserAPI) Update(ctx context.Context, req *user_api.UpdateRequest) (*empty.Empty, error) {
	ctx, cancel := context.WithTimeout(ctx, consts.ContextTimeout)
	defer cancel()

	err := s.service.Update(ctx, req.Id, converters.ToUserUpdateFromAPI(req.User))
	if err != nil {
		return nil, err
	}

	return &empty.Empty{}, nil
}
