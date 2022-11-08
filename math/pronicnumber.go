// pronicnumber.go
// description: Returns true if the number is pronic and false otherwise
// pronic: Consisting of a base or root or first power plus a higher power.
// details: A number n is a pronic number if
// it is equal to product of two consecutive numbers m and m+1.
// wikipedia: https://en.wikipedia.org/wiki/Pronic_number
// author: Akshay Dubey (https://github.com/itsAkshayDubey)
// see pronicnumber_test.go

package math

import stdMath "math"

// PronicNumber returns true if argument passed to the function is pronic and false otherwise.
func PronicNumber(n int) bool {
	if n < 0 || n%2 == 1 {
		return false
	}
	x := int(stdMath.Sqrt(float64(n)))
	return n == x*(x+1)
}
