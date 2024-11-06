package auth

import (
	"context"
)

func (c *cacher) DeleteUser(ctx context.Context, id int64) error {
	return c.client.Del(ctx, constructKey(id))
}
