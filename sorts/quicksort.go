//Package sorts a package for demonstrating sorting algorithms in Go
package sorts

func Quicksort(array []int) {
	if len(array) <= 1 {
		return
	}
	q := partition(array)
	Quicksort(array[:q])
	Quicksort(array[q+1:])
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
