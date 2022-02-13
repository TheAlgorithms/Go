// subset.go
// description: subset of a given set
// details:
//In mathematics, set A is a subset of a set B if all elements of A are also elements of B; B is then a superset of A. It is possible for A and B to be equal; if they are unequal, then A is a proper subset of B. The relationship of one set being a subset of another is called inclusion (or sometimes containment). A is a subset of B may also be expressed as B includes (or contains) A or A is included (or contained) in B.
//  - https://en.wikipedia.org/wiki/Subset
// author: [sixwaaaay](https://github.com/sixwaaaay)
// see subset_test.go

package backtrack

func Subset(nums []int) (set [][]int) {
	n := len(nums)
	// init index array for record the index of each element selected
	index := make([]bool, n)
	for i := 0; i < n; i++ {
		index[i] = false
	}
	var backtrack func(int)
	backtrack = func(depth int) {
		// use index to record the selected elements
		if depth == n {
			temp := []int{}
			for i := range index {
				if index[i] {
					temp = append(temp, nums[i])
				}
			}
			set = append(set, temp)
			return
		}
		index[depth] = true
		backtrack(depth + 1)
		index[depth] = false
		backtrack(depth + 1)
	}
	backtrack(0)
	return
}
