package config

import (
	user_ds "github.com/viniciusps01/todo/internal/feature/auth/data_source"
	task_ds "github.com/viniciusps01/todo/internal/feature/task/data_source"
)

type DataSourceProvider struct {
	TaskDataSource task_ds.ITaskDataSource
	AuthDataSource user_ds.IAuthDataSource
}
