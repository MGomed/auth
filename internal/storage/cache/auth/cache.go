package auth

import (
	"fmt"

	cache "github.com/MGomed/common/pkg/client/cache"
)

type cacher struct {
	client cache.RedisClient
}

// NewCacher is cacher struct constructor
func NewCacher(client cache.RedisClient) *cacher {
	return &cacher{
		client: client,
	}
}

func constructKey(id int64) string {
	return fmt.Sprintf("user:id:%v", id)
}
