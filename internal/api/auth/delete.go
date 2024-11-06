package auth

import (
	"context"

	empty "github.com/golang/protobuf/ptypes/empty"

	user_api "github.com/MGomed/auth/pkg/user_api"
)

// Delete removes user by id
func (s *API) Delete(ctx context.Context, req *user_api.DeleteRequest) (*empty.Empty, error) {
	err := s.service.Delete(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &empty.Empty{}, nil
}
