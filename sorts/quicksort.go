//Package sorts a package for demonstrating sorting algorithms in Go
package sorts

func partition(arr []int, low, high int) int {
	index := low - 1
	pivotElement := arr[high]
	for i := low; i < high; i++ {
		if arr[i] <= pivotElement {
			index += 1
			arr[index], arr[i] = arr[i], arr[index]
		}
	}
	arr[index+1], arr[high] = arr[high], arr[index+1]
	return index + 1
}

func QuickSort(arr []int, low, high int) {
	if len(arr) <= 1 {
		return
	}

	if low < high {
		pivot := partition(arr, low, high)
		QuickSort(arr, low, pivot-1)
		QuickSort(arr, pivot+1, high)
	}
}
