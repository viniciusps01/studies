package usecase

import (
	"github.com/viniciusps01/todo/internal/feature/auth/repository"
)

type ReadUserInputDTO struct {
	ID string
}

type ReadUserOutputDTO struct {
	ID        string `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type ReadUserUseCase struct {
	repository repository.IAuthRepository
}

func NewReadUserUseCase(r repository.IAuthRepository) ReadUserUseCase {
	return ReadUserUseCase{
		repository: r,
	}
}

func (u ReadUserUseCase) Exec(i ReadUserInputDTO) (*ReadUserOutputDTO, error) {
	res, err := u.repository.Read(i.ID)

	if err != nil {
		return nil, err
	}

	out := ReadUserOutputDTO{
		ID:        res.ID,
		Email:     res.Email,
		FirstName: res.FirstName,
		LastName:  res.LastName,
	}

	return &out, nil
}
