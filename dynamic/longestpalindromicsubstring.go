// longestpalindromicsubstring.go
// description: Implementation of finding the longest palindromic substring
// reference: https://en.wikipedia.org/wiki/Longest_palindromic_substring
// time complexity: O(n^2)
// space complexity: O(n^2)

package dynamic

// LongestPalindromicSubstring returns the longest palindromic substring in the input string
func LongestPalindromicSubstring(s string) string {
	n := len(s)
	if n == 0 {
		return ""
	}

	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
	}

	start := 0
	maxLength := 1
	for i := 0; i < n; i++ {
		dp[i][i] = true
	}

	for length := 2; length <= n; length++ {
		for i := 0; i < n-length+1; i++ {
			j := i + length - 1
			if length == 2 {
				dp[i][j] = (s[i] == s[j])
			} else {
				dp[i][j] = (s[i] == s[j]) && dp[i+1][j-1]
			}

			if dp[i][j] && length > maxLength {
				maxLength = length
				start = i
			}
		}
	}
	return s[start : start+maxLength]
}
