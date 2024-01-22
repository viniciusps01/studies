package config

import infra "github.com/viniciusps01/todo/internal/infra/cache"

type CacheProvider struct {
	RedisCache infra.ICache
}
