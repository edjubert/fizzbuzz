package main

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"

	"github.com/edjubert/fizzbuzz/constants"
	"github.com/edjubert/fizzbuzz/controllers"
	"github.com/edjubert/fizzbuzz/redis"
	"github.com/edjubert/fizzbuzz/utils"
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

func startServer(ctx context.Context, mux *http.ServeMux, cancel context.CancelFunc) {
	server := &http.Server{
		Handler: mux,
		BaseContext: func(l net.Listener) context.Context {
			slog.Noticef("Running server on port %v", l.Addr())
			ctx = context.WithValue(ctx, KEY_SERVER_ADDR, l.Addr().String())
			return ctx
		},
	}
	server.Addr = fmt.Sprintf(":%d", utils.GetPort(false))
	err := server.ListenAndServe()
	if errors.Is(err, http.ErrServerClosed) {
		slog.Warn("Server closed")
	} else if err != nil {
		slog.Error(err)
	}
	cancel()
}

func startHealthCheckServer(cancel context.CancelFunc) {
	addr := fmt.Sprintf(":%d", utils.GetPort(true))
	http.HandleFunc(
		"/",
		func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path != "/" {
				utils.NotFound(context.Background(), w, fmt.Sprintf("HealthCheck | %s", r.URL.Path))
				return
			}
			slog.Infof("[HealthCheck | %s] -> %d", r.URL.Path, http.StatusOK)
			_, _ = w.Write([]byte("ok"))
		},
	)
	slog.Noticef("Running health check server on port [::]%v", addr)
	_ = http.ListenAndServe(addr, nil)

	cancel()
}

func server(r redis.Cache) {
	mux := http.NewServeMux()
	handleRoutes(mux, r)

	ctx, cancelCtx := context.WithCancel(context.Background())
	defer cancelCtx()

	go startServer(ctx, mux, cancelCtx)
	go startHealthCheckServer(cancelCtx)

	<-ctx.Done()
}
