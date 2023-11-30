package usecase

import "github.com/viniciusps01/internal/feature/task/repository"

type ReadAllTaskInputDTO struct {
	UserID string `json:"user_id"`
	Limit  *int   `json:"limit"`
	Offset *int   `json:"offset"`
}

type ReadAllTaskOutput struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

type ReadAllTaskOutputDTO struct {
	Tasks []ReadAllTaskOutput `json:"tasks"`
}

type ReadAllTaskUseCase struct {
	taskRepository repository.ITaskRepository
}

func NewReadAllTaskUseCase(r repository.ITaskRepository) ReadAllTaskUseCase {
	return ReadAllTaskUseCase{
		taskRepository: r,
	}
}

func (u ReadAllTaskUseCase) Exec(i ReadAllTaskInputDTO) (*ReadAllTaskOutputDTO, error) {
	tasks, err := u.taskRepository.ReadAll(i.UserID, i.Limit, i.Offset)

	if err != nil {
		return nil, err
	}

	out := &ReadAllTaskOutputDTO{
		Tasks: []ReadAllTaskOutput{},
	}

	for _, t := range tasks {
		out.Tasks = append(out.Tasks, ReadAllTaskOutput{
			ID:          t.ID,
			Title:       t.Title,
			Description: t.Description,
			Done:        t.Done,
		})
	}

	return out, nil
}
