package auth

import (
	"context"

	service_model "github.com/MGomed/auth/internal/model"
)

// Create creates new user
func (s *service) Create(ctx context.Context, user *service_model.User) (int64, error) {
	id, err := s.repo.CreateUser(ctx, user)
	if err != nil {
		s.logger.Printf("Failed to add user in database: %v", err)

		return 0, err
	}

	s.logger.Printf("Successfully added user with id: %v", id)

	return id, nil
}
