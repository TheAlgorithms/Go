// generateparenthesesgo
// description: Generate Parentheses
// details:
// Given n pairs of parentheses, write a function to generate all combinations of well-formed parentheses.
// author(s) [red_byte](https://github.com/i-redbyte)
// see generateparentheses_test.go

package generateparentheses

import "strings"

func GenerateParenthesis(n int) []string {
	result := make([]string, 0)
	maxLen := 2 * n
	var recursiveComputation func(s []string, left int, right int)
	recursiveComputation = func(s []string, left int, right int) {
		if len(s) == maxLen {
			result = append(result, strings.Join(s, ""))
			return
		}
		if left < n {
			s = append(s, "(")
			recursiveComputation(s, left+1, right)
			s = s[:len(s)-1]
		}
		if right < left {
			s = append(s, ")")
			recursiveComputation(s, left, right+1)
			_ = s[:len(s)-1]
		}
	}
	recursiveComputation(make([]string, 0), 0, 0)
	return result
}
