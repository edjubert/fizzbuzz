package controllers

import (
	"context"
	"io/ioutil"
	"net/http"

	"github.com/edjubert/fizzbuzz/constants"
	"github.com/edjubert/fizzbuzz/redis"
	"github.com/edjubert/fizzbuzz/services/fizzbuzz"
	"github.com/edjubert/fizzbuzz/utils"
)

func FizzBuzzPost(w http.ResponseWriter, r *http.Request, redis redis.Cache, ctx context.Context) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utils.InternalServerError(ctx, w, err)
		return
	}

	params, err := utils.UnmarshalParams(body)
	if err != nil {
		utils.InternalServerError(ctx, w, err)
		return
	}

	msg := fizzbuzz.GetFizzBuzz(ctx, *params, redis)
	utils.Ok(ctx, w, msg)
}

func FizzBuzz(w http.ResponseWriter, r *http.Request, redis redis.Cache) {
	ctx := context.WithValue(context.Background(), constants.CTX_ADDR, constants.FIZZBUZZ)
	switch r.Method {
	case "POST":
		FizzBuzzPost(w, r, redis, ctx)
	default:
		utils.NotImplemented(ctx, w, r)
	}
}
