// Implementation of counting sort algorithm
// Reference: https://www.geeksforgeeks.org/counting-sort/

package sort

func countingSort(data []int) []int {
	if len(data) == 0 {
		return data
	}
	maxValue := data[0]
	minValue := data[0]
	for i := 0; i < len(data); i++ {
		if data[i] > maxValue {
			maxValue = data[i]
		}
		if data[i] < minValue {
			minValue = data[i]
		}
	}
	rangeOfElements := maxValue - minValue + 1
	countArray := make([]int, rangeOfElements, rangeOfElements)
	outputArray := make([]int, len(data))

	for i := 0; i < len(data); i++ {
		countArray[data[i]-minValue] += 1
	}

	for i := 1; i < len(countArray); i++ {
		countArray[i] += countArray[i-1]
	}

	for i := len(data) - 1; i >= 0; i-- {
		outputArray[countArray[data[i]-minValue]-1] = data[i]
		countArray[data[i]-minValue] -= 1
	}
	return outputArray
}
