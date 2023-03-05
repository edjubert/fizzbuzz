package controllers

import (
	"context"
	"net/http"

	"github.com/edjubert/leboncoin/constants"
	"github.com/edjubert/leboncoin/redis"
	"github.com/edjubert/leboncoin/services/statistics"
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
	}
}
