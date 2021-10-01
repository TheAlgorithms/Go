// generateparentheses_test.go
// description: Generate Parentheses
// author(s) [red_byte](https://github.com/i-redbyte)
// see generateparentheses.go

package generateparentheses

import "testing"

func TestGenerateParenthesis(t *testing.T) {
	t.Run("GenerateParenthesis", func(t *testing.T) {
		result := GenerateParenthesis(3)
		t.Log(result)
		exp := []string{"((()))", "(()())", "(())()", "()(())", "()()()"}
		for i, v := range result {
			if v != exp[i] {
				t.Errorf("Wrong result! Expected:%s, returned:%s ", result[i], exp[i])
			}
		}
	})
}
