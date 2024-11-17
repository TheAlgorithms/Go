// pronicnumber.go
// description: Returns true if the number is pronic and false otherwise
// details:
//  Pronic number: For any integer n, if there exists integer m
//  such that n = m * (m + 1) then n is called a pronic number.
// wikipedia: https://en.wikipedia.org/wiki/Pronic_number
// time complexity: O(1)
// space complexity: O(1)
// author: Akshay Dubey (https://github.com/itsAkshayDubey)
// see pronicnumber_test.go

package math

import "math"

// PronicNumber returns true if argument passed to the function is pronic and false otherwise.
func PronicNumber(n int) bool {
	if n < 0 || n%2 == 1 {
		return false
	}
	x := int(math.Sqrt(float64(n)))
	return n == x*(x+1)
}
