package tokenredis

import (
	"github.com/go-funcards/jwt"
	"github.com/go-funcards/token"
	"github.com/go-redis/redis/v8"
)

func New(cfg token.Config, generator jwt.Generator, rdb *redis.Client) token.Service {
	return token.New(cfg, generator, &Storage{Redis: rdb})
}
