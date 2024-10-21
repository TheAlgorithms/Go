// arithmeticmean.go
// description: Arithmetic mean
// details:
// The most common type of average is the arithmetic mean. If n numbers are given, each number denoted by ai (where i = 1,2, ..., n), the arithmetic mean is the sum of the as divided by n or - [Arithmetic mean](https://en.wikipedia.org/wiki/Average#Arithmetic_mean)
// time complexity: O(1)
// space complexity: O(1)
// author(s) [red_byte](https://github.com/i-redbyte)
// see arithmeticmean_test.go

// Package binary describes algorithms that use binary operations for different calculations.
package binary

// MeanUsingAndXor This function finds arithmetic mean using "AND" and "XOR" operations
func MeanUsingAndXor(a int, b int) int {
	return ((a ^ b) >> 1) + (a & b)
}

// MeanUsingRightShift This function finds arithmetic mean using right shift
func MeanUsingRightShift(a int, b int) int {
	return (a + b) >> 1
}
