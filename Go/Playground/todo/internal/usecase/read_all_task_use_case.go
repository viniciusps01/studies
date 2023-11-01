package usecase

import "github.com/viniciusps01/internal/infra/repository"

type ReadAllTaskInputDTO struct{}

type ReadAllTaskOutput struct {
	ID          int64
	Description string
	Done        bool
}

type ReadAllTaskOutputDTO struct {
	Tasks []ReadAllTaskOutput
}

type ReadAllTaskUseCase struct {
	taskRepository repository.ITaskRepository
}

func NewReadAllTaskUseCase(r repository.ITaskRepository) ReadAllTaskUseCase {
	return ReadAllTaskUseCase{
		taskRepository: r,
	}
}

func (u ReadAllTaskUseCase) Exec(ReadAllTaskInputDTO) (*ReadAllTaskOutputDTO, error) {
	tasks, err := u.taskRepository.ReadAll()

	if err != nil {
		return nil, err
	}

	out := &ReadAllTaskOutputDTO{
		Tasks: []ReadAllTaskOutput{},
	}

	for _, t := range tasks {
		out.Tasks = append(out.Tasks, ReadAllTaskOutput{
			ID:          t.ID,
			Description: t.Description,
			Done:        t.Done,
		})
	}

	return out, nil
}
