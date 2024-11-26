// inttoroman.go
// description: Convert an integer to a roman numeral
// details: This program converts an integer to a roman numeral. The program uses a lookup array to convert the integer to a roman numeral.
// time complexity: O(1)
// space complexity: O(1)

package conversion

import (
	"errors"
)

var (
	// lookup arrays used for converting from an int to a roman numeral extremely quickly.
	r0 = []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"} // 1 - 9
	r1 = []string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"} // 10 - 90
	r2 = []string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"} // 100 - 900
	r3 = []string{"", "M", "MM", "MMM"}                                       // 1,000 - 3,000
)

// IntToRoman converts an integer value to a roman numeral string. An error is
// returned if the integer is not between 1 and 3999.
func IntToRoman(n int) (string, error) {
	if n < 1 || n > 3999 {
		return "", errors.New("integer must be between 1 and 3999")
	}
	// Concatenate strings for each of 4 lookup array categories.
	//
	// Key behavior to note here is how math with integers is handled. Values are floored to the
	// nearest int, not rounded up. For example, 26/10 = 2 even though the actual result is 2.6.
	//
	// For example, lets use an input value of 126:
	// `r3[n%1e4/1e3]` --> 126 % 10_000 = 126 --> 126 / 1_000 = 0.126 (0 as int) --> r3[0] = ""
	// `r2[n%1e3/1e2]` --> 126 % 1_000 = 126 --> 126 / 100 = 1.26 (1 as int) --> r2[1] = "C"
	// `r1[n%100/10]` --> 126 % 100 = 26 --> 26 / 10 = 2.6 (2 as int) --> r1[2] = "XX"
	// `r0[n%10]` --> 126 % 10 = 6 --> r0[6] = "VI"
	// FINAL --> "" + "C" + "XX" + "VI" = "CXXVI"
	//
	// This is efficient in Go. The 4 operands are evaluated,
	// then a single allocation is made of the exact size needed for the result.
	return r3[n%1e4/1e3] + r2[n%1e3/1e2] + r1[n%100/10] + r0[n%10], nil
}
