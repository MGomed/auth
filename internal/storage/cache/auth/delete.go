package auth

import (
	"context"
)

// DeleteUser deletes user from cache
func (c *cacher) DeleteUser(ctx context.Context, id int64) error {
	return c.client.Del(ctx, constructKey(id))
}
