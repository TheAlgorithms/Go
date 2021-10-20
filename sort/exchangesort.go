// Implementation of exchange sort algorithm, a variant of bubble sort
// Reference: https://en.wikipedia.org/wiki/Sorting_algorithm#Exchange_sort

package sort

func Exchange(arr []int) []int {
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}

	return arr
}
