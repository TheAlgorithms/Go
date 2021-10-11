package main

import (
	"fmt"
	"math"
)

func sqrt(value int64) int64 {
	return int64(math.Sqrt(float64(value)))
}

func isPrime(value int64) bool {
	var i int64 = 2
	for ; i <= int64(math.Floor(float64(value)/2)); i++ {
		if value%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	const num int64 = 600851475143
	for i := sqrt(num); i >= 1; i-- {
		if num%i == 0 && isPrime(i) {
			fmt.Println(i)
			break
		}
	}
}
