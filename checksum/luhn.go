// lunh.go
// description: Luhn algorithm
// details: is a simple checksum formula used to validate a variety of identification numbers, such as credit card numbers, IMEI numbers, etc [Lunh](https://en.wikipedia.org/wiki/Luhn_algorithm)
// author(s) [red_byte](https://github.com/i-redbyte)
// see lunh_test.go

// Package checksum describes algorithms for finding various checksums
package checksum

// LuhnAlgorithm This function calculates the checksum using the Luna algorithm
func LuhnAlgorithm(s string) bool {
	n := len(s)
	number := 0
	result := 0
	for i := 0; i < n; i++ {
		number = int(s[i]) - '0'
		if i%2 == 0 {
			number *= 2
			if number > 9 {
				number -= 9
			}
		}
		result += number
	}
	return result%10 == 0
}
