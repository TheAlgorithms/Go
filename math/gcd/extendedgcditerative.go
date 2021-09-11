package gcd

// ExtendedIterative finds and returns gcd(a, b), x, y satisfying a*x + b*y = gcd(a, b).
func ExtendedIterative(a, b int64) (int64, int64, int64) {
	var u, y, v, x int64 = 1, 1, 0, 0
	for a > 0 {
		var q int64 = b / a
		x, u = u, x-q*u
		y, v = v, y-q*v
		b, a = a, b-q*a
	}
	return b, x, y
}
