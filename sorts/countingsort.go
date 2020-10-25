// Reference: https://en.wikipedia.org/wiki/Counting_sort

package sorts

func countingSort(arr []int) []int {
	if len(arr) == 0 {
		return make([]int, 0)
	}
	minElement := arr[0]
	maxElement := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] > maxElement {
			maxElement = arr[i]
		}
		if minElement > arr[i] {
			minElement = arr[i]
		}
	}

	maxElement -= minElement

	counts := make([]int, maxElement+1)
	result := make([]int, 0)
	for _, item := range arr {
		counts[item-minElement]++
	}
	for i := 0; i <= maxElement; i++ {
		for j := 0; j < counts[i]; j++ {
			result = append(result, i+minElement)
		}
	}
	return result
}
