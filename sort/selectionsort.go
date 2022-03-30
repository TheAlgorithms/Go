package sort

import "github.com/TheAlgorithms/Go/constraints"

func Selection[T constraints.Ordered](arr []T) []T {
	for i := 0; i < len(arr); i++ {
		min := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}

		arr[i], arr[min] = arr[min], arr[i]
	}
	return arr
}
