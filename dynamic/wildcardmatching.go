// wildcardmatching.go
// description: Solves the Wildcard Matching problem using dynamic programming
// reference: https://en.wikipedia.org/wiki/Wildcard_matching
// time complexity: O(m*n)
// space complexity: O(m*n)

package dynamic

// IsMatch checks if the string `s` matches the wildcard pattern `p`
func IsMatch(s, p string) bool {
	dp := make([][]bool, len(s)+1)
	for i := range dp {
		dp[i] = make([]bool, len(p)+1)
	}

	dp[0][0] = true
	for j := 1; j <= len(p); j++ {
		if p[j-1] == '*' {
			dp[0][j] = dp[0][j-1]
		}
	}

	for i := 1; i <= len(s); i++ {
		for j := 1; j <= len(p); j++ {
			if p[j-1] == s[i-1] || p[j-1] == '?' {
				dp[i][j] = dp[i-1][j-1]
			} else if p[j-1] == '*' {
				dp[i][j] = dp[i-1][j] || dp[i][j-1]
			}
		}
	}
	return dp[len(s)][len(p)]
}
