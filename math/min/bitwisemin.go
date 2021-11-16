// bitwisemin.go
// description: Gives min of two integers
// details:
// implementation of finding the minimum of two numbers using only binary operations without using conditions
// author(s) [red_byte](https://github.com/i-redbyte)
// see bitwisemin_test.go

package min

// Bitwise This function returns the minimum integer using bit operations
func Bitwise(base int, value int, values ...int) int {
	min := value
	for _, val := range values {
		min = min&((min-val)>>base) | val&(^(min-val)>>base)
	}
	return min
}
