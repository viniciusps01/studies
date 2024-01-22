package usecase

import (
	repository "github.com/viniciusps01/todo/internal/feature/task/repository"
)

type DeleteTaskInputDTO struct {
	ID     int64  `json:"id"`
	UserID string `json:"user_id"`
}

type DeleteTaskUseCase struct {
	taskRepository repository.ITaskRepository
}

func NewDeleteTaskUseCase(r repository.ITaskRepository) DeleteTaskUseCase {
	return DeleteTaskUseCase{
		taskRepository: r,
	}
}

func (u DeleteTaskUseCase) Exec(i DeleteTaskInputDTO) error {
	err := u.taskRepository.Delete(i.ID, i.UserID)

	if err != nil {
		return err
	}

	return nil
}
