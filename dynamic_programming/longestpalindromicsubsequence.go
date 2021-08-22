// longest palindromic subsequence
// http://www.geeksforgeeks.org/dynamic-programming-set-12-longest-palindromic-subsequence/

package longest_palindromic_subsequence

import "github.com/TheAlgorithms/Go/dynamic_programming/knapsack"

// LpsRec function
func LpsRec(word string, i, j int) int {
	if i == j {
		return 1
	}
	if i > j {
		return 0
	}
	if word[i] == word[j] {
		return 2 + LpsRec(word, i+1, j-1)
	}
	return knapsack.Max(LpsRec(word, i, j-1), LpsRec(word, i+1, j))
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
				dp[i][j] = knapsack.Max(dp[i+1][j], dp[i][j-1])
			}
		}
	}

	return dp[1][N-1]
}

/*
func main() {
	// word := "aaabbbbababbabbabbabf"
	word := "aaaabbbba"
	fmt.Printf("%d\n", lpsRec(word, 0, len(word)-1))
	fmt.Printf("%d\n", lpsDp(word))
}
*/
