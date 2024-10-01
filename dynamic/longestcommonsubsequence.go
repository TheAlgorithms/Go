// LONGEST COMMON SUBSEQUENCE
// DP - 4
// https://www.geeksforgeeks.org/longest-common-subsequence-dp-4/
// https://leetcode.com/problems/longest-common-subsequence/
// time complexity: O(m*n) where m and n are lengths of the strings
// space complexity: O(m*n) where m and n are lengths of the strings

package dynamic

func strToRuneSlice(s string) (r []rune, size int) {
	r = []rune(s)
	return r, len(r)
}

// LongestCommonSubsequence function
func LongestCommonSubsequence(a string, b string) int {
	aRunes, aLen := strToRuneSlice(a)
	bRunes, bLen := strToRuneSlice(b)

	// here we are making a 2d slice of size (aLen+1)*(bLen+1)
	lcs := make([][]int, aLen+1)
	for i := 0; i <= aLen; i++ {
		lcs[i] = make([]int, bLen+1)
	}

	// block that implements LCS
	for i := 0; i <= aLen; i++ {
		for j := 0; j <= bLen; j++ {
			if i == 0 || j == 0 {
				lcs[i][j] = 0
			} else if aRunes[i-1] == bRunes[j-1] {
				lcs[i][j] = lcs[i-1][j-1] + 1
			} else {
				lcs[i][j] = Max(lcs[i-1][j], lcs[i][j-1])
			}
		}
	}
	// returning the length of longest common subsequence
	return lcs[aLen][bLen]
}
