//Package sorts a package for demonstrating sorting algorithms in Go
package sorts

import (
	"math/rand"
	"time"
)

func bogoSort(arr []int) []int {

	rand.Seed(time.Now().UnixNano())

	result := make([]int, len(arr))
	copy(result, arr)

	for !isSorted(arr) {
		rand.Shuffle(len(arr), func(i, j int) {
			arr[i], arr[j] = arr[j], arr[i]
		})
	}

	return arr
}

func isSorted(arr []int) bool {
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			return false
		}
	}

	return true
}
