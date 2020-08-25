//Package sorts a package for demonstrating sorting algorithms in Go
package sorts

import (
	"math/rand"
	"time"
)

func quickSort(array []int, left int, right int) []int {
	if left < right {
		p := partition(array, left, right)
		quickSort(array, left, p-1)
		quickSort(array, p+1, right)
	}

	return array
}

// Lomuto's partition scheme
func partition(array []int, low int, high int) int {
	rand.Seed(time.Now().UnixNano())
	pivotIndex := rand.Intn(high-low) + low
	pivot := array[pivotIndex]

	array[pivotIndex], array[high] = array[high], array[pivotIndex]

	index := low

	for i := low; i < high; i++ {
		if array[i] < pivot {
			array[i], array[index] = array[index], array[i]
			index++
		}
	}

	array[high], array[index] = array[index], array[high]

	return index
}
