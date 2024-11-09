package auth

import (
	"context"

	service_model "github.com/MGomed/auth/internal/model"
)

// Create creates new user
func (s *service) Create(ctx context.Context, user *service_model.UserCreate) (int64, error) {
	var (
		id       int64
		userInfo *service_model.UserInfo
	)

	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var errTx error
		id, errTx = s.repo.CreateUser(ctx, user)
		if errTx != nil {
			s.logger.Printf("Failed to add user in database: %v", errTx)

			return errTx
		}

		userInfo, errTx = s.repo.GetUser(ctx, id)
		if errTx != nil {
			return errTx
		}

		s.logger.Printf("Successfully added user with id: %v", id)

		return nil
	})
	if err != nil {
		return 0, err
	}

	if err := s.cache.CreateUser(ctx, id, userInfo); err != nil {
		return 0, err
	}

	return id, nil
}
