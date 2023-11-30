package dbrepo

import (
	"app/internals/config"
	"database/sql"
)

type PostgresDbRepo struct {
	App *config.AppConfig
	DB  *sql.DB
}

func NewPostgresRepo(con *sql.DB, app *config.AppConfig) DatabaseRepo {
	return PostgresDbRepo{
		App: app,
		DB:  con,
	}
}
