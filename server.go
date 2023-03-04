package main

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"

	"github.com/edjubert/leboncoin/constants"
	"github.com/edjubert/leboncoin/controllers"
	"github.com/edjubert/leboncoin/redis"
	"github.com/edjubert/leboncoin/utils"
	"github.com/gookit/slog"
)

const KEY_SERVER_ADDR = "serverAddr"

func handleRoutes(mux *http.ServeMux, redis redis.Cache) {
	mux.HandleFunc(constants.HOME_BASE, controllers.Base)
	mux.HandleFunc(constants.FIZZBUZZ, func(w http.ResponseWriter, r *http.Request) {
		controllers.FizzBuzz(w, r, redis)
	})
	mux.HandleFunc(constants.STATISTICS, func(w http.ResponseWriter, r *http.Request) {
		controllers.Statistics(w, r, redis)
	})
}

func server(r redis.Cache) {
	mux := http.NewServeMux()
	handleRoutes(mux, r)

	ctx, cancelCtx := context.WithCancel(context.Background())
	defer cancelCtx()

	addr := fmt.Sprintf(":%d", utils.GetPort())
	server := &http.Server{
		Addr:    addr,
		Handler: mux,
		BaseContext: func(l net.Listener) context.Context {
			slog.Noticef("Running server on port %d", utils.GetPort())
			ctx = context.WithValue(ctx, KEY_SERVER_ADDR, l.Addr().String())
			return ctx
		},
	}

	go func() {
		err := server.ListenAndServe()
		if errors.Is(err, http.ErrServerClosed) {
			slog.Warn("Server closed")
		} else if err != nil {
			slog.Error(err)
		}
		cancelCtx()
	}()

	<-ctx.Done()
}
