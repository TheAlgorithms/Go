package conversion

import "errors"

var (
	// lookup arrays used for converting from an int to a roman numeral extremely quickly.
	r0 = []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}
	r1 = []string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
	r2 = []string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
	r3 = []string{"", "M", "MM", "MMM"}
)

// IntToRoman converts an integer value to a roman numeral string. An error is
// returned if the integer is not between 1 and 3999.
func IntToRoman(n int) (string, error) {
	if n < 1 || n > 3999 {
		return "", errors.New("integer must be between 1 and 3999")
	}
	// This is efficient in Go. The 4 operands are evaluated,
	// then a single allocation is made of the exact size needed for the result.
	return r3[n%1e4/1e3] + r2[n%1e3/1e2] + r1[n%100/10] + r0[n%10], nil
}
