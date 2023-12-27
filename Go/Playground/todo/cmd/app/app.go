package app

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/viniciusps01/internal/config"
	"github.com/viniciusps01/internal/environment"
	auth_ds "github.com/viniciusps01/internal/feature/auth/data_source"
	auth_repo "github.com/viniciusps01/internal/feature/auth/repository"
	task_ds "github.com/viniciusps01/internal/feature/task/data_source"
	task_repo "github.com/viniciusps01/internal/feature/task/repository"
)

func New() *config.AppConfig {
	env := environment.Load()

	DB, err := sql.Open("pgx", env.DBpath)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	defer DB.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	conn, err := DB.Conn(ctx)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	ds := &config.DataSourceProvider{
		TaskDataSource: task_ds.NewTaskDataSourcePostgres(conn),
		AuthDataSource: auth_ds.NewAuthDataSourcePostgres(conn),
	}

	r := &config.RepositoryProvider{
		TaskRepository: task_repo.NewTaskRepository(ds.TaskDataSource),
		AuthRepository: auth_repo.NewAuthRepository(ds.AuthDataSource),
	}

	app := &config.AppConfig{
		RepositoryProvider: r,
		Env:                env,
		Conn:               conn,
	}

	return app
}
