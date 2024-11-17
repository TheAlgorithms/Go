// longest palindromic subsequence
// time complexity: O(n^2)
// space complexity: O(n^2)
// http://www.geeksforgeeks.org/dynamic-programming-set-12-longest-palindromic-subsequence/

package dynamic

func lpsRec(word string, i, j int) int {
	if i == j {
		return 1
	}
	if i > j {
		return 0
	}
	if word[i] == word[j] {
		return 2 + lpsRec(word, i+1, j-1)
	}
	return Max(lpsRec(word, i, j-1), lpsRec(word, i+1, j))
}

// LpsRec function
func LpsRec(word string) int {
	return lpsRec(word, 0, len(word)-1)
}

// LpsDp function
func LpsDp(word string) int {
	N := len(word)
	dp := make([][]int, N)

	for i := 0; i < N; i++ {
		dp[i] = make([]int, N)
		dp[i][i] = 1
	}

	for l := 2; l <= N; l++ {
		// for length l
		for i := 0; i < N-l+1; i++ {
			j := i + l - 1
			if word[i] == word[j] {
				if l == 2 {
					dp[i][j] = 2
				} else {
					dp[i][j] = 2 + dp[i+1][j-1]
				}
			} else {
				dp[i][j] = Max(dp[i+1][j], dp[i][j-1])
			}
		}
	}

	return dp[0][N-1]
}
