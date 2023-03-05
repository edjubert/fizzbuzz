package statistics

import (
	"context"
	"net/http"

	"github.com/edjubert/leboncoin/redis"
	"github.com/edjubert/leboncoin/types"
	"github.com/edjubert/leboncoin/utils"
)

func Statistics(ctx context.Context, w http.ResponseWriter, r *http.Request, redis redis.Cache) {
	params, score, err := redis.GetMostUsedParams(ctx)
	if err != nil {
		utils.InternalServerError(ctx, w, err)
		return
	} else if params.IsEmpty() {
		utils.Response(ctx, w, "Empty statistics", http.StatusOK, nil)
		return
	}

	statsResponse := types.StatsResponse{
		Params: params,
		Score:  int(score),
	}
	str, err := utils.MarshalParams(statsResponse)
	if err != nil {
		utils.InternalServerError(ctx, w, err)
		return
	}

	utils.Ok(ctx, w, string(str))
}
