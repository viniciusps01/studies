package app

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/viniciusps01/internal/config"
	"github.com/viniciusps01/internal/environment"
	auth_repo "github.com/viniciusps01/internal/feature/auth/repository"
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

	r := &config.RepositoryProvider{
		TaskRepository: task_repo.NewTaskRepositoryPostgres(conn),
		AuthRepository: auth_repo.NewAuthRepositoryPostgres(conn),
	}

	app := &config.AppConfig{
		RepositoryProvider: r,
		Env:                env,
		Conn:               conn,
	}

	return app
}
