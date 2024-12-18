package userservice

import (
	"context"
	"errors"

	cache_errors "github.com/MGomed/auth/internal/storage/cache/errors"
	msg_bus_converters "github.com/MGomed/auth/internal/storage/message_bus/converters"
)

// Delete removes user by id
func (s *service) Delete(ctx context.Context, id int64) error {
	_, err := s.repo.DeleteUser(ctx, id)
	if err != nil {
		s.logger.Error("Failed to delete user with id - %v from database: %v", id, err)

		return err
	}

	if err := s.cache.DeleteUser(ctx, id); err != nil {
		if !errors.Is(err, cache_errors.ErrUserNotPresent) {
			return err
		}
	}

	s.logger.Debug("Successfully deleted user: %v", id)

	msg, err := msg_bus_converters.ToMessageTypeFromID(id)
	if err != nil {
		return err
	}

	if err := s.msgBus.SendMessage(ctx, msg); err != nil {
		return err
	}

	return nil
}
