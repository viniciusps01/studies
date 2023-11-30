package repository

import "github.com/viniciusps01/internal/feature/auth/entity"

type IAuthRepository interface {
	Create(user entity.User) (*entity.User, error)
	Read(ID string) (*entity.User, error)
	ReadAllUsers(limit, offset *int) (*[]entity.User, error)
	ReadUserByEmail(ID string) (*entity.User, error)
	Update(user entity.User) (*entity.User, error)
	Delete(ID string) error
}
