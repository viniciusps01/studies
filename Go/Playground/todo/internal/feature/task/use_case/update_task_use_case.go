package usecase

import (
	"strings"

	"github.com/viniciusps01/todo/internal/feature/task/repository"
	"github.com/viniciusps01/todo/pkg/apperrors"
	"github.com/viniciusps01/todo/pkg/validator"
)

type UpdateTaskInputDTO struct {
	ID          int64   `json:"id"`
	UserID      string  `json:"user_id"`
	Title       *string `json:"title"`
	Description *string `json:"description"`
	Done        *bool   `json:"done"`
}

type UpdateTaskOutputDTO struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

type UpdateTaskUseCase struct {
	taskRepository repository.ITaskRepository
}

func NewUpdateTaskUseCase(r repository.ITaskRepository) UpdateTaskUseCase {
	return UpdateTaskUseCase{
		taskRepository: r,
	}
}

func (u UpdateTaskUseCase) Exec(i UpdateTaskInputDTO) (*UpdateTaskOutputDTO, error) {
	t, err := u.taskRepository.Read(i.ID, i.UserID)

	if err != nil {
		return nil, err
	}

	if i.Title != nil {
		t.Title = *i.Title
	}

	if i.Description != nil {
		t.Description = *i.Description
	}

	if i.Done != nil {
		t.Done = *i.Done
	}

	r := validator.Validate(
		validator.ValidateRequired("title", t.Title),
		validator.ValidateRange("title", t.Title, 1, 128),
	)

	if !r.IsValid() {
		err := apperrors.BadRequestError{
			Message: strings.Join(r.Validations, ", "),
		}
		return nil, err
	}

	err = u.taskRepository.Update(*t)

	if err != nil {
		return nil, err
	}

	out := UpdateTaskOutputDTO{
		ID:          t.ID,
		Title:       t.Title,
		Description: t.Description,
		Done:        t.Done,
	}

	return &out, nil
}
