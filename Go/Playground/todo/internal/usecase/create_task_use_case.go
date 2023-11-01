package usecase

import (
	"strings"

	"github.com/viniciusps01/internal/entity"
	"github.com/viniciusps01/internal/infra/repository"
	"github.com/viniciusps01/pkg/apperrors"
	"github.com/viniciusps01/pkg/validator"
)

type CreateTaskInputDTO struct {
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

type CreateTaskOutputDTO struct {
	ID          int64  `json:"id"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

type CreateTaskUseCase struct {
	Repository repository.ITaskRepository
}

func NewCreateTask(r repository.ITaskRepository) CreateTaskUseCase {
	return CreateTaskUseCase{
		Repository: r,
	}
}

func (u *CreateTaskUseCase) Exec(i CreateTaskInputDTO) (*CreateTaskOutputDTO, error) {
	if err := validate(i); err != nil {
		return nil, err
	}

	ID, err := u.Repository.Create(*entity.NewTask(i.Description, i.Done))

	if err != nil {
		return nil, err
	}

	out := &CreateTaskOutputDTO{ID: ID, Description: i.Description, Done: i.Done}
	return out, nil
}

func validate(i CreateTaskInputDTO) error {
	r := validator.Validate(
		validator.ValidateRequired("description", i.Description),
		validator.ValidateRange("description", i.Description, 1, 128),
	)

	if r.IsValid() {
		return nil
	}

	err := apperrors.ValidationError{Message: strings.Join(r.Validations, ", ")}

	return err
}
