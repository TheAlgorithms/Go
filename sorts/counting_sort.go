package sorts

func CountingSort(arr []int, max int) []int {
	//use the max value to create a slice of item counts
	counts := make([]int, max+1)
	result := make([]int, 0)
	for _, item := range arr {
		counts[item]++
	}
	for i := 0; i < len(counts); i++ {
		count := counts[i]
		for j := 0; j < count; j++ {
			result = append(result, i)
		}
	}
	return result
}
