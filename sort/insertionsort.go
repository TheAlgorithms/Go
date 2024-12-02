// insertionsort.go
// description: Implementation of insertion sort algorithm
// worst-case time complexity: O(n^2)
// average-case time complexity: O(n^2)
// space complexity: O(1)

package sort

import "github.com/TheAlgorithms/Go/constraints"

func Insertion[T constraints.Ordered](arr []T) []T {
	for currentIndex := 1; currentIndex < len(arr); currentIndex++ {
		temporary := arr[currentIndex]
		iterator := currentIndex
		for ; iterator > 0 && arr[iterator-1] > temporary; iterator-- {
			arr[iterator] = arr[iterator-1]
		}
		arr[iterator] = temporary
	}
	return arr
}
