//Package sorts a package for demonstrating sorting algorithms in Go
package sorts

import (
	"math/rand"
)

// QuickSort using random pivot
func QuickSort(arr []int) []int {

	if len(arr) <= 1 {
		return arr
	}

	median := arr[rand.Intn(len(arr))]

	lowPart := make([]int, 0, len(arr))
	highPart := make([]int, 0, len(arr))
	middlePart := make([]int, 0, len(arr))

	for _, item := range arr {
		switch {
		case item < median:
			lowPart = append(lowPart, item)
		case item == median:
			middlePart = append(middlePart, item)
		case item > median:
			highPart = append(highPart, item)
		}
	}

	lowPart = QuickSort(lowPart)
	highPart = QuickSort(highPart)

	lowPart = append(lowPart, middlePart...)
	lowPart = append(lowPart, highPart...)

	return lowPart
}

func partition(arr []int, lo int, hi int) int {
	i := lo - 1
	j := hi + 1
	pivot := arr[(lo+hi)/2]

	for true {
		i++
		for arr[i] < pivot {
			i++
		}
		j--
		for arr[j] > pivot {
			j--
		}

		if i >= j {
			return j
		}

		arr[i], arr[j] = arr[j], arr[i]
	}

	// unreachable
	return j
}

func QuickSortHoareHelper(arr []int, lo int, hi int) []int {
	for hi > lo {
		p := partition(arr, lo, hi)
		QuickSortHoareHelper(arr, lo, p)
		lo = p + 1
	}

	return arr
}

// QuickSortHoare using Hoare partition scheme and centered pivot
func QuickSortHoare(arr []int) []int {
	return QuickSortHoareHelper(arr, 0, len(arr)-1)
}
