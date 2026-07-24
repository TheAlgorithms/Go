package lcm

import (
	"github.com/TheAlgorithms/Go/math/gcd"
)

// absInt64 returns the absolute value of x as an int64.
func absInt64(x int64) int64 {
	if x < 0 {
		return -x
	}
	return x
}

// Lcm returns the lcm of two numbers using the fact that lcm(a,b) * gcd(a,b) = | a * b |
//
// The computation is performed with integer arithmetic and the division by gcd
// is applied to |a| before multiplying by |b|, so the running intermediate never
// exceeds the true LCM. This avoids two failure modes of the previous
// float64-based formula:
//   - int64 multiply overflow: a*b wraps mod 2^64 before the float64 cast, so
//     e.g. Lcm(4294967296, 8589934592) returned 0 (impossible for positives).
//   - float64 precision loss: float64 cannot represent every int64 past 2^53,
//     so e.g. Lcm(9007199254740993, 1) returned 9007199254740992, violating the
//     identity lcm(x, 1) == x.
func Lcm(a, b int64) int64 {
	g := gcd.Iterative(a, b)
	if g == 0 {
		// gcd.Iterative(0, 0) == 0; LCM(0, x) is conventionally 0.
		return 0
	}
	return (absInt64(a) / absInt64(g)) * absInt64(b)
}
