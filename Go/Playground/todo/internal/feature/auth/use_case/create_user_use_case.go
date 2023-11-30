package usecase

import (
	"strings"

	"github.com/viniciusps01/internal/feature/auth/entity"
	"github.com/viniciusps01/internal/feature/auth/repository"

	"github.com/viniciusps01/pkg/apperrors"
	"github.com/viniciusps01/pkg/security"
	"github.com/viniciusps01/pkg/validator"
)

type CreateUserInputDTO struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type CreateUserOutputDTO struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

type CreateUserUseCase struct {
	repository repository.IAuthRepository
}

func NewCreateUserUseCase(r repository.IAuthRepository) CreateUserUseCase {
	return CreateUserUseCase{
		repository: r,
	}
}

func (i CreateUserInputDTO) validate() *validator.ValidationResult {
	r := validator.Validate(
		validator.ValidateRequired("first_name", i.FirstName),
		validator.ValidateRange("first_name", i.FirstName, 1, 50),

		validator.ValidateRequired("last_name", i.LastName),
		validator.ValidateRange("last_name", i.LastName, 1, 128),

		validator.ValidateRequired("email", i.Email),
		validator.ValidateEmail("email", i.Email),

		validator.ValidateRequired("password", i.Password),
	)

	return r
}

func (u CreateUserUseCase) Exec(i CreateUserInputDTO) (*CreateUserOutputDTO, error) {
	if v := i.validate(); !v.IsValid() {

		return nil, apperrors.BadRequestError{
			Message: (strings.Join(v.Validations, ", ")),
		}
	}

	pwdHash, err := security.EncryptPassword(i.Password)

	if err != nil {
		return nil, err
	}

	in := entity.NewUser(i.Email, i.FirstName, i.LastName, *pwdHash, nil)
	res, err := u.repository.Create(in)

	if err != nil {
		return nil, err
	}

	out := &CreateUserOutputDTO{
		ID:        res.ID,
		FirstName: res.FirstName,
		LastName:  res.LastName,
		Email:     res.Email,
	}

	return out, nil
}
