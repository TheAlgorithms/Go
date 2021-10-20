// countingsort.go
// description: Implementation of counting sort algorithm
// details: A simple counting sort algorithm implementation
// author [Phil](https://github.com/pschik)
// see sort_test.go for a test implementation, test function TestQuickSort

// package sort provides primitives for sorting slices and user-defined collections
package sort

import (
	"fmt"
	"runtime"
	"strings"
)

func Count(data []int) []int {
	var aMin, aMax = -1000, 1000
	count := make([]int, aMax-aMin+1)
	for _, x := range data {
		count[x-aMin]++
	}
	z := 0
	for i, c := range count {
		for c > 0 {
			data[z] = i + aMin
			z++
			c--
		}
	}
	return data
}
