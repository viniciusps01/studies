package repository

import (
	"github.com/viniciusps01/todo/internal/feature/task/data_source"
	"github.com/viniciusps01/todo/internal/feature/task/entity"
)

type TaskRepository struct {
	ds data_source.ITaskDataSource
}

func NewTaskRepository(ds data_source.ITaskDataSource) *TaskRepository {
	r := &TaskRepository{
		ds: ds,
	}

	return r
}

func (r *TaskRepository) Create(t entity.Task) (int64, error) {
	return r.ds.Create(t)
}

func (r *TaskRepository) Read(ID int64, userID string) (*entity.Task, error) {
	return r.ds.Read(ID, userID)
}

func (r *TaskRepository) Update(t entity.Task) error {
	return r.ds.Update(t)
}

func (r *TaskRepository) Delete(ID int64, userID string) error {
	return r.ds.Delete(ID, userID)
}

func (r *TaskRepository) ReadAll(userID string, limit, offset *int) ([]*entity.Task, error) {
	return r.ds.ReadAll(userID, limit, offset)
}
