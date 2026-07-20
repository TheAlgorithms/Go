// aliquotsum.go
// description: Returns s(n)
// details:
// the aliquot sum s(n) of a positive integer n
// is the sum of all proper divisors of n,
// that is, all divisors of n other than n itself
// wikipedia: https://en.wikipedia.org/wiki/Aliquot_sum
// time complexity: O(sqrt(n))
// space complexity: O(1)
// author: Akshay Dubey (https://github.com/itsAkshayDubey)
// see aliquotsum_test.go

package math

// AliquotSum returns the sum of all proper divisors of n.
func AliquotSum(n int) (int, error) {
	switch {
	case n < 0:
		return 0, ErrPosArgsOnly
	case n == 0:
		return 0, ErrNonZeroArgsOnly
	default:
		if n == 1 {
			return 0, nil
		}

		sum := 1
		for divisor := 2; divisor <= n/divisor; divisor++ {
			if n%divisor == 0 {
				sum += divisor

				pairedDivisor := n / divisor
				if pairedDivisor != divisor {
					sum += pairedDivisor
				}
			}
		}
		return sum, nil
	}
}
