package main

import "fmt"

func reversed(value int) int {
	reversed := 0
	for ; value > 0; value = value / 10 {
		reversed = value%10 + 10*reversed
	}
	return reversed
}

func isPalindrome(value int) bool {
	return value == reversed(value)
}

func main() {
	largestPalindrome := 0
	x := 100
	for ; x <= 999; x = x + 1 {
		y := 100
		for ; y <= 999; y = y + 1 {
			if z := x * y; isPalindrome(z) && z > largestPalindrome {
				largestPalindrome = z
			}
		}
	}
	fmt.Println(largestPalindrome)
}
