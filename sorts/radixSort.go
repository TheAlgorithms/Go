//Radix sort is an integer sorting algorithm that sorts data with integer keys by grouping the keys by individual digits
//that share the same significant position and value (place value).
//Radix sort uses counting sort as a subroutine to sort an array of numbers.
//Source Wikipedia - https://en.wikipedia.org/wiki/Radix_sort
package main

func LargetSum(array []int) int {
	largestNum := 0

	for i := 0; i < len(array); i++ {
		if array[i] > largestNum {
			largestNum = array[i]
		}
	}
	return largestNum
}

func radixSort(array []int) []int {
	largestNum := LargetSum(array)
	size := len(array)
	significantDigit := 1
	semiSorted := make([]int, size, size)
	for largestNum / significantDigit > 0 {
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
		significantDigit *= 10
	}
	return array
}
