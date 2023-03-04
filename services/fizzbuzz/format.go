package fizzbuzz

import "strings"

func (d *FizzBuzz) String() string {
	return strings.Join(d.data, ",")
}
