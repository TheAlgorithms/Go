/*
Extended Euclidean Algorithm

the function extendedEuclidean return gcd, alpha and beta

e.g.: go build extended_euclidean_algorithm.go
      ./extended_euclidean_algorithm 20 6
*/

package main

import (
	"fmt"
	"os"
	"strconv"
)

func extendedEuclidean(x, y int) (gcd, alpha, beta int) {
	oldAlpha, oldBeta := 1, 0
	alpha, beta = 0, 1
	for y != 0 {
		quotient := x / y
		rest := x % y
		newAlpha := oldAlpha - alpha*quotient
		newBeta := oldBeta - beta*quotient

		x = y
		y = rest
		oldAlpha = alpha
		oldBeta = beta
		alpha = newAlpha
		beta = newBeta
	}
	gcd = x
	alpha = oldAlpha
	beta = oldBeta
	return gcd, alpha, beta
}

func main() {
	var x int
	var y int
	x, _ = strconv.Atoi(os.Args[1])
	y, _ = strconv.Atoi(os.Args[2])
	gcd, alpha, beta := extendedEuclidean(x, y)
	fmt.Println("gcd =", gcd)
	fmt.Println("alpha =", alpha)
	fmt.Println("beta =", beta)
}
