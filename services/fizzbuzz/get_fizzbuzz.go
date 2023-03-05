package fizzbuzz

import (
	"github.com/edjubert/leboncoin/types"
)

func GetFizzBuzz(params types.Params) string {
	raw := FizzBuzz{}
	return raw.Generate(params).String()
}
