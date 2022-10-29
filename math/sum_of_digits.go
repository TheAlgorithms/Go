package math

// SumOfDigits returns sum of digits of the integer
func SumOfDigits(num int64) (sum int64) {
	// If num is negative then negative sign is neglected.
	if num < 0 {
		num = -1 * num
	}

	for num > 0 {
		sum = sum + (num % 10)
		num = num / 10
	}
	return sum
}
