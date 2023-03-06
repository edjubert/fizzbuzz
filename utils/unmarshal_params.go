package utils

import (
	"encoding/json"

	"github.com/edjubert/fizzbuzz/types"
	"github.com/gookit/slog"
)

func UnmarshalParams[T types.HttpParams](str []byte) (*T, error) {
	params := new(T)
	if err := json.Unmarshal(str, &params); err != nil {
		slog.Error(err)
		return nil, err
	}

	return params, nil
}
