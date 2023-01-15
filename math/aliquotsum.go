// aliquotsum.go
// description: Returns s(n)
// details:
// the aliquot sum s(n) of a positive integer n
// is the sum of all proper divisors of n,
// that is, all divisors of n other than n itself
// wikipedia: https://en.wikipedia.org/wiki/Aliquot_sum
// author: Akshay Dubey (https://github.com/itsAkshayDubey)
// see aliquotsum_test.go

package math

// This function returns s(n) for given number
func AliquotSum(n int) (int, error) {
	switch {
	case n < 0:
		return 0, ErrPosArgsOnly
	case n == 0:
		return 0, ErrNonZeroArgsOnly
	case n == 1:
		return 0, nil
	default:
		var sum int
		for i := 1; i <= n/2; i++ {
			if n%i == 0 {
				sum += i
			}
		}
		return sum, nil
	}
}
