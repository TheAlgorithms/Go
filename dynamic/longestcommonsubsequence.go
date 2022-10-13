// LONGEST COMMON SUBSEQUENCE
// DP - 4
// https://www.geeksforgeeks.org/longest-common-subsequence-dp-4/

package dynamic

import (
	"unicode/utf8"
)

// LongestCommonSubsequence function
func LongestCommonSubsequence(a string, b string) int {
	var a_len = utf8.RuneCountInString(a)
	var b_len = utf8.RuneCountInString(b)

	// here we are making a 2d slice of size (a_len+1)*(b_len+1)
	lcs := make([][]int, a_len+1)
	for i := 0; i <= a_len; i++ {
		lcs[i] = make([]int, b_len+1)
	}

	// block that implements LCS
	for i := 0; i <= a_len; i++ {
		for j := 0; j <= b_len; j++ {
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
	return lcs[a_len][b_len]
}

// func main(){
// 	// declaring two strings and asking for input

// 	var a,b string
// 	fmt.Scan(&a, &b)
// 	// calling the LCS function
// 	fmt.Println("The length of longest common subsequence is:", longestCommonSubsequence(a,b))
// }
