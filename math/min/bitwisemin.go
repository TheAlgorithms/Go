// bitwisemin.go
// description: Gives min of two integers
// details:
// implementation of finding the minimum of two numbers using only binary operations without using conditions
// author(s) [red_byte](https://github.com/i-redbyte)
// see bitwisemin_test.go

package min

func Bitwise(base int, values ...int) int {
	result := values[0]
	for _, val := range values {
		result = result&((result-val)>>base) | val&(^(result-val)>>base)
	}
	return result
}
