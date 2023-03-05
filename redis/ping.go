package redis

import (
	"errors"

	"github.com/gookit/slog"
)

func (r *redisStruct) Ping() error {
	status := r.client.Ping(r.ctx)
	if _, err := status.Result(); err != nil {
		slog.Error(err)
		return errors.New("No connection")
	}

	return nil
}
