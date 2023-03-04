package statistics

import (
	"net/http"

	"github.com/edjubert/leboncoin/redis"
	"github.com/edjubert/leboncoin/utils"
)

func Statistics(w http.ResponseWriter, r *http.Request, redis redis.Cache) {
	utils.SendResponse(w, "ok", "/", http.StatusOK, nil)
}
