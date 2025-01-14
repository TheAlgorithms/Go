// longestarithmeticsubsequence.go
// description: Implementation of the Longest Arithmetic Subsequence problem
// reference: https://en.wikipedia.org/wiki/Longest_arithmetic_progression
// time complexity: O(n^2)
// space complexity: O(n^2)

package dynamic

// LongestArithmeticSubsequence returns the length of the longest arithmetic subsequence
func LongestArithmeticSubsequence(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return n
	}

	dp := make([]map[int]int, n)
	for i := range dp {
		dp[i] = make(map[int]int)
	}

	maxLength := 1

	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			diff := nums[i] - nums[j]
			dp[i][diff] = dp[j][diff] + 1
			if dp[i][diff]+1 > maxLength {
				maxLength = dp[i][diff] + 1
			}
		}
	}

	return maxLength
}
