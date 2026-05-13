// LONGEST COMMON SUBSTRING
// time complexity: O(m * n) where m and n are lengths of the two strings
// space complexity: O(m * n) for the DP table
// https://en.wikipedia.org/wiki/Longest_common_substring

package dynamic

// LongestCommonSubstring returns the length of the longest common substring
// between two strings using dynamic programming.
// Unlike Longest Common Subsequence, the substring must be contiguous.
func LongestCommonSubstring(str1, str2 string) int {
	m := len(str1)
	n := len(str2)

	if m == 0 || n == 0 {
		return 0
	}

	// dp[i][j] represents the length of the longest common substring
	// ending at str1[i-1] and str2[j-1]
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	maxLength := 0

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if str1[i-1] == str2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
				if dp[i][j] > maxLength {
					maxLength = dp[i][j]
				}
			}
			// else dp[i][j] = 0 (default value)
		}
	}

	return maxLength
}
