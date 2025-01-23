package dynamic

import "github.com/TheAlgorithms/Go/math/max"

// MaxCoins returns the maximum coins we can collect by bursting the balloons
func MaxCoins(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	nums = append([]int{1}, nums...)
	nums = append(nums, 1)

	dp := make([][]int, n+2)
	for i := range dp {
		dp[i] = make([]int, n+2)
	}

	for length := 1; length <= n; length++ {
		for left := 1; left+length-1 <= n; left++ {
			right := left + length - 1
			for k := left; k <= right; k++ {
				coins := nums[left-1] * nums[k] * nums[right+1]
				dp[left][right] = max.Int(dp[left][right], dp[left][k-1]+dp[k+1][right]+coins)
			}
		}
	}

	return dp[1][n]
}
