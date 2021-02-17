package sieve

import (
	"testing"
)

func TestSieve(t *testing.T) {
	firstTenPrimes := [10]int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29}

	t.Run("First 10 primes test", func(t *testing.T) {
		var testTenPrimes [10]int

		ch := make(chan int)
		go Generate(ch)

		for i := 0; i < 10; i++ {
			testTenPrimes[i] = <-ch
			ch1 := make(chan int)
			go Sieve(ch, ch1, testTenPrimes[i])
			ch = ch1
		}

		if firstTenPrimes != testTenPrimes {
			t.Errorf("The first 10 primes do not match")
		}

	})
}
