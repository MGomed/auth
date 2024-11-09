package redis

import (
	"context"
	"log"
	"time"

	go_redis "github.com/gomodule/redigo/redis"
)

type handler func(ctx context.Context, conn go_redis.Conn) error

type client struct {
	logger      *log.Logger
	pool        *go_redis.Pool
	connTimeout time.Duration
}

// NewClient is client struct constructor
func NewClient(logger *log.Logger, pool *go_redis.Pool, connTimeout time.Duration) *client {
	return &client{
		logger:      logger,
		pool:        pool,
		connTimeout: connTimeout,
	}
}

// HashSet wraps HSET redis command
func (c *client) HashSet(ctx context.Context, key string, values interface{}) error {
	err := c.execute(ctx, func(_ context.Context, conn go_redis.Conn) error {
		_, err := conn.Do("HSET", go_redis.Args{key}.AddFlat(values)...)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

// Set wraps SET redis command
func (c *client) Set(ctx context.Context, key string, value interface{}) error {
	err := c.execute(ctx, func(_ context.Context, conn go_redis.Conn) error {
		_, err := conn.Do("SET", go_redis.Args{key}.Add(value)...)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

// HGetAll wraps HGETALL redis command
func (c *client) HGetAll(ctx context.Context, key string, dest interface{}) error {
	var values []interface{}
	err := c.execute(ctx, func(_ context.Context, conn go_redis.Conn) error {
		var errEx error
		values, errEx = go_redis.Values(conn.Do("HGETALL", key))
		if errEx != nil {
			return errEx
		}

		return nil
	})
	if err != nil {
		return err
	}

	if len(values) == 0 {
		return nil
	}

	return go_redis.ScanStruct(values, dest)
}

// Get wraps GET redis command
func (c *client) Get(ctx context.Context, key string) (interface{}, error) {
	var value interface{}
	err := c.execute(ctx, func(_ context.Context, conn go_redis.Conn) error {
		var errEx error
		value, errEx = conn.Do("GET", key)
		if errEx != nil {
			return errEx
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return value, nil
}

// Del wraps DEL redis command
func (c *client) Del(ctx context.Context, key string) error {
	err := c.execute(ctx, func(_ context.Context, conn go_redis.Conn) error {
		var errEx error
		_, errEx = conn.Do("DEL", key)
		if errEx != nil {
			return errEx
		}

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

// Expire wraps EXPIRE redis command
func (c *client) Expire(ctx context.Context, key string, expiration time.Duration) error {
	err := c.execute(ctx, func(_ context.Context, conn go_redis.Conn) error {
		_, err := conn.Do("EXPIRE", key, int(expiration.Seconds()))
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

// Ping wraps PING redis command
func (c *client) Ping(ctx context.Context) error {
	err := c.execute(ctx, func(_ context.Context, conn go_redis.Conn) error {
		_, err := conn.Do("PING")
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

func (c *client) execute(ctx context.Context, handler handler) error {
	conn, err := c.getConnect(ctx)
	if err != nil {
		return err
	}
	defer func() {
		err = conn.Close()
		if err != nil {
			c.logger.Printf("failed to close redis connection: %v\n", err)
		}
	}()

	err = handler(ctx, conn)
	if err != nil {
		return err
	}

	return nil
}

func (c *client) getConnect(ctx context.Context) (go_redis.Conn, error) {
	getConnTimeoutCtx, cancel := context.WithTimeout(ctx, c.connTimeout)
	defer cancel()

	conn, err := c.pool.GetContext(getConnTimeoutCtx)
	if err != nil {
		c.logger.Printf("failed to get redis connection: %v\n", err)

		_ = conn.Close()
		return nil, err
	}

	return conn, nil
}
