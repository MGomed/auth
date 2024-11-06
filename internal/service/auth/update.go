package auth

import (
	"context"

	service_model "github.com/MGomed/auth/internal/model"
)

// Update modifies user information
func (s *service) Update(ctx context.Context, user *service_model.UserUpdate) error {
	_, err := s.repo.UpdateUser(ctx, user)
	if err != nil {
		s.logger.Printf("Failed to update user from database: %v", err)

		return err
	}

	s.logger.Printf("Successfully updated user with id: %v", user.ID)

	return nil
}
