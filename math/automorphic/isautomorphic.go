// isautomorphic.go
// description: Checks whether a whole number integer is Automorphic or not. If number < 0 then returns false.
// details:
// In mathematics, a number n is said to be a Automorphic number if the square of n ends in the same digits as n itself.
// ref: (https://en.wikipedia.org/wiki/Automorphic_number)
// time complexity: O(log10(N))
// space complexity: O(1)
// author: [SilverDragonOfR](https://github.com/SilverDragonOfR)

// Package sumofdigits provides method to determine whether a number is automorphic or not.
package automorphic

func IsAutomorphic(n int) bool {

	// handling the negetive number case
	if n < 0 {
		return false
	}

	var n_sq int = n * n

	for n > 0 {

		// checking whether the number and it's square have same rightMost digit
		if (n % 10) != (n_sq % 10) {
			return false
		}

		// update the number and it's square by removing the rightMost digit
		n /= 10
		n_sq /= 10
	}

	return true
}
