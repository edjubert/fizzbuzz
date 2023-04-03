package fizzbuzz

import (
	"strconv"

	"github.com/edjubert/fizzbuzz/types"
)

func (d *FizzBuzz) Generate(params types.Params) *FizzBuzz {
	int1 := params.Int1
	int2 := params.Int2
	limit := params.Limit
	str1 := params.Str1
	str2 := params.Str2

	for _, i := range []int{int1, int2, limit} {
		if i < 1 {
			return d
		}
	}

	for i := 1; i <= limit; i++ {
		str := strconv.Itoa(i)

		onlyStr1 := i%int1 == 0 && i%int2 != 0
		onlyStr2 := i%int1 != 0 && i%int2 == 0
		both := i%int1 == 0 && i%int2 == 0

		if onlyStr1 {
			str = str1
		} else if onlyStr2 {
			str = str2
		} else if both {
			str = str1 + str2
		}

		d.data = append(d.data, str)
	}

	return d
}
