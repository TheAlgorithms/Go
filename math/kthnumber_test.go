package math

import (
	"github.com/TheAlgorithms/Go/search"
	"testing"
)

func TestFindKthMax(t *testing.T) {
	sortTests := []struct {
		input    []int
		k        int
		expected int
		err      error
		name     string
	}{
		{
			input:    []int{6, 7, 0, -1, 10, 70, 8, 22, 3, 9},
			k:        3,
			expected: 10,
			name:     "3th largest number",
		},
		{
			input:    []int{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1},
			k:        3,
			expected: -1,
			name:     "3th largest number",
		},
		{
			input:    []int{-1, -1, -1, -1, -1, -1},
			k:        7,
			expected: -1,
			err:      search.ErrNotFound,
			name:     "This should be an error",
		},
		{
			input:    []int{},
			k:        1,
			expected: -1,
			err:      search.ErrNotFound,
			name:     "This should be an error",
		},
	}
	for _, test := range sortTests {
		t.Run(test.name, func(t *testing.T) {
			actual, err := FindKthMax(test.input, test.k)
			if err != test.err {
				t.Errorf("name:%v FindKthMax() = %v, want err: %v", test.name, err, test.err)
			}
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
		err      error
		name     string
	}{
		{
			input:    []int{6, 7, 0, -1, 10, 70, 8, 22, 3, 9},
			k:        3,
			expected: 3,
			name:     "3th smallest number",
		},
		{
			input:    []int{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1},
			k:        3,
			expected: -1,
			name:     "3th smallest number",
		},
		{
			input:    []int{-1, -1, -1, -1, -1, -1},
			k:        7,
			expected: -1,
			err:      search.ErrNotFound,
			name:     "This should be an error",
		},
		{
			input:    []int{},
			k:        1,
			expected: -1,
			err:      search.ErrNotFound,
			name:     "This should be an error",
		},
	}
	for _, test := range sortTests {
		t.Run(test.name, func(t *testing.T) {
			actual, err := FindKthMin(test.input, test.k)
			if err != test.err {
				t.Errorf("name:%v FindKthMin() = %v, want err: %v", test.name, err, test.err)
			}
			if actual != test.expected {
				t.Errorf("test %s failed", test.name)
				t.Errorf("actual %v expected %v", actual, test.expected)
			}
		})
	}
}
