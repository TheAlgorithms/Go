// mininsertionspalindrome.go
// description: Minimum insertions to form a palindrome DP solution
// details:
// Given a string, the task is to find the minimum number of characters to be inserted to convert it to palindrome.
// [geeksforgeeks](https://www.geeksforgeeks.org/minimum-insertions-to-form-a-palindrome-dp-28/)
// see mininsertionspalindrome_test.go for test implementation

package dynamic

// FindMinInsertionsPalindrome returns minimum number of insertions required to convert given string to a palindrome
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
				dp[left][right] = dp[left+1][right-1]
			} else {
				// Min :- Minimum function imported from binomialcoefficient.go in package dynamic
				dp[left][right] = Min(dp[left][right-1], dp[left+1][right]) + 1
			}
			left++
		}
	}
	return dp[0][n-1]
}
