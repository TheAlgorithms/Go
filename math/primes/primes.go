package primes

// package main

// import "fmt"

func primes(max int) []int {
	max++
	numbers := make([]bool, max)
	ret := make([]int, max)
	length := 0
	for i := 2; i < max; i++ {
		if !numbers[i] {
			for j := i; j < max; j += i {
				numbers[j] = true
			}
			ret[length] = i
			length++
		}
	}
	return ret[:length]
}

// func main() {
// 	fmt.Println(primes(101))
// }
