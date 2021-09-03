//Package sorts a package for demonstrating sorting algorithms in Go
package sorts

func QuickSort(array []int) []int {
	if len(array) <= 1 {
		return array
	}
	q := partition(array)
	QuickSort(array[:q])
	QuickSort(array[q+1:])
	return array
}

func partition(array []int) int {
	r := len(array) - 1
	x := array[r] // last element of slice
	i := -1
	for j := 0; j < r; j++ {
		if array[j] <= x {
			i++
			array[i], array[j] = array[j], array[i]
		}
	}
	array[i+1], array[r] = array[r], array[i+1]
	return i + 1
}
