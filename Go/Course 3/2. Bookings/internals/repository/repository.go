package repository

import (
	"app/internals/config"
	"app/internals/repository/dbrepo"
)

type Repository struct {
	App *config.AppConfig
	DB  *dbrepo.DatabaseRepo
}
