package controllers

import (
	"net/http"

	"github.com/edjubert/leboncoin/redis"
	"github.com/edjubert/leboncoin/services/statistics"
)

func Statistics(w http.ResponseWriter, r *http.Request, redis redis.Cache) {
	statistics.Statistics(w, r, redis)
}
