// File: abbreviation.go
// Description: Abbreviation problem
// Details:
//   https://www.hackerrank.com/challenges/abbr/problem
// Problem description (from hackerrank):
//   You can perform the following operations on the string, a:
//   1. Capitalize zero or more of a's lowercase letters.
//   2. Delete all of the remaining lowercase letters in a.
//   Given 2 strings a and b, determine if it's possible to make a equal to be using above operations.
// Example:
//	 Given a = "ABcde" and b = "ABCD"
//   We can capitalize "c" and "d" in a to get "ABCde" then delete all the lowercase letters (which is only "e") in a to get "ABCD" which equals b.
// Author: [duongoku](https://github.com/duongoku)
// Time Complexity: O(n*m) where n is the length of a and m is the length of b
// Space Complexity: O(n*m) where n is the length of a and m is the length of b
// See abbreviation_test.go for test cases

package dynamic

// strings for getting uppercases and lowercases
import (
	"strings"
)

// Returns true if it is possible to make a equals b (if b is an abbreviation of a), returns false otherwise
func Abbreviation(a string, b string) bool {
	dp := make([][]bool, len(a)+1)
	for i := range dp {
		dp[i] = make([]bool, len(b)+1)
	}
	dp[0][0] = true

	for i := 0; i < len(a); i++ {
		for j := 0; j <= len(b); j++ {
			if dp[i][j] {
				if j < len(b) && strings.ToUpper(string(a[i])) == string(b[j]) {
					dp[i+1][j+1] = true
				}
				if string(a[i]) == strings.ToLower(string(a[i])) {
					dp[i+1][j] = true
				}
			}
		}
	}

	return dp[len(a)][len(b)]
}
