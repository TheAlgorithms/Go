package sorts

// CountingSort : Time O(n) | Space O(n)
func CountingSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	min, max := findMinAndMax(arr)

	buckets := make([]int, max-min+1)

	for _, item := range arr {
		buckets[item-min]++
	}

	index := 0
	for i := 0; i < len(buckets); i++ {
		for buckets[i] > 0 {
			arr[index] = i + min
			index++
			buckets[i]--
		}
	}
	return arr
}

func findMinAndMax(a []int) (min int, max int) {
	min = a[0]
	max = a[0]
	for _, value := range a {
		if value < min {
			min = value
		}
		if value > max {
			max = value
		}
	}
	return min, max
}
