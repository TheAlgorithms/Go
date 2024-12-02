// primefactorization.go
// description: Prime factorization of a number
// time complexity: O(sqrt(n))
// space complexity: O(sqrt(n))

package prime

// Factorize is a function that computes the exponents
// of each prime in the prime factorization of n
func Factorize(n int64) map[int64]int64 {
	result := make(map[int64]int64)

	for i := int64(2); i*i <= n; i += 1 {
		for {
			if n%i != 0 {
				break
			}
			result[i] += 1
			n /= i
		}

	}
	if n > 1 {
		result[n] += 1
	}
	return result
}
