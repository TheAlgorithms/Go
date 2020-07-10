package main

import(
	"fmt"
)

func SieveOfEratosthenes(n int) []int {
	isPrime := []bool{}
	primes := []int{}
	for i := 0; i <= n; i++{
		isPrime = append(isPrime, true)
	}
	for p := 2; p * p <= n; p++{ 
        if isPrime[p] {
            for i := p * p; i <= n; i += p{
                isPrime[i] = false
			} 
        } 
	}

	for p := 2; p <= n; p++{ 
		if isPrime[p]{
			primes = append(primes, p)
		}
	}
	return primes
}

func main() {
	fmt.Println(SieveOfEratosthenes(50))
} 
