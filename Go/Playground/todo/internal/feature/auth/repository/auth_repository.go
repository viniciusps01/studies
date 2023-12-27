package repository

import (
	"github.com/viniciusps01/internal/feature/auth/data_source"
	"github.com/viniciusps01/internal/feature/auth/entity"
)

type AuthRepository struct {
	ds data_source.IAuthDataSource
}

func NewAuthRepository(ds data_source.IAuthDataSource) AuthRepository {
	return AuthRepository{
		ds: ds,
	}
}

func (r AuthRepository) Create(u entity.User) (*entity.User, error) {
	return r.ds.Create(u)
}

func (r AuthRepository) Read(ID string) (*entity.User, error) {
	return r.ds.Read(ID)
}

func (r AuthRepository) ReadUserByEmail(email string) (*entity.User, error) {
	return r.ds.ReadUserByEmail(email)
}

func (r AuthRepository) Update(user entity.User) (*entity.User, error) {
	return r.ds.Update(user)
}

func (r AuthRepository) Delete(ID string) error {
	return r.ds.Delete(ID)
}

func (r AuthRepository) ReadAllUsers(limit, offset *int) (*[]entity.User, error) {
	return r.ds.ReadAllUsers(limit, offset)
}
