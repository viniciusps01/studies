package data_source

import (
	"context"

	"github.com/viniciusps01/todo/internal/feature/auth/entity"
)

type IAuthCacheDataSource interface {
	PutUser(context context.Context, user entity.User) error
	GetUser(context context.Context, ID string) (*entity.User, error)
}
