// dicethrow.go
// description: Solves the Dice Throw Problem using dynamic programming
// reference: https://www.geeksforgeeks.org/dice-throw-problem/
// time complexity: O(m * n)
// space complexity: O(m * n)

package dynamic

// DiceThrow returns the number of ways to get sum `sum` using `m` dice with `n` faces
func DiceThrow(m, n, sum int) int {
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, sum+1)
	}

	for i := 1; i <= n; i++ {
		if i <= sum {
			dp[1][i] = 1
		}
	}

	for i := 2; i <= m; i++ {
		for j := 1; j <= sum; j++ {
			for k := 1; k <= n; k++ {
				if j-k >= 0 {
					dp[i][j] += dp[i-1][j-k]
				}
			}
		}
	}

	return dp[m][sum]
}
