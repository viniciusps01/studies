package repository

import (
	"context"
	"time"

	"github.com/viniciusps01/todo/internal/feature/auth/data_source"
	"github.com/viniciusps01/todo/internal/feature/auth/entity"
)

type AuthRepository struct {
	ds    data_source.IAuthDataSource
	cache data_source.IAuthCacheDataSource
}

func NewAuthRepository(
	ds data_source.IAuthDataSource,
	c data_source.IAuthCacheDataSource,
) AuthRepository {
	return AuthRepository{
		ds:    ds,
		cache: c,
	}
}

func (r AuthRepository) Create(u entity.User) (*entity.User, error) {
	return r.ds.Create(u)
}

func (r AuthRepository) Read(ID string) (*entity.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*60)

	defer cancel()

	user, err := r.cache.GetUser(ctx, ID)

	if err == nil {
		return user, nil
	}

	user, err = r.ds.Read(ID)

	if err != nil {
		return nil, err
	}

	r.cache.PutUser(ctx, *user)

	return user, nil
}

func (r AuthRepository) ReadUserByEmail(email string) (*entity.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*60)

	defer cancel()

	user, err := r.ds.ReadUserByEmail(email)

	if err != nil {
		return nil, err
	}

	r.cache.PutUser(ctx, *user)

	return user, nil

}

func (r AuthRepository) Update(user entity.User) (*entity.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*60)

	defer cancel()

	uUser, err := r.ds.Update(user)

	if err != nil {
		return nil, err
	}

	r.cache.PutUser(ctx, *uUser)

	return uUser, nil
}

func (r AuthRepository) Delete(ID string) error {
	return r.ds.Delete(ID)
}

func (r AuthRepository) ReadAllUsers(limit, offset *int) (*[]entity.User, error) {
	return r.ds.ReadAllUsers(limit, offset)
}
