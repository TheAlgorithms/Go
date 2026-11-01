package dynamic

// MaxNonAdjacentSum finds the maximum sum of non-adjacent elements
// in a given slice of integers.
//
// The algorithm uses dynamic programming with two variables to
// optimize space complexity.
//
// Time complexity: O(n)
// Space complexity: O(1)
//
// The empty subset is allowed, so the function returns 0 for
// an empty slice or a slice containing only negative numbers.
func MaxNonAdjacentSum(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	prevTwo := 0
	prevOne := 0

	for _, num := range nums {
		current := prevOne
		if prevTwo+num > current {
			current = prevTwo + num
		}

		prevTwo = prevOne
		prevOne = current
	}

	return prevOne
}
