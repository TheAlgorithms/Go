// filename : krishnamurthy.go
// description: A program which contains the function that returns true if a given number is Krishnamurthy number or not.
// details: A number is a Krishnamurthy number if the sum of all the factorials of the digits is equal to the number.
// Ex: 1! = 1, 145 = 1! + 4! + 5!
// time complexity: O(log n)
// space complexity: O(1)
// author(s): [GooMonk](https://github.com/GooMonk)
// see krishnamurthy_test.go
package math

import "github.com/TheAlgorithms/Go/constraints"

// IsKrishnamurthyNumber returns if the provided number n is a Krishnamurthy number or not.
func IsKrishnamurthyNumber[T constraints.Integer](n T) bool {
	if n <= 0 {
		return false
	}

	// Preprocessing: Using a slice to store the digit Factorials
	digitFact := make([]T, 10)
	digitFact[0] = 1 // 0! = 1

	for i := 1; i < 10; i++ {
		digitFact[i] = digitFact[i-1] * T(i)
	}

	// Subtract the digit Facotorial from the number
	nTemp := n
	for n > 0 {
		nTemp -= digitFact[n%10]
		n /= 10
	}
	return nTemp == 0
}
