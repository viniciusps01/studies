package usecase

import (
	"strings"

	"github.com/viniciusps01/internal/infra/repository"
	"github.com/viniciusps01/pkg/apperrors"
	"github.com/viniciusps01/pkg/validator"
)

type UpdateTaskInputDTO struct {
	ID          int64   `json:"id"`
	Description *string `json:"description"`
	Done        *bool   `json:"done"`
}

type UpdateTaskOutputDTO struct {
	ID          int64  `json:"id"`
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
	t, err := u.taskRepository.Read(i.ID)

	if err != nil {
		return nil, err
	}

	if i.Description != nil {
		t.Description = *i.Description
	}

	if i.Done != nil {
		t.Done = *i.Done
	}

	r := validator.Validate(
		validator.ValidateRequired("description", t.Description),
		validator.ValidateRange("description", t.Description, 1, 128),
	)

	if !r.IsValid() {
		err := apperrors.ValidationError{
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
		Description: t.Description,
		Done:        t.Done,
	}

	return &out, nil
}
