package binary

// ABS - this function return absolute value use binary operation
// Principle of operation:
// 1) Get the mask by right shift by the bit width of the variable (int64, int32, int16, int8)
// 2) For negative numbers, above step sets mask as 1 1 1 1 1 1 1 1 and 0 0 0 0 0 0 0 0 for positive numbers.
// 3) Add the mask to the given number.
// 4) XOR of mask + n and mask gives the absolute value.
func ABS(base, n int) int {
	m := n >> (base - 1)
	return (n + m) ^ m
}
