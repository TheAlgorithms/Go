package sort

import "github.com/TheAlgorithms/Go/constraints"

func heapify[T constraints.Ordered](slice []T, N, i int) {
	largest := i
	l := 2*i + 1
	r := 2*i + 2

	if l < N && slice[largest] < slice[l] {
		largest = l
	}
	if r < N && slice[largest] < slice[r] {
		largest = r
	}

	// change root, if needed
	if largest != i {
		slice[i], slice[largest] = slice[largest], slice[i]

		// heapify the root
		heapify(slice, N, largest)
	}
}

func HeapSort[T constraints.Ordered](slice []T) []T {
	N := len(slice)

	// build a maxheap
	for i := N/2 - 1; i >= 0; i-- {
		heapify(slice, N, i)
	}

	for i := N - 1; i > 0; i-- {
		slice[i], slice[0] = slice[0], slice[i]
		heapify(slice, i, 0)
	}

	return slice
}
