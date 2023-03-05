package controllers

import (
	"context"
	"net/http"

	"github.com/edjubert/leboncoin/constants"
	"github.com/edjubert/leboncoin/services/home"
)

func Base(w http.ResponseWriter, r *http.Request) {
	ctx := context.WithValue(
		context.Background(),
		constants.CTX_ADDR,
		constants.HOME_BASE,
	)

	if r.URL.Path != constants.HOME_BASE {
		NotFound(w, r)
		return
	}

	switch r.Method {
	case "GET":
		home.Base(ctx, w, r)
	}
}
