package prime

// A primality test is an algorithm for determining whether an input number is prime.Among other fields of mathematics, it is used for cryptography.
//Unlike integer factorization, primality tests do not generally give prime factors, only stating whether the input number is prime or not.
//Source - Wikipedia https://en.wikipedia.org/wiki/Primality_test

// NaiveApproach checks if an integer is prime or not. Returns a bool.
func NaiveApproach(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i < n; i++ {

		if n%i == 0 {
			return false
		}
	}
	return true
}

// PairApproach checks primality of an integer and returns a bool. More efficient than the naive approach as number of iterations are less.
func PairApproach(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
