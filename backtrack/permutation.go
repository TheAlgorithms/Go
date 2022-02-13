// permutation.go
// description: generate permutations of a given set
// details:
// In mathematics, a permutation of a set is, loosely speaking, an arrangement of its members into a sequence or linear order, or if the set is already ordered, a rearrangement of its elements. The word "permutation" also refers to the act or process of changing the linear order of an ordered set.
// author: [sixwaaaay](https://github.com/sixwaaaay)
// see permutation_test.go

package backtrack

func Permutation(nums []int) (set [][]int) {
	n := len(nums)
	var backtrack func(int)
	backtrack = func(depth int) {
		if depth == n {
			set = append(set, append([]int{}, nums...))
			return
		}
		for i := depth; i < len(nums); i++ {
			nums[depth], nums[i] = nums[i], nums[depth]
			backtrack(depth + 1)
			nums[depth], nums[i] = nums[i], nums[depth]
		}
	}
	backtrack(0)
	return
}
