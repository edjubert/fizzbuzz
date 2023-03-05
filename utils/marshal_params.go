package utils

import (
	"encoding/json"

	"github.com/gookit/slog"
)

func MarshalParams[T any](params T) ([]byte, error) {
	str, err := json.Marshal(params)
	if err != nil {
		slog.Error(err)
		return nil, err
	}

	return str, nil
}
