package auth

import (
	"context"

	errors "github.com/MGomed/auth/internal/api/errors"
	converters "github.com/MGomed/auth/internal/converters"
	user_api "github.com/MGomed/auth/pkg/user_api"
)

// Create creates new user
func (s *API) Create(ctx context.Context, req *user_api.CreateRequest) (*user_api.CreateResponse, error) {
	if err := validateName(req.User.Name); err != nil {
		return nil, err
	}

	if err := validateEmail(req.User.Email); err != nil {
		return nil, err
	}

	if req.User.Password != req.User.PasswordConfirm {
		return nil, errors.ErrPasswordMismatch
	}

	id, err := s.service.Create(ctx, converters.ToUserCreateFromAPI(req.User))
	if err != nil {
		return nil, err
	}

	return &user_api.CreateResponse{Id: id}, nil
}
