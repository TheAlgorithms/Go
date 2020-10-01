// Formula from https://en.wikipedia.org/wiki/Least_common_multiple
// Find the least common multiple of two integers
// by using the greatest common divisor.

package main

import (
	"fmt"
)

// abs: absolute value for an integer
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

// gcd: greatest common divisor -
// the largest positive integer that divides both a and b
func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

// lcm: least common multiple -
// the smallest positive integer that is divisible by both a and b
func lcm(a, b int) int {
	if a == 0 && b == 0 {
		return 0
	}
	a, b = abs(a), abs(b)
	return a * b / gcd(a, b)
}

func main() {
	fmt.Println("Least common multiple of 21 and 6 : ", lcm(21, 6))
	fmt.Println("Least common multiple of -21 and 6 : ", lcm(-21, 6))
	fmt.Println("Least common multiple of -6 and -21 : ", lcm(-6, -21))
	fmt.Println("Least common multiple of 21 and -6 : ", lcm(21, -6))
	// all should equal 42

	fmt.Println("Least common multiple of 18 and 0 : ", lcm(18, 0))
	fmt.Println("Least common multiple of 0 and -7 : ", lcm(0, -7))
	fmt.Println("Least common multiple of 0 and 0 : ", lcm(0, 0))
	// all should equal 0
}
