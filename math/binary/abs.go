package binary

// ABS - this function return absolute value use binary operation
func ABS(base, n int) int {
	m := n >> (base - 1)
	return (n + m) ^ m
}
