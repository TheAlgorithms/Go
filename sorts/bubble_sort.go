//Package sorts a package for demonstrating sorting algorithms in Go
package sorts

func bubbleSort(array []int) []int {
	swapped := true
	for swapped {
		swapped = false
		for i := 0; i < len(array)-1; i++ {
			if array[i+1] < array[i] {
				array[i+1], array[i] = array[i], array[i+1]
				swapped = true
			}
		}
	}
	return array
}
