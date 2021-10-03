// reversebits.go
// description: Reverse bits
// details:
// Reverse the bits of an integer number
// author(s) [red_byte](https://github.com/i-redbyte)
// see reversebits_test.go

package binary

func ReverseBits(number uint) uint {
	result := uint(0)
	powerTwo := 31
	for number != 0 {
		result += (number & 1) << powerTwo
		number = number >> 1
		powerTwo -= 1
	}
	return result
}
