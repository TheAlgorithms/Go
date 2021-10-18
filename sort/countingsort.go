// countingsort.go
// description: Implementation of counting sort algorithm
// details: A simple counting sort algorithm implementation
// author [Phil](https://github.com/pschik)
// see sort_test.go for a test implementation, test function TestQuickSort

// package sort provides primitives for sorting slices and user-defined collections
package sort

func Count(data []int, maxValue int) []int {
	bucketLen := maxValue + 1
	bucket := make([]int, bucketLen)
	sortedIndex := 0
	length := len(data)
	for i := 0; i < length; i++ {
		bucket[data[i]] += 1
	}
	for j := 0; j < bucketLen; j++ {
		for bucket[j] > 0 {
			data[sortedIndex] = j
			sortedIndex += 1
			bucket[j] -= 1
		}
	}
	return data
}
