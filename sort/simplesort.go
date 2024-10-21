// simplesort.go
// description: Implementation of a simple sorting algorithm
// details:
// A simple sorting algorithm that look counter intuitive at first glance and very similar to Exchange Sort
// An improved version is included with slight changes to make the sort slightly more efficient
// reference: https://arxiv.org/abs/2110.01111v1
// see sort_test.go for a test implementation, test function TestSimple and TestImprovedSimple
// worst-case time complexity: O(n^2)
// average-case time complexity: O(n^2)
// space complexity: O(1)

package sort

import "github.com/TheAlgorithms/Go/constraints"

func Simple[T constraints.Ordered](arr []T) []T {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			if arr[i] < arr[j] {
				// swap arr[i] and arr[j]
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	return arr
}

// ImprovedSimple is a improve SimpleSort by skipping an unnecessary comparison of the first and last.
// This improved version is more similar to implementation of insertion sort
func ImprovedSimple[T constraints.Ordered](arr []T) []T {
	for i := 1; i < len(arr); i++ {
		for j := 0; j < len(arr)-1; j++ {
			if arr[i] < arr[j] {
				// swap arr[i] and arr[j]
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	return arr
}
