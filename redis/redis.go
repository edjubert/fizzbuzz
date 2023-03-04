package redis

import (
	"context"
	"errors"
	"time"

	"github.com/gookit/slog"
	"github.com/redis/go-redis/v9"
)

type redisStruct struct {
	client *redis.Client
	ctx    context.Context
}

func (r *redisStruct) Ping() error {
	status := r.client.Ping(r.ctx)
	if _, err := status.Result(); err != nil {
		slog.Error(err)
		return errors.New("No connection")
	}

	return nil
}

func (r *redisStruct) Close() error {
	if err := r.client.Close(); err != nil {
		slog.Error(err)
		return errors.New("Cannot close connection")
	}

	return nil
}

type Cache interface {
	Ping() error
	Close() error
}

var _ Cache = (*redisStruct)(nil)

func NewRedisClient(addr string) Cache {
	ctx := context.Background()
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "",
		DB:       0,
	})

	return &redisStruct{
		client: client,
		ctx:    ctx,
	}
}

func GetRedisClient(addr string, wait bool) Cache {
	redisCache := NewRedisClient(addr)
	if redisCache.Ping() != nil {
		slog.Warn("Redis not reachable", addr)
	}

	for wait && redisCache.Ping() != nil {
		slog.Infof("Waiting for redis on %s", addr)
		time.Sleep(5 * time.Second)
	}

	return redisCache
}
