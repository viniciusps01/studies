package repository

import "github.com/viniciusps01/todo/internal/feature/task/entity"

type ITaskRepository interface {
	Create(task entity.Task) (int64, error)
	Read(ID int64, userID string) (*entity.Task, error)
	Update(task entity.Task) error
	Delete(ID int64, userID string) error
	ReadAll(userID string, limit, offset *int) ([]*entity.Task, error)
}
