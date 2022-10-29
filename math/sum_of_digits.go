package math

// SumOfDigits returns sum of digits of the integer
func SumOfDigits(num int64) (sum int64) {
	// If num is negative then negative sign is neglected.
	// -9223372036854775808 is the min value of int64, but if multiplied by -1 we get
	// 9223372036854775808 which is > 9223372036854775807, the max value of int64, causing overflow
	if num < 0 && num > -9223372036854775808 {
		num = -1 * num
	}

	for num > 0 {
		sum = sum + (num % 10)
		num = num / 10
	}
	return sum
}
