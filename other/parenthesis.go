package main

import (
	"fmt"

	"github.com/george-e-shaw-iv/Go/data-structures/stack"
)

func ValidParenthesis(input string) bool {
	var s stack.Stack

	for _, ch := range input {
		if string(ch) == "(" {
			s.Push(string(ch))
		} else if string(ch) == ")" {
			if s.IsEmpty() {
				return false
			}

			s.Pop()
		} else {
			// Skip all non-parenthesis characters
			continue
		}
	}

	// If stack is empty, the parenthesis in the given string are valid.
	return s.IsEmpty()
}

func main() {
	candidates := []string{
		"(()())", "((foo)(bar))", "())(((())", "(g))sd(f((())",
	}

	fmt.Println("Validating the candidates for parenthesises:")
	for _, candidate := range candidates {
		fmt.Printf("%s: %t\n", candidate, ValidParenthesis(candidate))
	}
}
