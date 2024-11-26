// reversebits.go
// description: Reverse bits
// details:
// Reverse the bits of an integer number
// time complexity: O(log(n))
// space complexity: O(1)
// author(s) [red_byte](https://github.com/i-redbyte)
// see reversebits_test.go

package binary

// ReverseBits This function initialized the result by 0 (all bits 0) and process the given number starting
// from its least significant bit. If the current bit is 1, set the corresponding most significant bit in the result
// and finally move on to the next bit in the input number.
// Repeat this till all its bits are processed.
func ReverseBits(number uint) uint {
	result := uint(0)
	intSize := 31
	for number != 0 {
		result += (number & 1) << intSize
		number = number >> 1
		intSize -= 1
	}
	return result
}
