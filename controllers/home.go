package controllers

import (
	"context"
	"net/http"

	"github.com/edjubert/fizzbuzz/constants"
	"github.com/edjubert/fizzbuzz/services/home"
	"github.com/edjubert/fizzbuzz/utils"
)

func Base(w http.ResponseWriter, r *http.Request) {
	ctx := context.WithValue(
		context.Background(),
		constants.CTX_ADDR,
		constants.HOME_BASE,
	)

	if r.URL.Path != constants.HOME_BASE {
		utils.NotFound(ctx, w, r.URL.Path)
		return
	}

	switch r.Method {
	case "GET":
		home.Base(ctx, w, r)
	default:
		utils.NotImplemented(ctx, w, r)
	}
}
