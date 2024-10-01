// fibonacci.go
// description: Get the nth Fibonacci Number
// details:
// In mathematics, the Fibonacci numbers, commonly denoted Fn, form a sequence, called the Fibonacci sequence, such that each number is the sum of the two preceding ones, starting from 0 and 1. [Fibonacci number](https://en.wikipedia.org/wiki/Fibonacci_number)
// time complexity: O(log n)
// space complexity: O(1)
// author(s) [red_byte](https://github.com/i-redbyte)
// see fibonacci_test.go

package fibonacci

import (
	"math"
)

// Matrix This function calculates the n-th fibonacci number using the matrix method. [See](https://en.wikipedia.org/wiki/Fibonacci_number#Matrix_form)
func Matrix(n uint) uint {
	a, b := 1, 1
	c, rc, tc := 1, 0, 0
	d, rd := 0, 1

	for n != 0 {
		if n&1 == 1 {
			tc = rc
			rc = rc*a + rd*c
			rd = tc*b + rd*d
		}

		ta := a
		tb := b
		tc = c
		a = a*a + b*c
		b = ta*b + b*d
		c = c*ta + d*c
		d = tc*tb + d*d

		n >>= 1
	}
	return uint(rc)
}

// Formula This function calculates the n-th fibonacci number using the [formula](https://en.wikipedia.org/wiki/Fibonacci_number#Relation_to_the_golden_ratio)
// Attention! Tests for large values fall due to rounding error of floating point numbers, works well, only on small numbers
func Formula(n uint) uint {
	sqrt5 := math.Sqrt(5)
	phi := (sqrt5 + 1) / 2
	powPhi := math.Pow(phi, float64(n))
	return uint(powPhi/sqrt5 + 0.5)
}

// Recursive calculates the n-th fibonacci number recursively by adding the previous two Fibonacci numbers.
// This algorithm is extremely slow for bigger numbers, but provides a simpler implementation.
func Recursive(n uint) uint {
	if n <= 1 {
		return n
	}

	return Recursive(n-1) + Recursive(n-2)
}
