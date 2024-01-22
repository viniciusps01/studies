package infra

import (
	"context"
	"time"
)

type ICache interface {
	Put(ctx context.Context, key string, value []byte, expiration time.Duration) error
	Get(ctx context.Context, key string) (*string, error)
	Close() error
}
