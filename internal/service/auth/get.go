package auth

import (
	"context"

	service_model "github.com/MGomed/auth/internal/model"
)

// Get gets user by id
func (s *service) Get(ctx context.Context, id int64) (*service_model.User, error) {
	user, err := s.repo.GetUser(ctx, id)
	if err != nil {
		s.logger.Printf("Failed to get user with id - %v from database: %v", id, err)

		return nil, err
	}

	s.logger.Printf("Successfully got user!")

	return user, nil
}
