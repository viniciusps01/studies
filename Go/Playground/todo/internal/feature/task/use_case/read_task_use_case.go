package usecase

import repository "github.com/viniciusps01/todo/internal/feature/task/repository"

type ReadTaskInputDTO struct {
	ID     int64  `json:"id"`
	UserID string `json:"user_id"`
}

type ReadTaskOutputDto struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
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
	t, err := u.taskRepository.Read(i.ID, i.UserID)

	if err != nil {
		return nil, err
	}

	out := &ReadTaskOutputDto{
		ID:          t.ID,
		Title:       t.Title,
		Description: t.Description,
		Done:        t.Done,
	}

	return out, nil
}
