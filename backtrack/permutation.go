// permutation.go
// description: generate subsets use backtrack algorithm
// details:
// In mathematics, a permutation of a set is, loosely speaking,
// an arrangement of its members into a sequence or linear order,
// or if the set is already ordered, a rearrangement of its elements.
// The word "permutation" also refers to the act or process of changing
// the linear order of an ordered set.
// author: [sixwaaaay](https://github.com/sixwaaaay)
// see permutation_test.go

package backtrack

import "sort"

// Permutation is a function to generate all permutations of a given array
func Permutation(nums []int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	perm := make([]int, n)
	vis := make([]bool, n)
	var permutations [][]int
	var backtrackPerm func(int)
	backtrackPerm = func(depth int) {
		if depth == n {
			permutations = append(permutations, append([]int{}, perm...))
			return
		}
		for i := range nums {
			if vis[i] || (i > 0 && !vis[i-1] && nums[i] == nums[i-1]) {
				continue
			} // get first element that never used/visited
			perm[depth] = nums[i]
			vis[i] = true
			backtrackPerm(depth + 1)
			vis[i] = false
		}
	}
	backtrackPerm(0)
	return permutations
}
