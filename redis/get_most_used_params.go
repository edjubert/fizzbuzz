package redis

import (
	"context"
	"errors"
	"net/url"
	"strconv"
	"strings"

	"github.com/edjubert/leboncoin/types"
	"github.com/gookit/slog"
)

func (r *redisStruct) GetMostUsedParams(ctx context.Context) (types.Params, float64, error) {
	keys, err := r.client.ZRevRangeWithScores(ctx, REDIS_ZSET, 0, 0).Result()
	if err != nil {
		return types.Params{}, 0, err
	}

	for _, key := range keys {
		msg, ok := key.Member.(string)
		if !ok {
			slog.Errorf("This is not ok sir", msg, ok)
		}

		v := strings.Split(msg, ":")
		if len(v) != 5 {
			return types.Params{}, 0, errors.New("Wrong split len")
		}

		int1, err := strconv.Atoi(v[0])
		if err != nil {
			return types.Params{}, 0, err
		}
		int2, err := strconv.Atoi(v[1])
		if err != nil {
			return types.Params{}, 0, err
		}

		limit, err := strconv.Atoi(v[2])
		if err != nil {
			return types.Params{}, 0, err
		}

		str1, err := url.QueryUnescape(v[3])
		if err != nil {
			return types.Params{}, 0, err
		}

		str2, err := url.QueryUnescape(v[4])
		if err != nil {
			return types.Params{}, 0, err
		}

		return types.Params{
			Int1:  int1,
			Int2:  int2,
			Limit: limit,
			Str1:  str1,
			Str2:  str2,
		}, key.Score, nil
	}

	return types.Params{}, 0, nil
}
