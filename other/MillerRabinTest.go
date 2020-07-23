package main

import (
	"fmt"
	"math/rand"
)


// Miller–Rabin primality test

func power(x int, y int, p int) int {
	res := 1
	x = x % p

	for y > 0 {
		if (y & 1) == 1 {
			res = (res * x) % p
		}
		y = y >> 1
		x = (x*x) % p
	}
	return res
}

func millerTest(d int, n int) bool {
	var a =  2 + (rand.Int() % (n-4))
	x := power(a, d, n)

	if x == 1 || x == n - 1 {
		return true
	}

	for d != n -1 {
		x = (x*x) % n
		d *= 2

		if x == 1 {
			return false
		}
		if x == n -1 {
			return true
		}
	}
	return false
}

func isPrimeNumber(n int, k int) bool {
	if n <= 1 || n == 4 {
		return false
	}
	if n <= 3 {
		return true
	}

	d := n - 1
	for d % 2 == 0 {
		d /= 2
	}

	for i := 0; i < k; i++ {
		if millerTest(d, n) == false {
			return false
		}
	}

	return true
}

func main() {
	k := 4
	fmt.Println("All primes smaller than 100:")
	for i := 0; i < 100; i++ {
		if isPrimeNumber(i, k) {
			fmt.Println(i)
		}
	}
}