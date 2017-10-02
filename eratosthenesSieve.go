package main

import (
	"fmt"
)

//Only works for primes smaller or equal to 10e7
func sieve(upperBound int64) []int64 {
	_sieveSize := upperBound + 10
	//Creates set to mark wich numbers are primes and wich are not
	//true: not primes, false: primes
	//this to favor default initialization of arrays in go
	var bs [10000010]bool
	//creates a slice to save the primes it finds
	primes := make([]int64, 0, 1000)
	
	bs[0] = true
	bs[1] = true
	//iterate over the numbers set
	for i := int64(0); i <= _sieveSize; i++ {
		//if find one number that is not marked as a compund number, mark all its multiples
		if !bs[i] {
			for j := i * i; j <= _sieveSize; j += i {
				bs[j] = true
			}
			//Add the prime you just find to the slice of primes
			primes = append(primes, i)
		}
	}
	return primes
}

func main() {
	//prints first N primes into console
	N := 100
	primes := sieve(N)
	fmt.Println(primes)
}
