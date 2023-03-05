package redis

import (
	"context"

	"github.com/redis/go-redis/v9"
)

var _ Cache = (*redisStruct)(nil)

func NewRedisClient(addr string) Cache {
	ctx := context.Background()
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "",
		DB:       0,
	})

	return &redisStruct{
		client: client,
		ctx:    ctx,
	}
}
