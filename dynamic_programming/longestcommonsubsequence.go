// LONGEST COMMON SUBSEQUENCE
// DP - 4
// https://www.geeksforgeeks.org/longest-common-subsequence-dp-4/

package longest_common_subsequence

import "github.com/TheAlgorithms/Go/dynamic_programming/knapsack"

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
				lcs[i][j] = knapsack.Max(lcs[i-1][j], lcs[i][j-1])
			}
		}
	}
	// returning the length of longest common subsequence
	return lcs[m][n]
}

// func main(){
// 	// declaring two strings and asking for input

// 	var a,b string
// 	fmt.Scan(&a, &b)
// 	// calling the LCS function
// 	fmt.Println("The length of longest common subsequence is:", longestCommonSubsequence(a,b, len(a), len(b)))
// }
