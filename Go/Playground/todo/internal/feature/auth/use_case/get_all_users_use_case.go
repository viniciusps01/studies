package usecase

import (
	"github.com/viniciusps01/internal/feature/auth/entity"
	"github.com/viniciusps01/internal/feature/auth/repository"
)

type ReadAllUsersInputDTO struct {
	Limit  *int
	Offset *int
}

type ReadAllUsersOutputDTO struct {
	ID          string              `json:"id"`
	FirstName   string              `json:"first_name"`
	LastName    string              `json:"last_name"`
	Email       string              `json:"email"`
	Permissions []entity.Permission `json:"permissions"`
}

type ReadAllUsersUseCase struct {
	repository repository.IAuthRepository
}

func NewReadAllUsersUseCase(r repository.IAuthRepository) ReadAllUsersUseCase {
	return ReadAllUsersUseCase{
		repository: r,
	}
}

func (u ReadAllUsersUseCase) Exec(i ReadAllUsersInputDTO) (*[]ReadAllUsersOutputDTO, error) {
	res, err := u.repository.ReadAllUsers(i.Limit, i.Offset)

	if err != nil {
		return nil, err
	}

	users := make([]ReadAllUsersOutputDTO, len(*res))

	for i, u := range *res {
		users[i] = ReadAllUsersOutputDTO{
			ID:          u.ID,
			Email:       u.Email,
			FirstName:   u.FirstName,
			LastName:    u.LastName,
			Permissions: u.Permissions,
		}
	}

	return &users, nil
}
