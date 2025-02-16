package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/lucaiatropulus/social/internal/dao"
)

type UserStore struct {
	redisDB *redis.Client
}

const userExpirationTime = time.Hour

func (s *UserStore) Get(ctx context.Context, userID int64) (*dao.User, error) {
	cacheKey := s.generateCacheKey(userID)

	data, err := s.redisDB.Get(ctx, cacheKey).Result()

	if err == redis.Nil {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	var user dao.User

	err = s.decodeUser(data, &user)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *UserStore) Set(ctx context.Context, user *dao.User) error {
	cacheKey := s.generateCacheKey(user.ID)

	json, err := json.Marshal(user)

	if err != nil {
		return err
	}

	return s.redisDB.SetEX(ctx, cacheKey, json, userExpirationTime).Err()
}

func (s *UserStore) Delete(ctx context.Context, userID int64) error {
	cacheKey := s.generateCacheKey(userID)

	return s.redisDB.Del(ctx, cacheKey).Err()
}

func (s *UserStore) generateCacheKey(userID int64) string {
	return fmt.Sprintf("user-%v", userID)
}

func (s *UserStore) decodeUser(data string, user *dao.User) error {
	if data == "" {
		return nil
	}

	err := json.Unmarshal([]byte(data), &user)

	if err != nil {
		return err
	}

	return nil
}
