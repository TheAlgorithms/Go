// abs.go
// description: Absolute value
// details:
// In mathematics, the absolute value or modulus of a real number x, denoted |x|, is the non-negative value of x without regard to its sign. [Absolute value](https://en.wikipedia.org/wiki/Average#Arithmetic_mean)
// author(s) [red_byte](https://github.com/i-redbyte)
// see abs_test.go

package abs

// ABS - this function return absolute value
func ABS(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
