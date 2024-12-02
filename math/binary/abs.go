// abs.go
// description: returns absolute value using binary operation
// time complexity: O(1)
// space complexity: O(1)

package binary

// Abs returns absolute value using binary operation
// Principle of operation:
// 1) Get the mask by right shift by the base
// 2) Base is the size of an integer variable in bits, for example, for int32 it will be 32, for int64 it will be 64
// 3) For negative numbers, above step sets mask as 1 1 1 1 1 1 1 1 and 0 0 0 0 0 0 0 0 for positive numbers.
// 4) Add the mask to the given number.
// 5) XOR of mask + n and mask gives the absolute value.
func Abs(base, n int) int {
	m := n >> (base - 1)
	return (n + m) ^ m
}
