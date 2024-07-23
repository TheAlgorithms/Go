package sort

import (
	"reflect"
	"testing"
)

func TestCircle(t *testing.T) {
	testCases := []struct {
		input    []int
		expected []int
	}{
		{input: []int{3, 1, 2}, expected: []int{1, 2, 3}},
		{input: []int{5, 3, 4, 1, 2}, expected: []int{1, 2, 3, 4, 5}},
		{input: []int{1, 2, 3, 4, 5}, expected: []int{1, 2, 3, 4, 5}},
		{input: []int{}, expected: []int{}},
	}

	for _, tc := range testCases {
		result := Circle(tc.input)
		if !reflect.DeepEqual(result, tc.expected) {
			t.Errorf("expected %v, got %v", tc.expected, result)
		}
	}
}
