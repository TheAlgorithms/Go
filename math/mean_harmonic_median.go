// Functions to find different types of averages in a list of floats.
// The list is sorted, where needed, with the help of the "sort" library.
// See Go/sorts for sorting algorithms.

package main

import (
	"fmt"
	"sort"
)

// find the mean of the list
// (sum of elements / amount of elements)
func mean(numbers []float64) float64 {
	var sum float64 = 0
	for _, v := range numbers {
		sum += v
	}
	return sum / float64(len(numbers))
}

// find the harmonic mean of the list
// (amount of elements / sum of reciprocals of elements)
func meanHarmonic(numbers []float64) float64 {
	var sum float64 = 0
	for _, v := range numbers {
		sum += 1 / v
	}
	return float64(len(numbers)) / sum
}

// find the median of the list -
// the number in the middle of the sorted list if odd number of elements
// else the mean of the two numbers in the middle of the sorted list
func median(numbers []float64) float64 {
	sort.Float64s(numbers)
	n := len(numbers)
	if middle := numbers[n/2-1 : n/2+1]; n%2 == 0 {
		return mean(middle)
		//return (numbers[n/2-1] + numbers[n/2]) / 2
	}
	return numbers[n/2]
}

func main() {
	numbers := []float64{4, 2, 1, 4, 3, 5, 3, 5}
	for i := 0; i < 2; i++ {
		fmt.Printf("Mean of %v is %v \n", numbers, mean(numbers))
		fmt.Printf("Harmonic mean of %v is %v \n", numbers, meanHarmonic(numbers))
		fmt.Printf("Median of %v is %v \n", numbers, median(numbers))
		fmt.Println()
		numbers = numbers[:len(numbers)-1]
	}
}
