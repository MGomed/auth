package auth

import (
	"context"
	"errors"

	cache_errors "github.com/MGomed/auth/internal/storage/cache/errors"
)

// Delete removes user by id
func (s *service) Delete(ctx context.Context, id int64) error {
	_, err := s.repo.DeleteUser(ctx, id)
	if err != nil {
		s.logger.Printf("Failed to delete user with id - %v from database: %v", id, err)

		return err
	}

	if err := s.cache.DeleteUser(ctx, id); err != nil {
		if !errors.Is(err, cache_errors.ErrUserNotPresent) {
			return err
		}
	}

	s.logger.Printf("Successfully deleted user: %v", id)

	return nil
}
