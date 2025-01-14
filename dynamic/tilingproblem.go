// tilingproblem.go
// description: Solves the Tiling Problem using dynamic programming
// reference: https://en.wikipedia.org/wiki/Tiling_problem
// time complexity: O(n)
// space complexity: O(n)

package dynamic

// TilingProblem returns the number of ways to tile a 2xN grid using 2x1 dominoes
func TilingProblem(n int) int {
	if n <= 1 {
		return 1
	}
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1

	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}
