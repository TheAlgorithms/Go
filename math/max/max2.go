// max2.go
// description: Gives max of two integers
// details:
// implementation of finding the maximum of two numbers using only binary operations without using conditions
// author(s) [red_byte](https://github.com/i-redbyte)
// see max2_test.go

package max

func Max2(a int, b int, base int) int {
	z := a - b
	i := (z >> base) & 1
	return a - (i * z)
}
