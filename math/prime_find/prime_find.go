package main

// Finds all prime numbers from 1 to max implemented using eratosthenes sieve
func PrimesWithMax(max int) []int {
	var lp []int // stores minimum prime divisor for every number from 2 to max
	lp = make([]int, max+1)

	var primes []int

	for i := 2; i <= max; i++ {
		if lp[i] == 0 {
			lp[i] = i
			primes = append(primes, i)
		}
		for j := 0; j < len(primes) && primes[j] <= lp[i] && i*primes[j] <= max; j++ {
			lp[i*primes[j]] = primes[j]
		}
	}

	return primes
}