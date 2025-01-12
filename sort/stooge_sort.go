// implementation of the Stooge sort
// more info at https://en.wikipedia.org/wiki/Stooge_sort
// worst-case time complexity    O(n^2.709511)
// worst-case space complexity   O(n)

package sort

import (
	"github.com/TheAlgorithms/Go/constraints"

	// math imported for floor division
	"math"
)

func innerStooge[T constraints.Ordered](arr []T, i int32, j int32) []T {
	if arr[i] > arr[j] {
		arr[i], arr[j] = arr[j], arr[i]
	}
	if (j - i + 1) > 2 {
		t := int32(math.Floor(float64(j-i+1) / 3.0))
		arr = innerStooge(arr, i, j-t)
		arr = innerStooge(arr, i+t, j)
		arr = innerStooge(arr, i, j-t)
	}
	return arr
}

func Stooge[T constraints.Ordered](arr []T) []T {
	if len(arr) == 0 {
		return arr
	}

	return innerStooge(arr, 0, int32(len(arr)-1))
}
