package data_source

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/viniciusps01/todo/internal/feature/auth/entity"
	infra "github.com/viniciusps01/todo/internal/infra/cache"
)

const ()

func getUserKey(ID string) string {
	return fmt.Sprintf("%s:%s", "user", ID)
}

type AuthCacheDataSource struct {
	cache infra.ICache
}

func NewAuthCacheDataSource(c infra.ICache) IAuthCacheDataSource {
	return AuthCacheDataSource{
		cache: c,
	}
}

func (c AuthCacheDataSource) PutUser(ctx context.Context, user entity.User) error {
	json, err := json.Marshal(user)

	if err != nil {
		return err
	}

	return c.cache.Put(ctx, getUserKey(user.ID), json, time.Second*60)
}

func (c AuthCacheDataSource) GetUser(ctx context.Context, ID string) (*entity.User, error) {
	userStr, err := c.cache.Get(ctx, getUserKey(ID))

	if err != nil {
		return nil, err
	}

	var user entity.User

	err = json.Unmarshal([]byte(*userStr), &user)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
