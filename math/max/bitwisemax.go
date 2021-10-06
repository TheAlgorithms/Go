// bitwiseMax.go
// description: Gives max of two integers
// details:
// implementation of finding the maximum of two numbers using only binary operations without using conditions
// author(s) [red_byte](https://github.com/i-redbyte)
// see bitwiseMax_test.go

package max

func BitwiseMax(a int, b int, base int) int {
	z := a - b
	i := (z >> base) & 1
	return a - (i * z)
}
