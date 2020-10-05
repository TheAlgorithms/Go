// Package sorts a package for demonstrating sorting algorithms in Go
// Implementation of stooge sort algorithm
// Reference: https://en.wikipedia.org/wiki/Stooge_sort

package sorts

func swap(arr []int, l int, h int) {
	tmp := arr[l]
	arr[l] = arr[h]
	arr[h] = tmp
}

// At the initial run, the l value is 0 and the h is the length of the array - 1
func stoogeSort(arr []int, l int, h int) []int {
	if l >= h {
		return arr
	}
	if arr[l] > arr[h] {
		swap(arr, l, h)
	}
	if h-l+1 > 2 {
		t := (h - l + 1) / 3
		stoogeSort(arr, l, h-t)
		stoogeSort(arr, l+t, h)
		stoogeSort(arr, l, h-t)
	}

	return arr
}
