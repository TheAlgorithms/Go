// Find the minimum and maximum numbers of a list.
// Find the smallest and largest distance from 0.
// Find the range.

package main

import (
	"fmt"
)

// absolute value, also implemented in "math" as math.Abs
func abs(a float64) float64 {
	if a < 0 {
		return -a
	}
	return a
}

// find the minimum
func min(numbers []float64) float64 {
	min := numbers[0]
	for _, n := range numbers {
		if n < min {
			min = n
		}
	}
	return min
}

//find the smallest distance from 0
func minDist(numbers []float64) float64 {
	min := numbers[0]
	for _, n := range numbers {
		if abs(n) < abs(min) {
			min = n
		}
	}
	return min
}

// find the maximum
func max(numbers []float64) float64 {
	max := numbers[0]
	for _, n := range numbers {
		if n > max {
			max = n
		}
	}
	return max
}

//find the largest distance from 0
func maxDist(numbers []float64) float64 {
	max := numbers[0]
	for _, n := range numbers {
		if abs(n) > abs(max) {
			max = n
		}
	}
	return max
}

func findRange(numbers []float64) float64 {
	return max(numbers) - min(numbers)
}

func main() {
	numbers := []float64{-4.6, 58.003, 32, -2, 4.1, -120, 65.5431, -30.7}
	fmt.Println("Min : ", min(numbers))
	fmt.Println("Min distance from 0 : ", minDist(numbers))
	fmt.Println("Max : ", max(numbers))
	fmt.Println("Max distance from 0 : ", maxDist(numbers))
	fmt.Println("Range : ", findRange(numbers))
}
