package usecase

import "github.com/viniciusps01/internal/infra/repository"

type DeleteTaskInputDTO struct {
	ID int64 `json:"id"`
}

type DeleteTaskOutputDTO struct {
	ID int64 `json:"id"`
}

type DeleteTaskUseCase struct {
	taskRepository repository.ITaskRepository
}

func NewDeleteTaskUseCase(r repository.ITaskRepository) DeleteTaskUseCase {
	return DeleteTaskUseCase{
		taskRepository: r,
	}
}

func (u DeleteTaskUseCase) Exec(i DeleteTaskInputDTO) (*DeleteTaskOutputDTO, error) {
	err := u.taskRepository.Delete(i.ID)

	if err != nil {
		return nil, err
	}

	out := &DeleteTaskOutputDTO{
		ID: i.ID,
	}

	return out, nil
}
