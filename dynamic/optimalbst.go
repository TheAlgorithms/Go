// optimalbst.go
// description: Solves the Optimal Binary Search Tree problem using dynamic programming
// reference: https://en.wikipedia.org/wiki/Optimal_binary_search_tree
// time complexity: O(n^3)
// space complexity: O(n^2)

package dynamic

import "github.com/TheAlgorithms/Go/math/min"

// OptimalBST returns the minimum cost of constructing a Binary Search Tree
func OptimalBST(keys []int, freq []int, n int) int {
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		dp[i][i] = freq[i]
	}

	for length := 2; length <= n; length++ {
		for i := 0; i < n-length+1; i++ {
			j := i + length - 1
			dp[i][j] = int(^uint(0) >> 1)
			sum := sum(freq, i, j)
			for k := i; k <= j; k++ {
				cost := sum + dp[i][k-1] + dp[k+1][j]
				dp[i][j] = min.Int(dp[i][j], cost)
			}
		}
	}
	return dp[0][n-1]
}

func sum(freq []int, i, j int) int {
	total := 0
	for k := i; k <= j; k++ {
		total += freq[k]
	}
	return total
}
