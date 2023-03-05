package redis

import (
	"errors"

	"github.com/gookit/slog"
)

func (r *redisStruct) Close() error {
	if err := r.client.Close(); err != nil {
		slog.Error(err)
		return errors.New("Cannot close connection")
	}

	return nil
}
