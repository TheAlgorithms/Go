package dynamic

// ClimbingStairs calculates the number of distinct ways to climb
// n stairs when each step can be either 1 or 2 stairs.
//
// It uses dynamic programming with two variables to optimize
// space complexity.
//
// Time complexity: O(n)
// Space complexity: O(1)
//
// Returns 1 for n = 0 and 0 for negative values of n.
//
// Note: The result may overflow int for sufficiently large n.
func ClimbingStairs(n int) int {
	if n < 0 {
		return 0
	}

	if n <= 1 {
		return 1
	}

	prevTwo := 1
	prevOne := 1

	for i := 2; i <= n; i++ {
		current := prevOne + prevTwo
		prevTwo = prevOne
		prevOne = current
	}

	return prevOne
}
