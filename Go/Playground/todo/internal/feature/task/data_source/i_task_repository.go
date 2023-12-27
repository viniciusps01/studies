package data_source

import "github.com/viniciusps01/internal/feature/task/entity"

type ITaskDataSource interface {
	Create(task entity.Task) (int64, error)
	Read(ID int64, userID string) (*entity.Task, error)
	Update(task entity.Task) error
	Delete(ID int64, userID string) error
	ReadAll(userID string, limit, offset *int) ([]*entity.Task, error)
}
