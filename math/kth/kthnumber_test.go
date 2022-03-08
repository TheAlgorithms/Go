package kth

import (
	"testing"
)

func TestFindKthMax(t *testing.T) {
	sortTests := []struct {
		input    []int
		k        int
		expected int
		name     string
	}{
		{
			input:    []int{6, 7, 0, -1, 10, 70, 8, 22, 3, 9},
			k:        3,
			expected: 10,
			name:     "3th largest number",
		},
		{
			input:    []int{-5, 7, 4, -2, 6, 5, 8, 3, 2, -7, -1, 0, -3, 9, -6, 10, 10, 9, 1, -8, -9, -10},
			k:        1,
			expected: 10,
			name:     "largest number",
		},
		{
			input:    []int{1},
			k:        2,
			expected: -1,
			name:     "k is out of array bound",
		},
		{
			input:    []int{},
			k:        1,
			expected: -1,
			name:     "Empty Slice",
		},
	}
	for _, test := range sortTests {
		t.Run(test.name, func(t *testing.T) {
			actual := FindKthMax(test.input, test.k)
			if actual != test.expected {
				t.Errorf("test %s failed", test.name)
				t.Errorf("actual %v expected %v", actual, test.expected)
			}
		})
	}
}

func TestFindKthMin(t *testing.T) {
	sortTests := []struct {
		input    []int
		k        int
		expected int
		name     string
	}{
		{
			input:    []int{6, 7, 0, -1, 10, 70, 8, 22, 3, 9},
			k:        3,
			expected: 3,
			name:     "3th smallest number",
		},
		{
			input:    []int{-5, 7, 4, -2, 6, 5, 8, 3, 2, -7, -1, 0, -3, 9, -6, 10, -10, 9, 1, -8, -9, -10},
			k:        1,
			expected: -10,
			name:     "smallest number",
		},
		{
			input:    []int{1},
			k:        2,
			expected: -1,
			name:     "k is out of array bound",
		},
		{
			input:    []int{},
			k:        1,
			expected: -1,
			name:     "Empty Slice",
		},
	}
	for _, test := range sortTests {
		t.Run(test.name, func(t *testing.T) {
			actual := FindKthMin(test.input, test.k)
			if actual != test.expected {
				t.Errorf("test %s failed", test.name)
				t.Errorf("actual %v expected %v", actual, test.expected)
			}
		})
	}
}
