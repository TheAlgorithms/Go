package main

import  (
	"fmt"
	"os"
	"strconv"
)

var (
	m map[uint64]uint64
)

func fibonacciRecursionSlow(n uint64) uint64 {
	if n <= 2 {
		return 1
	}
	return  fibonacciRecursionSlow(n - 1) +  fibonacciRecursionSlow(n - 2)
}

func fibonacciRecursionFast(n uint64) uint64 {
	if _, ok := m[n]; ok {
		return m[n]
	}
	
	m[n] =  fibonacciRecursionFast(n - 1) + m[n - 2]
	return m[n]
}

func main() {
	n, err := strconv.ParseUint(os.Args[1], 10, 64)
	if err != nil {
		panic(err)
	}

	m = map[uint64]uint64{1: 1, 2: 1}
	
	fmt.Printf("Fast - %d\n", fibonacciRecursionFast(n))
	fmt.Printf("Slow - %d\n", fibonacciRecursionSlow(n))
}