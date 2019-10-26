package main

import (
	"fmt"

	"github.com/george-e-shaw-iv/Go/data-structures/stack"
)

func IsPalindrome(input string) bool {
	var s stack.Stack
	var output string

	for _, ch := range input {
		s.Push(string(ch))
	}

	for !s.IsEmpty() {
		output += s.Top().(string)
		s.Pop()
	}

	return input == output
}

func main() {
	candidates := []string{
		"racecar", "salad", "tacocat", "bananas",
	}

	fmt.Println("Checking candidates for Palindromes:")
	for _, candidate := range candidates {
		fmt.Printf("%s: %t\n", candidate, IsPalindrome(candidate))
	}
}
