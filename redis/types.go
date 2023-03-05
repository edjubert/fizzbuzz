package redis

import (
	"context"

	"github.com/edjubert/leboncoin/types"
	"github.com/redis/go-redis/v9"
)

type redisStruct struct {
	client *redis.Client
	ctx    context.Context
}

type Cache interface {
	Ping() error
	Close() error

	UpdateScore(ctx context.Context, params types.Params) error
	SaveParamsAndMsg(ctx context.Context, params types.Params, msg string) error
	GetMsgFromParams(ctx context.Context, params types.Params) (string, error)
}

const REDIS_ZSET = "ZSet"
const REDIS_HSET = "HSet"
