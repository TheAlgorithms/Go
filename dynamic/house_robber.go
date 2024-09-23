// See https://leetcode.com/problems/house-robber/
// author: Danish Eqbal (https://github.com/dcode-github)
package dynamic

func HouseRobber(nums []int) int {

	var max = func(i1, i2 int) int {
		if i1 > i2 {
			return i1
		}
		return i2
	}

	if len(nums) == 1 {
		return nums[0]
	}

	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		dp[i] = max(nums[i]+dp[i-2], dp[i-1])
	}

	return dp[len(nums)-1]
}
