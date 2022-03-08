// permutation_test.go
// description: test generating permutation of a given array
// author: [sixwaaaay](https://github.com/sixwaaaay)
// see permutation.go

package backtrack

import (
	"reflect"
	"testing"
)

var testCasesForPermutation = []struct {
	name         string  // test case name
	nums         []int   // nums array
	permutations [][]int // expected subsets
}{
	{
		name: "array without duplicate elements",
		nums: []int{
			2, 3, 4,
		},
		permutations: [][]int{
			{2, 3, 4},
			{2, 4, 3},
			{3, 2, 4},
			{3, 4, 2},
			{4, 2, 3},
			{4, 3, 2},
		},
	},
	{name: "array with two duplicate numbers",
		nums: []int{
			1, 3, 3,
		},
		permutations: [][]int{
			{1, 3, 3},
			{3, 1, 3},
			{3, 3, 1},
		},
	},
	{
		name: "array with three repeating numbers",
		nums: []int{0, 1, 0, 9, 0},
		permutations: [][]int{
			{0, 0, 0, 1, 9}, {0, 0, 0, 9, 1}, {0, 0, 1, 0, 9}, {0, 0, 1, 9, 0},
			{0, 0, 9, 0, 1}, {0, 0, 9, 1, 0}, {0, 1, 0, 0, 9}, {0, 1, 0, 9, 0},
			{0, 1, 9, 0, 0}, {0, 9, 0, 0, 1}, {0, 9, 0, 1, 0}, {0, 9, 1, 0, 0},
			{1, 0, 0, 0, 9}, {1, 0, 0, 9, 0}, {1, 0, 9, 0, 0}, {1, 9, 0, 0, 0},
			{9, 0, 0, 0, 1}, {9, 0, 0, 1, 0}, {9, 0, 1, 0, 0}, {9, 1, 0, 0, 0}},
	},
	{
		name: "empty array",
		nums: nil,
		permutations: [][]int{
			{},
		},
	},
	{
		name: "array with element greater than length",
		nums: []int{9, 10},
		permutations: [][]int{
			{9, 10},
			{10, 9},
		},
	},
}

// test permutation function
func TestPermutation(t *testing.T) {
	for _, testCase := range testCasesForPermutation {
		t.Run(testCase.name, func(t *testing.T) {
			output := Permutation(testCase.nums)
			if len(output) != len(testCase.permutations) {
				t.Errorf("expected length %d, got %d", len(testCase.permutations), len(output))
			}
			if reflect.DeepEqual(output, testCase.permutations) == false {
				t.Errorf("expected %v, got %v", testCase.permutations, output)
			}
		})
	}
}
