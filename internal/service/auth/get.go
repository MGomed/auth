package auth

import (
	"context"
	"errors"

	service_model "github.com/MGomed/auth/internal/model"
	cache_errors "github.com/MGomed/auth/internal/storage/cache/errors"
)

// Get gets user by id
func (s *service) Get(ctx context.Context, id int64) (*service_model.UserInfo, error) {
	user, err := s.cache.GetUser(ctx, id)
	if err == nil {
		return user, nil
	} else if !errors.Is(err, cache_errors.ErrUserNotPresent) {
		return nil, err
	}

	user, err = s.repo.GetUser(ctx, id)
	if err != nil {
		s.logger.Printf("Failed to get user with id - %v from database: %v", id, err)

		return nil, err
	}

	s.logger.Printf("Successfully got user!")

	return user, nil
}
