// subset_test.go
// description: Test for generating all subsets of a set
// author(s) [sixwaaaay](https://github.com//sixwaaaay)
// see subset.go

package backtrack

import (
	"reflect"
	"testing"
)

var testCasesForSubset = []struct {
	name    string  // test case name
	nums    []int   // nums array
	subsets [][]int // expected subsets
}{
	{
		name: "array without duplicate elements",
		nums: []int{
			2, 3, 4,
		},
		subsets: [][]int{
			nil,
			{4},
			{3},
			{3, 4},
			{2},
			{2, 4},
			{2, 3},
			{2, 3, 4},
		},
	},
	{name: "array with two duplicate numbers",
		nums: []int{
			1, 3, 3,
		},
		subsets: [][]int{
			nil,
			{3},
			{3, 3},
			{1},
			{1, 3},
			{1, 3, 3},
		},
	},
	{
		name: "array with three repeating numbers",
		nums: []int{0, 1, 0, 9, 0},
		subsets: [][]int{
			nil, {9}, {1}, {1, 9}, {0}, {0, 9}, {0, 1}, {0, 1, 9},
			{0, 0}, {0, 0, 9}, {0, 0, 1}, {0, 0, 1, 9}, {0, 0, 0},
			{0, 0, 0, 9}, {0, 0, 0, 1}, {0, 0, 0, 1, 9},
		},
	},
	{
		name: "empty array",
		nums: nil,
		subsets: [][]int{
			nil,
		},
	},
	{
		name: "array with element greater than length",
		nums: []int{7, 12},
		subsets: [][]int{
			nil,
			{12},
			{7},
			{7, 12},
		},
	},
}

func TestSubset(t *testing.T) {
	for _, testCase := range testCasesForSubset {
		t.Run(testCase.name, func(t *testing.T) {
			subsets := Subset(testCase.nums)
			if len(subsets) != len(testCase.subsets) {
				t.Errorf("expected length %d, got %d", len(testCase.subsets), len(subsets))
			}
			if reflect.DeepEqual(subsets, testCase.subsets) == false {
				t.Errorf("expected %v, got %v", testCase.subsets, subsets)
			}
		})
	}
}
