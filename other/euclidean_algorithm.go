/*
Euclidean Algorithm

e.g.: go build euclidean_algorithm.go
      ./euclidean_algorithm 20 6
*/

package main

import (
	"fmt"
	"os"
	"strconv"
)

func euclidean(x, y, verbose int) int {
	if verbose != 0 {
		fmt.Println(x, y) /* Print table if verbose is true */
	}
	if y == 0 {
		fmt.Println("gcd =", x)
		return x
	}
	return euclidean(y, x%y, verbose)
}

func main() {
	var x int
	var y int
	x, _ = strconv.Atoi(os.Args[1])
	y, _ = strconv.Atoi(os.Args[2])
	euclidean(x, y, 1)
}
