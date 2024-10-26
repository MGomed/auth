package auth

import (
	"context"
	"errors"

	pgx "github.com/jackc/pgx/v4"
)

// Delete removes user by id
func (s *service) Delete(ctx context.Context, id int64) error {
	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var errTx error
		_, errTx = s.repo.DeleteUser(ctx, id)
		if errTx != nil {
			s.logger.Printf("Failed to delete user with id - %v from database: %v", id, errTx)

			return errTx
		}

		_, errTx = s.repo.GetUser(ctx, id)
		if !errors.Is(errTx, pgx.ErrNoRows) {
			return errTx
		}

		s.logger.Printf("Successfully deleted user: %v", id)

		return nil
	})

	return err
}
