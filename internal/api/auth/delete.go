package auth

import (
	"context"

	empty "github.com/golang/protobuf/ptypes/empty"

	consts "github.com/MGomed/auth/consts"
	user_api "github.com/MGomed/auth/pkg/user_api"
)

// Delete removes user by id
func (s *UserAPI) Delete(ctx context.Context, req *user_api.DeleteRequest) (*empty.Empty, error) {
	ctx, cancel := context.WithTimeout(ctx, consts.ContextTimeout)
	defer cancel()

	err := s.service.Delete(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &empty.Empty{}, nil
}
