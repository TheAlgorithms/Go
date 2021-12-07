// catalannumber.go
// description: Returns the Catalan number
// details:
// In combinatorial mathematics, the Catalan numbers are a sequence of natural numbers that occur in various counting problems, often involving recursively defined objects. - [Catalan number](https://en.wikipedia.org/wiki/Catalan_number)
// The input is the number of the Catalan number n, at the output we get the value of the number
// author(s) [red_byte](https://github.com/i-redbyte)
// see catalannumber_test.go

package catalan

import (
	f "github.com/TheAlgorithms/Go/math/factorial"
)

// CatalanNumber This function returns the `nth` Catalan number
func CatalanNumber(n int) int {
	return f.Iterative(n*2) / (f.Iterative(n) * f.Iterative(n+1))
}
