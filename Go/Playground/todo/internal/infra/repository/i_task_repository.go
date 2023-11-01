package repository

import "github.com/viniciusps01/internal/entity"

type ITaskRepository interface {
	Create(task entity.Task) (int64, error)
	Read(ID int64) (*entity.Task, error)
	Update(task entity.Task) error
	Delete(ID int64) error
	ReadAll() ([]*entity.Task, error)
}
