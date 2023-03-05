package statistics

import (
	"context"
	"net/http"

	"github.com/edjubert/leboncoin/redis"
	"github.com/edjubert/leboncoin/utils"
)

func Statistics(ctx context.Context, w http.ResponseWriter, r *http.Request, redis redis.Cache) {
	utils.Ok(ctx, w, "ok")
}
