package redis

import (
	"context"

	"github.com/edjubert/leboncoin/types"
	"github.com/gookit/slog"
	"github.com/redis/go-redis/v9"
)

func (r *redisStruct) UpdateScore(ctx context.Context, params types.Params) error {
	member := getRedisMemberKey(params)

	score, _ := r.client.ZScore(ctx, REDIS_ZSET, member).Result()
	if err := r.client.ZAdd(ctx, REDIS_ZSET, redis.Z{
		Score:  score + 1,
		Member: member,
	}).Err(); err != nil {
		slog.Error(err)
		return err
	}

	return nil
}
