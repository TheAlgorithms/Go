// pronicnumber.go
// description: Returns true if the number is pronic, false otherwise
// details: A number n is a pronic number if
// it is equal to product of two consecutive numbers m and m+1.
// wikipedia: https://en.wikipedia.org/wiki/Pronic_number
// author: Akshay Dubey (https://github.com/itsAkshayDubey)
// see pronicnumber_test.go

package math

// This function returns true if argument passed to the function is pronic, false otherwise.
func PronicNumber(n int) (bool, error) {
	switch {
	case n < 0:
		return false, ErrPosArgsOnly
	case n == 0:
		return true, nil
	default:
		for i := 1; i <= n/2; i++ {
			if i*(i+1) == n && i != n {
				return true, nil
			}
		}
		return false, nil
	}
}
