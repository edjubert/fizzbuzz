package redis

import (
	"context"

	"github.com/edjubert/leboncoin/types"
)

func (r *redisStruct) GetMsgFromParams(ctx context.Context, params types.Params) (string, error) {
	member := getRedisMemberKey(params)
	if exists, err := r.client.HExists(ctx, REDIS_HSET, member).Result(); err != nil {
		return "", err
	} else if !exists {
		return "", nil
	}

	msg, err := r.client.HGet(ctx, REDIS_HSET, member).Result()
	if err != nil {
		return "", err
	}

	return msg, nil
}
