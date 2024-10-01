// perfectnumber.go
// description: provides the function IsPerfectNumber and related utilities
// details:
// provides the functions
// - IsPerfectNumber which checks if the input is a perfect number,
// - SumOfProperDivisors which returns the sum of proper divisors of the input.
// A number is called perfect, if it is a sum of its proper divisors,
// cf. https://en.wikipedia.org/wiki/Perfect_number,
// https://mathworld.wolfram.com/PerfectNumber.html
// time complexity: O(sqrt(n))
// space complexity: O(1)
// https://oeis.org/A000396
// author(s) [Piotr Idzik](https://github.com/vil02)
// see perfectnumber_test.go

package math

// Returns the sum of proper divisors of inNumber.
func SumOfProperDivisors(inNumber uint) uint {
	var res = uint(0)
	if inNumber > 1 {
		res = uint(1)
	}
	for curDivisor := uint(2); curDivisor*curDivisor <= inNumber; curDivisor++ {
		if inNumber%curDivisor == 0 {
			res += curDivisor
			if curDivisor*curDivisor != inNumber {
				res += inNumber / curDivisor
			}
		}
	}
	return res
}

// Checks if inNumber is a perfect number
func IsPerfectNumber(inNumber uint) bool {
	return inNumber > 0 && SumOfProperDivisors(inNumber) == inNumber
}
