// bitwiseMax.go
// description: Gives max of two integers
// details:
// implementation of finding the maximum of two numbers using only binary operations without using conditions
// author(s) [red_byte](https://github.com/i-redbyte)
// time complexity: O(1)
// space complexity: O(1)
// see bitwiseMax_test.go

package max

// Bitwise computes using bitwise operator the maximum of all the integer input and returns it
func Bitwise(a int, b int, base int) int {
	z := a - b
	i := (z >> base) & 1
	return a - (i * z)
}
