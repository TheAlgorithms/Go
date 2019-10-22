package main

import (
	"fmt"
)

func gcd(a, b uint64) uint64 {
	if a == 0 {
		return b
	}
	return gcd(b%a, a)
}

func main() {
	var a, b uint64
	fmt.Println("Please input first integer: ")
	_, err := fmt.Scanf("%d", &a)
	if err != nil {
		fmt.Printf("wrong input number, error: %s", err.Error())
	}

	fmt.Println("Please input second integer: ")
	_, err = fmt.Scanf("%d", &b)
	if err != nil {
		fmt.Printf("wrong input number, error: %s", err.Error())
	}

	fmt.Printf("Greatest common divisor of %d and %d is: %d\n", a, b, gcd(a, b))
}
