package cache

import (
	"context"
	"time"
)

//go:generate mockgen -destination=./mocks/cache_mock.go -package=mocks -source=interfaces.go

// RedisClient wraps some redis API calls
type RedisClient interface {
	HashSet(ctx context.Context, key string, values interface{}) error
	Set(ctx context.Context, key string, value interface{}) error
	HGetAll(ctx context.Context, key string, dest interface{}) error
	Get(ctx context.Context, key string) (interface{}, error)
	Del(ctx context.Context, key string) error
	Expire(ctx context.Context, key string, expiration time.Duration) error
	Ping(ctx context.Context) error
}
