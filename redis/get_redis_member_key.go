package redis

import (
	"fmt"
	"net/url"

	"github.com/edjubert/fizzbuzz/types"
)

func getRedisMemberKey(params types.Params) string {
	return fmt.Sprintf(
		"%d:%d:%d:%s:%s",
		params.Int1,
		params.Int2,
		params.Limit,
		url.QueryEscape(params.Str1),
		url.QueryEscape(params.Str2),
	)
}
