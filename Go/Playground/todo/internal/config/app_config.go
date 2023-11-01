package config

import (
	"database/sql"

	"github.com/viniciusps01/internal/environment"
)

type AppConfig struct {
	RepositoryProvider *RepositoryProvider
	Env                *environment.Environment
	Conn               *sql.Conn
}

func NewAppConfig(
	r *RepositoryProvider,
	env *environment.Environment,
	db *sql.Conn,
) AppConfig {
	return AppConfig{
		RepositoryProvider: r,
		Env:                env,
		Conn:               db,
	}
}
