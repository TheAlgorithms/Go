package search

// Linear Simple linear search algorithm that iterates over all elements of an array in the worst case scenario
func Linear(array []int, query int) int {
	for i, item := range array {
		if item == query {
			return i
		}
	}
	return -1
}
