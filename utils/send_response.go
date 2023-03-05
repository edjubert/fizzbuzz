package utils

import (
	"context"
	"net/http"

	"github.com/edjubert/leboncoin/constants"
	"github.com/gookit/slog"
)

func Ok(ctx context.Context, w http.ResponseWriter, msg string) {
	Response(ctx, w, msg, http.StatusOK, nil)
}

func InternalServerError(ctx context.Context, w http.ResponseWriter, err error) {
	Response(ctx, w, "", http.StatusInternalServerError, err)
}

func NotFound(ctx context.Context, w http.ResponseWriter) {
	Response(ctx, w, "page not found", http.StatusNotFound, nil)
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
