package config

import "github.com/viniciusps01/internal/infra/repository"

type RepositoryProvider struct {
	TaskRepository repository.ITaskRepository
}
