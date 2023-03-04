package utils

import (
	"encoding/json"

	"github.com/edjubert/leboncoin/types"
)

func UnmarshalParams[T types.HttpParams](str []byte) (*T, error) {
	params := new(T)
	err := json.Unmarshal(str, &params)
	if err != nil {
		return nil, err
	}

	return params, nil
}
