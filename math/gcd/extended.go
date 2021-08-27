package gcd

// Extended simple extended gcd
func Extended(a, b int64) (int64, int64, int64) {
	if a == 0 {
		return b, 0, 1
	}
	gcd, xPrime, yPrime := Extended(b%a, a)
	return gcd, yPrime - (b/a)*xPrime, xPrime
}
