// checkisnumberpoweroftwo.go
// description: Is the number a power of two
// details:
// Checks if a number is a power of two- [Power of two](https://en.wikipedia.org/wiki/Power_of_two)
// time complexity: O(1)
// space complexity: O(1)
// author(s) [red_byte](https://github.com/i-redbyte)
// see checkisnumberpoweroftwo_test.go

package binary

// IsPowerOfTwo This function uses the fact that powers of 2 are represented
// like 10...0 in binary, and numbers one less than the power of 2
// are represented like 11...1.
// Therefore, using the and function:
//
//	  10...0
//	& 01...1
//	  00...0 -> 0
//
// This is also true for 0, which is not a power of 2, for which we
// have to add and extra condition.
func IsPowerOfTwo(x int) bool {
	return x > 0 && (x&(x-1)) == 0
}

// IsPowerOfTwoLeftShift This function takes advantage of the fact that left shifting a number
// by 1 is equivalent to multiplying by 2. For example, binary 00000001 when shifted by 3 becomes 00001000,
// which in decimal system is 8 or = 2 * 2 * 2
func IsPowerOfTwoLeftShift(number uint) bool {
	for p := uint(1); p <= number; p = p << 1 {
		if number == p {
			return true
		}
	}
	return false
}
