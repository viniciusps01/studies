package usecase

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/viniciusps01/todo/internal/feature/auth/repository"
	"github.com/viniciusps01/todo/pkg/apperrors"
	"github.com/viniciusps01/todo/pkg/security"
)

type AuthenticateInputDTO struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
	ExpiresIn int    `json:"expires_in"`
}

type AuthenticateOutputDTO struct {
	Token string `json:"token"`
}

type AuthenticateUseCase struct {
	repository repository.IAuthRepository
}

func NewAuthenticateUseCase(r repository.IAuthRepository) AuthenticateUseCase {
	return AuthenticateUseCase{
		repository: r,
	}
}

func (u AuthenticateUseCase) Exec(i AuthenticateInputDTO) (*AuthenticateOutputDTO, error) {
	user, err := u.repository.ReadUserByEmail(i.Email)

	if err != nil {
		return nil, err
	}

	err = security.VerifyPassword(i.Password, user.Password)

	if err != nil {
		return nil, apperrors.AuthenticationError{
			Message: "authentication failure",
		}
	}

	now := time.Now()
	expireAt := jwt.NewNumericDate(now.Add(time.Second * 60 * 15))
	issuedAt := jwt.NewNumericDate(now)

	claims := security.Claims{
		UserID: user.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: expireAt,
			IssuedAt:  issuedAt,
		},
	}

	jwt, err := security.GenerateToken(claims)

	if err != nil {
		return nil, err
	}

	out := AuthenticateOutputDTO{
		Token: *jwt,
	}

	return &out, nil
}
