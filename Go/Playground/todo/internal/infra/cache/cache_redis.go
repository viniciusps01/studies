package infra

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisCache struct {
	redis *redis.Client
}

func NewRedisCache(url string) ICache {
	opt, err := redis.ParseURL(url)
	if err != nil {
		panic(err)
	}

	client := redis.NewClient(opt)

	return RedisCache{
		redis: client,
	}
}

func (c RedisCache) Put(
	ctx context.Context,
	key string,
	value []byte,
	expiration time.Duration,
) error {
	status := c.redis.Set(ctx, key, value, expiration)

	if err := status.Err(); err != nil {
		return err
	}

	return nil
}

func (c RedisCache) Get(ctx context.Context, key string) (*string, error) {
	cmd := c.redis.Get(ctx, key)

	if err := cmd.Err(); err != nil {
		return nil, err
	}

	value, err := cmd.Result()

	if err != nil {
		return nil, err
	}

	return &value, nil
}

func (c RedisCache) Close() error {
	return c.redis.Close()
}
