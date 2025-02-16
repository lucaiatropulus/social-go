package cache

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/lucaiatropulus/social/internal/dao"
)

type Storage struct {
	Users interface {
		Get(ctx context.Context, userID int64) (*dao.User, error)
		Set(ctx context.Context, user *dao.User) error
		Delete(ctx context.Context, userID int64) error
	}
}

func NewRedisStorage(redisDB *redis.Client) Storage {
	return Storage{
		Users: &UserStore{redisDB},
	}
}
