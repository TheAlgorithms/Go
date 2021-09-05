// LONGEST COMMON SUBSEQUENCE
// DP - 4
// https://www.geeksforgeeks.org/longest-common-subsequence-dp-4/

package dynamic_programming

// LongestCommonSubsequence function
func LongestCommonSubsequence(a string, b string, m int, n int) int {
	// m is the length of string a and n is the length of string b

	// here we are making a 2d slice of size (m+1)*(n+1)
	lcs := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		lcs[i] = make([]int, n+1)
	}

	// block that implements LCS
	for i := 0; i <= m; i++ {
		for j := 0; j <= n; j++ {
			if i == 0 || j == 0 {
				lcs[i][j] = 0
			} else if a[i-1] == b[j-1] {
				lcs[i][j] = lcs[i-1][j-1] + 1
			} else {
				lcs[i][j] = Max(lcs[i-1][j], lcs[i][j-1])
			}
		}
	}
	// returning the length of longest common subsequence
	return lcs[m][n]
}
