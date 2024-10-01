// factorial.go
// description: Calculating factorial
// details:
// The factorial of a non-negative integer n, denoted by n!, is the product of all positive integers less than or equal to n - [Factorial](https://en.wikipedia.org/wiki/Factorial)
// time complexity: O(n)
// space complexity: O(1)
// author(s) [red_byte](https://github.com/i-redbyte)
// see factorial_test.go

// Package factorial describes algorithms Factorials calculations.
package factorial

import (
	"errors"
)

var ErrNegativeArgument = errors.New("input argument must be non-negative integer")

// Iterative returns the iteratively brute forced factorial of n
func Iterative(n int) (int, error) {
	if n < 0 {
		return 0, ErrNegativeArgument
	}
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result, nil
}

// Recursive This function recursively computes the factorial of a number
func Recursive(n int) (int, error) {
	if n < 0 {
		return 0, ErrNegativeArgument
	}
	if n <= 1 {
		return 1, nil
	}
	prev, _ := Recursive(n - 1)
	return n * prev, nil
}

// UsingTree This function finds the factorial of a number using a binary tree
func UsingTree(n int) (int, error) {
	if n < 0 {
		return 0, ErrNegativeArgument
	}
	if n == 0 {
		return 1, nil
	}
	if n == 1 || n == 2 {
		return n, nil
	}
	return prodTree(2, n), nil
}

func prodTree(l int, r int) int {
	if l > r {
		return 1
	}
	if l == r {
		return l
	}
	if r-l == 1 {
		return l * r
	}
	m := (l + r) / 2
	return prodTree(l, m) * prodTree(m+1, r)
}
