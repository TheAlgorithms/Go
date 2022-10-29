// pronicnumber.go
// description: Returns true if the number is pronic, false otherwise
// details: A number n is a pronic number if
// it is equal to product of two consecutive numbers m and m+1.
// wikipedia: https://en.wikipedia.org/wiki/Pronic_number
// author: Akshay Dubey (https://github.com/itsAkshayDubey)
// see pronicnumber_test.go

package math

import stdMath "math"

// This function returns true if argument passed to the function is pronic, false otherwise.
func PronicNumber(n int) bool {
	switch {
	case n < 0 || n%2 == 1:
		return false
	case n == 0:
		return true
	default:
		x := stdMath.Floor(stdMath.Sqrt(float64(n)))
		return n == int(x*(x+1))
	}
}
