// palindromicnumber.go
// description: Returns true if the number is palindromic and false otherwise
// details:
//  Palindromic number: A number that remains the same when
//  its digits are reversed is called a palindromic number.
// wikipedia: https://en.wikipedia.org/wiki/Palindromic_number
// author: Akshay Dubey (https://github.com/itsAkshayDubey)
// see palindromicnumber_test.go

package math

// PalindromicNumber returns true if argument passed to the function is palindromic and false otherwise.
func PalindromicNumber(n int) bool {
	switch {
	case n < 0:
		return false
	case n < 10:
		return true
	default:
		n_copy := n
		reversed := 0
		for n_copy > 0 {
			last_digit := n_copy % 10
			n_copy /= 10
			reversed = reversed*10 + last_digit
		}
		return n == reversed
	}
}
