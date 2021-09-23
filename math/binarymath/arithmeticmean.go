// arithmeticmean.go
// description: Arithmetic mean
// details:
// The most common type of average is the arithmetic mean. If n numbers are given, each number denoted by ai (where i = 1,2, ..., n), the arithmetic mean is the sum of the as divided by n or - [Arithmetic mean](https://en.wikipedia.org/wiki/Average#Arithmetic_mean)
// author(s) [red_byte](https://github.com/i-redbyte)
// see arithmeticmean_test.go

// Package binarymath describes algorithms that use binary operations for different calculations.
package binarymath

func MeanUsingAndXor(a int, b int) int {
	return ((a ^ b) >> 1) + (a & b)
}

func MeanUsingRightShift(a int, b int) int {
	return (a + b) >> 1
}
