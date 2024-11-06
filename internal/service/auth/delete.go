package auth

import (
	"context"
)

// Delete removes user by id
func (s *service) Delete(ctx context.Context, id int64) error {
	_, err := s.repo.DeleteUser(ctx, id)
	if err != nil {
		s.logger.Printf("Failed to delete user with id - %v from database: %v", id, err)

		return err
	}

	if err := s.cache.DeleteUser(ctx, id); err != nil {
		return err
	}

	s.logger.Printf("Successfully deleted user: %v", id)

	return nil
}
