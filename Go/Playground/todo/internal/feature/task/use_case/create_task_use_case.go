package usecase

import (
	"strings"

	"github.com/viniciusps01/internal/feature/task/entity"
	repository "github.com/viniciusps01/internal/feature/task/repository"

	"github.com/viniciusps01/pkg/apperrors"
	"github.com/viniciusps01/pkg/validator"
)

type CreateTaskInputDTO struct {
	Title       string `json:"title"`
	UserID      string `json:"user_id"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

type CreateTaskOutputDTO struct {
	ID          int64   `json:"id"`
	Title       string  `json:"title"`
	Description *string `json:"description"`
	Done        bool    `json:"done"`
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

	ID, err := u.Repository.Create(*entity.NewTask(i.UserID, i.Title, i.Description, i.Done))

	if err != nil {
		return nil, err
	}

	out := &CreateTaskOutputDTO{ID: ID, Title: i.Title, Description: &i.Description, Done: i.Done}
	return out, nil
}

func validate(i CreateTaskInputDTO) error {
	r := validator.Validate(
		validator.ValidateRequired("title", i.Title),
		validator.ValidateRange("title", i.Title, 1, 128),
	)

	if r.IsValid() {
		return nil
	}

	err := apperrors.BadRequestError{Message: strings.Join(r.Validations, ", ")}

	return err
}
