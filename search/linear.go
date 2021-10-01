package search

// Linear Simple linear search algorithm that iterates over all elements of an array in the worst case scenario
func Linear(array []int, query int) (int, error) {
	for i, item := range array {
		if item == query {
			return i, nil
		}
	}
	return -1, ErrNotFound
}
