// mininsertionspalindrome.go
// description: Minimum insertions to form a palindrome DP solution
// details:
// Given a string, the task is to find the minimum number of characters to be inserted to convert it to palindrome.
// [geeksforgeeks](https://www.geeksforgeeks.org/minimum-insertions-to-form-a-palindrome-dp-28/)
// see mininsertionspalindrome_test.go for test implementation
 
package dynamic

//Utility Function
func Minimum(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func FindMinInsertionsPalindrome(str string) int {
	n := len(str)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	for gap := 1; gap < n; gap++ {
		left := 0
		for right := gap; right < n; right++ {
			if str[left] == str[right] {
				dp[left][right] = dp[left + 1][right - 1]
			} else {
				dp[left][right] = Minimum(dp[left][right - 1],dp[left + 1][right]) + 1
			}
			left += 1
		}
	}
	return dp[0][n-1]
}