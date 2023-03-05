package redis

import (
	"context"

	"github.com/edjubert/leboncoin/types"
	"github.com/gookit/slog"
)

func (r *redisStruct) SaveParamsAndMsg(ctx context.Context, params types.Params, msg string) error {
	member := getRedisMemberKey(params)

	if err := r.client.HSet(ctx, REDIS_HSET, member, msg).Err(); err != nil {
		slog.Error(err)
		return err
	}

	return nil
}
