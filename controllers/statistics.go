package controllers

import (
	"context"
	"net/http"

	"github.com/edjubert/fizzbuzz/constants"
	"github.com/edjubert/fizzbuzz/redis"
	"github.com/edjubert/fizzbuzz/services/statistics"
	"github.com/edjubert/fizzbuzz/utils"
)

func Statistics(w http.ResponseWriter, r *http.Request, redis redis.Cache) {
	ctx := context.WithValue(
		context.Background(),
		constants.CTX_ADDR,
		constants.STATISTICS,
	)

	switch r.Method {
	case "GET":
		statistics.Statistics(ctx, w, r, redis)
	default:
		utils.NotImplemented(ctx, w, r)
	}
}
