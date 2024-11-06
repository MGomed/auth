package auth

import (
	"context"
	"fmt"

	service_model "github.com/MGomed/auth/internal/model"
	cache_converters "github.com/MGomed/auth/internal/storage/cache/converters"
	cache_errors "github.com/MGomed/auth/internal/storage/cache/errors"
	cache_model "github.com/MGomed/auth/internal/storage/cache/model"
)

func (c *cacher) GetUser(ctx context.Context, id int64) (*service_model.UserInfo, error) {
	var user = &cache_model.UserInfo{}

	if err := c.client.HGetAll(ctx, constructKey(id), user); err != nil {
		return nil, err
	}

	if user.CreatedAt.IsZero() {
		return nil, fmt.Errorf("%w: %v", cache_errors.ErrUserNotPresent, id)
	}

	return cache_converters.ToUserInfoFromCache(user), nil
}