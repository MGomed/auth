package userservice

import (
	"context"

	service_model "github.com/MGomed/auth/internal/model"
)

// Update modifies user information
func (s *service) Update(ctx context.Context, id int64, user *service_model.UserUpdate) error {
	var userInfo *service_model.UserInfo

	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var errTx error
		_, errTx = s.repo.UpdateUser(ctx, id, user)
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
		return err
	}

	if err := s.cache.CreateUser(ctx, id, userInfo); err != nil {
		return err
	}

	s.logger.Printf("Successfully updated user with id: %v", id)

	return nil
}
