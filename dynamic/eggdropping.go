package dynamic

import (
	"github.com/TheAlgorithms/Go/math/max"
	"github.com/TheAlgorithms/Go/math/min"
)

// EggDropping finds the minimum number of attempts needed to find the critical floor
// with `eggs` number of eggs and `floors` number of floors
func EggDropping(eggs, floors int) int {
	// Edge case: If there are no floors, no attempts needed
	if floors == 0 {
		return 0
	}
	// Edge case: If there is one floor, one attempt needed
	if floors == 1 {
		return 1
	}
	// Edge case: If there is one egg, need to test all floors one by one
	if eggs == 1 {
		return floors
	}

	// Initialize DP table
	dp := make([][]int, eggs+1)
	for i := range dp {
		dp[i] = make([]int, floors+1)
	}

	// Fill the DP table for 1 egg
	for j := 1; j <= floors; j++ {
		dp[1][j] = j
	}

	// Fill the DP table for more than 1 egg
	for i := 2; i <= eggs; i++ {
		for j := 2; j <= floors; j++ {
			dp[i][j] = int(^uint(0) >> 1) // initialize with a large number
			for x := 1; x <= j; x++ {
				// Recurrence relation to fill the DP table
				res := max.Int(dp[i-1][x-1], dp[i][j-x]) + 1
				dp[i][j] = min.Int(dp[i][j], res)
			}
		}
	}
	return dp[eggs][floors]
}
