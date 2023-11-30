package usecase

import (
	"github.com/viniciusps01/internal/feature/auth/entity"
	"github.com/viniciusps01/internal/feature/auth/repository"
	"github.com/viniciusps01/pkg/apperrors"
	"github.com/viniciusps01/pkg/security"
)

type CheckAuthInputDTO struct {
	Token string
}

type AuthUser struct {
	ID          string
	Email       string
	FirstName   string
	LastName    string
	Permissions []entity.Permission
	Claims      security.Claims
}

type CheckAuthUseCase struct {
	repository repository.IAuthRepository
}

func NewCheckAuthUseCase(r repository.IAuthRepository) CheckAuthUseCase {
	return CheckAuthUseCase{
		repository: r,
	}
}

func (u AuthUser) HasPermission(ID int) bool {
	for _, p := range u.Permissions {
		if p.ID == ID {
			return true
		}
	}

	return false
}

func (u CheckAuthUseCase) Exec(i CheckAuthInputDTO) (*AuthUser, error) {
	claims, err := security.ValidateToken(i.Token)

	if err != nil {
		return nil, apperrors.AuthenticationError{
			Message: "Authentication failure",
		}

	}

	res, err := u.repository.Read(claims.UserID)

	if err != nil {
		return nil, err
	}

	out := AuthUser{
		ID:          res.ID,
		Email:       res.Email,
		FirstName:   res.FirstName,
		LastName:    res.LastName,
		Permissions: res.Permissions,
		Claims:      *claims,
	}

	return &out, nil
}
