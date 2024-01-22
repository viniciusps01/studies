package usecase

import "github.com/viniciusps01/todo/internal/feature/auth/repository"

type DeleteUserInputDTO struct {
	ID string
}

type DeleteUserOutputDTO struct{}

type DeleteUserUseCase struct {
	repository repository.IAuthRepository
}

func NewDeleteUserUseCase(r repository.IAuthRepository) DeleteUserUseCase {
	return DeleteUserUseCase{
		repository: r,
	}
}

func (u DeleteUserUseCase) Exec(i DeleteUserInputDTO) (*DeleteUserOutputDTO, error) {
	err := u.repository.Delete(i.ID)

	if err != nil {
		return nil, err
	}

	return &DeleteUserOutputDTO{}, nil
}
