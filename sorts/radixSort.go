package main

import (
	"fmt"
	"strconv"
)


func findLargestNum(array []int) int {
	largestNum := 0

	for i := 0; i < len(array); i++ {
		if array[i] > largestNum {
			largestNum = array[i]
		}
	}
	return largestNum
}


func radixSort(array []int) []int {

  fmt.Println("Running Radix Sort on Unsorted List\n")


	largestNum := findLargestNum(array)
	size := len(array)
	significantDigit := 1
	semiSorted := make([]int, size, size)


	for largestNum / significantDigit > 0 {
  
    fmt.Println("\tSorting: " + strconv.Itoa(significantDigit) + "'s place", array)

		bucket := [10]int{0}


		for i := 0; i < size; i++ {
			bucket[(array[i] / significantDigit) % 10]++
		}




		for i := 1; i < 10; i++ {
			bucket[i] += bucket[i - 1] 
		}


		for i := size - 1; i >= 0; i-- {
			bucket[(array[i] / significantDigit) % 10]--
			semiSorted[bucket[(array[i] / significantDigit) % 10]] = array[i]
		}


		for i := 0; i < size; i++ {
			array[i] = semiSorted[i]
		}
    
    fmt.Println("\n\tBucket: ", bucket)
  

		significantDigit *= 10 
	}

	return array
}

func main() {

  fmt.Println("\n\nRunning Radix Sort Example in Go!")

	unsortedList :=[]int{10, 2, 303, 4021, 293, 1, 0, 429, 480, 92, 2999, 14}
	fmt.Println("Unsorted List:", unsortedList, "\n")

	sortedList := radixSort(unsortedList)

	fmt.Println("\nSorted List:", sortedList, "\n")
