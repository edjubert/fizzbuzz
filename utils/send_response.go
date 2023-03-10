package utils

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/edjubert/fizzbuzz/constants"
	"github.com/gookit/slog"
)

func Ok(ctx context.Context, w http.ResponseWriter, msg string) {
	Response(ctx, w, msg, http.StatusOK, nil)
}

func InternalServerError(ctx context.Context, w http.ResponseWriter, err error) {
	Response(ctx, w, "", http.StatusInternalServerError, err)
}

func NotImplemented(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	msg := fmt.Sprintf("[%s] %s -> Not implemented", r.URL.Path, r.Method)
	Response(ctx, w, msg, http.StatusNotImplemented, errors.New(msg))
}

func NotFound(ctx context.Context, w http.ResponseWriter, url string) {
	ctx = context.WithValue(
		ctx,
		constants.CTX_ADDR,
		url,
	)
	Response(ctx, w, "Page not found", http.StatusNotFound, errors.New("Not found"))
}

func Response(ctx context.Context, w http.ResponseWriter, msg string, status int, err error) {
	addr := ctx.Value(constants.CTX_ADDR)
	if addr != "" {
		if err != nil {
			slog.Errorf("[%s] -> %d: %s", addr, status, err.Error())
		} else {
			slog.Infof("[%s] -> %d", addr, status)
			slog.Tracef("%s", msg)
		}
	}

	w.WriteHeader(status)
	w.Write([]byte(msg))
}
