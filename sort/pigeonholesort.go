package sort

// PigeonholeSort algorithm's working at wikipedia.
// https://en.wikipedia.org/wiki/Pigeonhole_sort
func PigeonholeSort(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}

	max := arr[0]
	min := arr[0]

	for _, element := range arr {
		if min > element {
			min = element
		}
		if max < element {
			max = element
		}
	}

	size := max - min + 1

	holes := make([]int, size)

	for _, element := range arr {
		holes[element-min]++
	}

	i := 0

	for j := 0; j < size; j++ {
		for holes[j] > 0 {
			holes[j]--
			arr[i] = j + min
			i++
		}
	}

	return arr
}
