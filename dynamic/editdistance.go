// EDIT DISTANCE PROBLEM
// https://www.geeksforgeeks.org/edit-distance-dp-5/
// https://leetcode.com/problems/edit-distance/

package dynamic

// Utility function which does what you expect.
func minOfTwo(a int, b int) int {

	if a < b {
		return a
	}
	return b
}

// Utility function to compute minimum of three numbers using minOfTwo.
func minOfThree(a int, b int, c int) int {

	return minOfTwo(a, minOfTwo(b, c))
}

// EditDistanceRecursive is a naive implementation with exponential time complexity.
func EditDistanceRecursive(first string, second string, pointerFirst int, pointerSecond int) int {

	if pointerFirst == 0 {
		return pointerSecond
	}

	if pointerSecond == 0 {
		return pointerFirst
	}

	// Characters match, so we recur for the remaining portions
	if first[pointerFirst-1] == second[pointerSecond-1] {
		return EditDistanceRecursive(first, second, pointerFirst-1, pointerSecond-1)
	}

	// We have three choices, all with cost of 1 unit
	return 1 + minOfThree(EditDistanceRecursive(first, second, pointerFirst, pointerSecond-1), // Insert
		EditDistanceRecursive(first, second, pointerFirst-1, pointerSecond),   // Delete
		EditDistanceRecursive(first, second, pointerFirst-1, pointerSecond-1)) // Replace
}

// EditDistanceDP is an optimised implementation which builds on the ideas of the recursive implementation.
// We use dynamic programming to compute the DP table where dp[i][j] denotes the edit distance value
// of first[0..i-1] and second[0..j-1]. Time complexity is O(m * n) where m and n are lengths of the strings,
// first and second respectively.
func EditDistanceDP(first string, second string) int {

	m := len(first)
	n := len(second)

	// Create the DP table
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}

	for i := 0; i <= m; i++ {
		for j := 0; j <= n; j++ {

			if i == 0 {
				dp[i][j] = j
				continue
			}

			if j == 0 {
				dp[i][j] = i
				continue
			}

			if first[i-1] == second[j-1] {
				dp[i][j] = dp[i-1][j-1]
				continue
			}

			dp[i][j] = 1 + minOfThree(dp[i][j-1], dp[i-1][j], dp[i-1][j-1])
		}
	}

	return dp[m][n]
}
