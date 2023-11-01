package usecase

import "github.com/viniciusps01/internal/infra/repository"

type ReadTaskInputDTO struct {
	ID int64 `json:"id"`
}

type ReadTaskOutputDto struct {
	ID          int64  `json:"id"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

type ReadTaskUseCase struct {
	taskRepository repository.ITaskRepository
}

func NewReadTaskUseCase(r repository.ITaskRepository) ReadTaskUseCase {
	return ReadTaskUseCase{
		taskRepository: r,
	}
}

func (u ReadTaskUseCase) Exec(i ReadTaskInputDTO) (*ReadTaskOutputDto, error) {
	t, err := u.taskRepository.Read(i.ID)

	if err != nil {
		return nil, err
	}

	out := &ReadTaskOutputDto{
		ID:          t.ID,
		Description: t.Description,
		Done:        t.Done,
	}

	return out, nil
}
