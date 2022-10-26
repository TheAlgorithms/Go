// bogosort_test.go
// description: Test implementations for bogosort.go
// details:
// Test case implementationfor the bogosort algorithm
// author(s): [Focusucof](https://github.com/Focusucof)
// see bogosort.go

package sort

import (
	"reflect"
	"testing"
)

func TestBogo(t *testing.T) {
	sortTests := []struct {
		input    []int
		expected []int
		name     string
	}{
		//Sorted slice
		{
			input:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			expected: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			name:     "Sorted Unsigned",
		},
		//Reversed slice
		{
			input:    []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
			expected: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			name:     "Reversed Unsigned",
		},
		//Sorted slice
		{
			input:    []int{-3, -2, -1, 0, 1, 2, 3},
			expected: []int{-3, -2, -1, 0, 1, 2, 3},
			name:     "Sorted Signed",
		},
		//Reversed slice
		{
			input:    []int{3, 2, 1, 0, -1, -2, -3},
			expected: []int{-3, -2, -1, 0, 1, 2, 3},
			name:     "Reversed Signed",
		},
		//Reversed slice, even length
		{
			input:    []int{3, 2, 1, -1, -2, -3},
			expected: []int{-3, -2, -1, 1, 2, 3},
			name:     "Reversed Signed #2",
		},
		//Random order with repetitions
		{
			input:    []int{4, -2, 3, 2, -1, 0, -3, -4, 1},
			expected: []int{-4, -3, -2, -1, 0, 1, 2, 3, 4},
			name:     "Random order Signed",
		},
		//Single-entry slice
		{
			input:    []int{1},
			expected: []int{1},
			name:     "Singleton",
		},
		// Empty slice
		{
			input:    []int{},
			expected: []int{},
			name:     "Empty Slice",
		},
	}
	for _, test := range sortTests {
		t.Run(test.name, func(t *testing.T) {
			actual := BogoSort[int](test.input)
			sorted := reflect.DeepEqual(actual, test.expected)
			if !sorted {
				t.Errorf("test %s failed", test.name)
				t.Errorf("actual %v expected %v", actual, test.expected)
			}
		})
	}
}

func BenchmarkBogo(b *testing.B) {
	sortTests := []struct {
		input    []int
		expected []int
		name     string
	}{
		//Sorted slice
		{
			input:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			expected: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			name:     "Sorted Unsigned",
		},
		//Reversed slice
		{
			input:    []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
			expected: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			name:     "Reversed Unsigned",
		},
		//Sorted slice
		{
			input:    []int{-3, -2, -1, 0, 1, 2, 3},
			expected: []int{-3, -2, -1, 0, 1, 2, 3},
			name:     "Sorted Signed",
		},
		//Reversed slice
		{
			input:    []int{3, 2, 1, 0, -1, -2, -3},
			expected: []int{-3, -2, -1, 0, 1, 2, 3},
			name:     "Reversed Signed",
		},
		//Reversed slice, even length
		{
			input:    []int{3, 2, 1, -1, -2, -3},
			expected: []int{-3, -2, -1, 1, 2, 3},
			name:     "Reversed Signed #2",
		},
		//Random order with repetitions
		{
			input:    []int{4, -2, 3, 2, -1, 0, -3, -4, 1},
			expected: []int{-4, -3, -2, -1, 0, 1, 2, 3, 4},
			name:     "Random order Signed",
		},
		//Single-entry slice
		{
			input:    []int{1},
			expected: []int{1},
			name:     "Singleton",
		},
		// Empty slice
		{
			input:    []int{},
			expected: []int{},
			name:     "Empty Slice",
		},
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, test := range sortTests {
			BogoSort[int](test.input)
		}
	}
}
