// Package nested provides functions for testing
// strings proper brackets nesting.
package nested

// IsBalanced returns true if provided input string is properly nested.
//
// Input is a sequence of brackets: '(', ')', '[', ']', '{', '}'.
//
// A sequence of brackets `s` is considered properly nested
// if any of the following conditions are true:
//   - `s` is empty;
//   - `s` has the form (U) or [U] or {U} where U is a properly nested string;
//   - `s` has the form VW where V and W are properly nested strings.
//
// For example, the string "()()[()]" is properly nested but "[(()]" is not.
//
// **Note** Providing characters other then brackets would return false,
// despite brackets sequence in the string. Make sure to filter
// input before usage.
// time complexity: O(n)
// space complexity: O(n)

func IsBalanced(input string) bool {
	if len(input) == 0 {
		return true
	}

	if len(input)%2 != 0 {
		return false
	}

	// Brackets such as '{', '[', '(' are valid UTF-8 characters,
	// which means that only one byte is required to code them,
	// so can be stored as bytes.
	var stack []byte

	for i := 0; i < len(input); i++ {
		if input[i] == '(' || input[i] == '{' || input[i] == '[' {
			stack = append(stack, input[i])
		} else {
			if len(stack) > 0 {
				pair := string(stack[len(stack)-1]) + string(input[i])
				stack = stack[:len(stack)-1]

				if pair != "[]" && pair != "{}" && pair != "()" {
					// This means that two types of brackets has
					// been mixed together, for example "([)]",
					// which makes seuqence invalid by definition.
					return false
				}
			} else {
				// This means that closing bracket is encountered
				// before opening one, which makes all sequence
				// invalid by definition.
				return false
			}
		}
	}

	// If sequence is properly nested, all elements in stack
	// has been paired with closing elements. If even one
	// element has not been paired with a closing bracket,
	// means that sequence is invalid by definition.
	return len(stack) == 0
}
