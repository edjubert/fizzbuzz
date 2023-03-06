package home

import (
	"context"
	"net/http"

	"github.com/edjubert/fizzbuzz/utils"
)

func Base(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	utils.Ok(ctx, w, "ok")
}
