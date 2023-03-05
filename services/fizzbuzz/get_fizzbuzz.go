package fizzbuzz

import (
	"context"

	"github.com/edjubert/leboncoin/redis"
	"github.com/edjubert/leboncoin/types"
)

func GetFizzBuzz(ctx context.Context, params types.Params, redis redis.Cache) string {
	raw := FizzBuzz{}

	if msg, _ := redis.GetMsgFromParams(ctx, params); msg != "" {
		redis.UpdateScore(ctx, params)
		return msg
	}

	msg := raw.Generate(params).String()
	if msg != "" {
		redis.UpdateScore(ctx, params)
		redis.SaveParamsAndMsg(ctx, params, msg)
	}

	return msg
}
