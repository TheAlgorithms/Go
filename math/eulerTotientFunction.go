/*
In number theory, Euler's totient function counts the positive integers up to a given integer n that are relatively prime to n. 
It is written using the Greek letter phi as φ(n) or ϕ(n), and may also be called Euler's phi function. 
In other words, it is the number of integers k in the range 1 ≤ k ≤ n for which the greatest common divisor gcd(n, k) is equal to 1.
The integers k of this form are sometimes referred to as totatives of n
Source -Wikipedia https://en.wikipedia.org/wiki/Euler's_totient_function#Euler.27s_product_formula
 */

package main

func getPrimeFactors(n int, primes []int) map[int]int {
	primeFacts := make(map[int]int) 
	for n != 1 {
		for i := 0; i < len(primes); i ++ {
			if n % primes[i] == 0 {
				val, ok := primeFacts[primes[i]]
				if !ok {
					val = 0
				}
				primeFacts[primes[i]] = val + 1
				n = n/primes[i]
				break
			}
		}
	}
	return primeFacts
}

func getPrimes(N int) []int {
	isComposite := make([]bool, N)
	primes := []int{}
	for i := 2; i < N; i ++ {
		if !isComposite[i] {
			primes = append(primes, i)
			for x := i+i; x < N; x += i {
				isComposite[x] = true
			}
		}
	}
	return primes
}

func totient(n int, primes []int) int {
	primeFacts := getPrimeFactors(n, primes)

	ans := n

	for prime := range primeFacts {
		ans = ans*(prime-1)/prime
	}
	return ans
}
