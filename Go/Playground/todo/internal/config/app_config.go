package config

import (
	"database/sql"

	"github.com/viniciusps01/todo/internal/environment"
	infra "github.com/viniciusps01/todo/internal/infra/cache"
)

type AppConfig struct {
	RepositoryProvider *RepositoryProvider
	Env                *environment.Environment
	Conn               *sql.Conn
	RedisCache         infra.ICache
}

func NewAppConfig(
	r *RepositoryProvider,
	env *environment.Environment,
	db *sql.Conn,
	redisCache infra.ICache,
) AppConfig {
	return AppConfig{
		RepositoryProvider: r,
		Env:                env,
		Conn:               db,
		RedisCache:         redisCache,
	}
}
