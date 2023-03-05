package redis

import (
	"time"

	"github.com/gookit/slog"
)

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
