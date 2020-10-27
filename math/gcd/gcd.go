package gcd

// Gcd finds and returns the greatest common divisor of a given integer.
func Gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return Gcd(b, a%b)
}
