package main

import "fmt"

func main() {
	sum := 0
	fib := []int{1, 1, 2}
	for fib[2] < 4000000 {
		if fib[2]%2 == 0 {
			sum += fib[2]
		}
		fib[0] = fib[1]
		fib[1] = fib[2]
		fib[2] = fib[0] + fib[1]
	}
	fmt.Print(sum)
}
