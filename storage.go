package tokenredis

import (
	"context"
	"encoding/json"
	"github.com/go-funcards/jwt"
	"github.com/go-redis/redis/v8"
	"time"
)

type Storage struct {
	Redis *redis.Client
}

func (s *Storage) Set(ctx context.Context, refreshToken string, user jwt.User, expiration time.Duration) error {
	data, err := json.Marshal(user)
	if err != nil {
		return err
	}
	if err = s.Redis.Set(ctx, refreshToken, string(data), expiration).Err(); err != nil {
		return err
	}
	return nil
}

func (s *Storage) Get(ctx context.Context, refreshToken string) (user jwt.User, err error) {
	data, err := s.Redis.Get(ctx, refreshToken).Result()
	if err != nil {
		return user, err
	}
	err = json.Unmarshal([]byte(data), &user)
	return
}

func (s *Storage) Del(ctx context.Context, refreshToken string) {
	s.Redis.Del(ctx, refreshToken)
}
