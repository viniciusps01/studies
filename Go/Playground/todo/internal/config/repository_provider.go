package config

import (
	user_repo "github.com/viniciusps01/internal/feature/auth/repository"
	task_repo "github.com/viniciusps01/internal/feature/task/repository"
)

type RepositoryProvider struct {
	TaskRepository task_repo.ITaskRepository
	AuthRepository user_repo.IAuthRepository
}
