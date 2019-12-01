//Package sorts a package for demonstrating sorting algorithms in Go
package sorts

func insertionSort(arr []int) {
	for out := 1; out < len(arr); out++ {
		temp := arr[out]
		in := out

		for ; in > 0 && arr[in-1] >= temp; in-- {
			arr[in] = arr[in-1]
		}
		arr[in] = temp
	}
	return arr
}
