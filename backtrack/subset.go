// subset.go
// description: subset of a given set
// details:
// In mathematics,
// set A is a subset of a set B if all elements of A are also elements of B;
// B is then a superset of A. It is possible for A and B to be equal;
// if they are unequal, then A is a proper subset of B.
// The relationship of one set being a subset of another
// is called inclusion (or sometimes containment).
// A is a subset of B may also be expressed as B includes (or contains) A
// or A is included (or contained) in B.
// - https://en.wikipedia.org/wiki/Subset
// author: [sixwaaaay](https://github.com/sixwaaaay)
// see subset_test.go

package backtrack

import "sort"

// Subset is a function to generate all subsets of a given array
func Subset(nums []int) [][]int {
	sort.Ints(nums)
	var subsets [][]int
	index := make([]bool, len(nums))
	backtrackSubset(nums, index, 0, &subsets)
	return subsets
}

func backtrackSubset(nums []int, index []bool, depth int, subsets *[][]int) {

	if depth == len(nums) { // use index to record the selected elements
		var subset []int
		for i := 0; i < len(nums); i++ {
			if index[i] {
				subset = append(subset, nums[i])
			}
		} // add index mapped array to subsets
		*subsets = append(*subsets, subset)
		return
	}
	// not select element at current depth (index[depth] equals to false)
	backtrackSubset(nums, index, depth+1, subsets)
	if depth > 0 && !index[depth-1] && nums[depth-1] == nums[depth] {
		return
	}
	index[depth] = true // select element at current depth
	backtrackSubset(nums, index, depth+1, subsets)
	index[depth] = false
}
