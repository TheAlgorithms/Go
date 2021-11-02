// bitwisemin.go
// description: Gives min of two integers
// details:
// implementation of finding the minimum of two numbers using only binary operations without using conditions
// author(s) [red_byte](https://github.com/i-redbyte)
// see bitwisemin_test.go

package min

func BitwiseMin(a int, b int, base int) int {
	return a&((a-b)>>base) | b&(^(a-b)>>base)
}
