//Package sorts a package for demonstrating sorting algorithms in Go
package sorts

func selectionSort(arr []int) []int {
	var min int = 0
	var tmp int = 0

	for i := 0; i < len(arr); i++ {
		min = i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}

		tmp = arr[i]
		arr[i] = arr[min]
		arr[min] = tmp
	}
	return arr
}
