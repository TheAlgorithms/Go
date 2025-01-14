// eggdropping.go
// description: Solves the Egg Dropping Problem using dynamic programming
// reference: https://en.wikipedia.org/wiki/Egg_dropping_puzzle
// time complexity: O(k*n)
// space complexity: O(k*n)

package dynamic

import (
	"github.com/TheAlgorithms/Go/math/max"
	"github.com/TheAlgorithms/Go/math/min"
)

// EggDropping finds the minimum number of attempts needed to find the critical floor
// with `eggs` number of eggs and `floors` number of floors
func EggDropping(eggs, floors int) int {
	dp := make([][]int, eggs+1)
	for i := range dp {
		dp[i] = make([]int, floors+1)
	}

	for i := 1; i <= eggs; i++ {
		dp[i][0] = 0
		dp[i][1] = 1
	}

	for j := 0; j <= floors; j++ {
		dp[1][j] = j
	}

	for i := 2; i <= eggs; i++ {
		for j := 2; j <= floors; j++ {
			dp[i][j] = int(^uint(0) >> 1) // initialize with a large number
			for x := 1; x <= j; x++ {
				res := max.Int(dp[i-1][x-1], dp[i][j-x]) + 1
				dp[i][j] = min.Int(dp[i][j], res)
			}
		}
	}
	return dp[eggs][floors]
}
