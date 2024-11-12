package auth

import (
	"context"

	service_model "github.com/MGomed/auth/internal/model"
	cache_converters "github.com/MGomed/auth/internal/storage/cache/converters"
)

// CreateUser creates user in cache
func (c *cacher) CreateUser(ctx context.Context, id int64, user *service_model.UserInfo) error {
	return c.client.HashSet(ctx, constructKey(id), cache_converters.ToUserInfoFromService(user))
}
