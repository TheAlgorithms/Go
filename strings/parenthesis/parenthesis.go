package parenthesis

// BalancedParenthesis algorithm checks if every opened parenthesis
// is closed correctly

// this implementation uses a stack

func BalancedParenthesis(text string) bool {

	// set up stack and opener-closer map
	var stack []rune
	openerCloser := map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
	}

	for _, c := range text {
		// ignore characters other than brackets
		if c == '(' || c == ')' || c == '[' || c == ']' || c == '{' || c == '}' {
			// if the current character is an opening bracket,
			// push its closer into the stack and continue
			if closer, ok := openerCloser[c]; ok {
				stack = append(stack, closer)
				continue
			}

			// else we are dealing with a closer
			// the stack should not be empty and the top
			// of the stack should be the current character
			l := len(stack) - 1
			if l < 0 || c != stack[l] {
				return false
			}

			// remove the top element of stack
			stack = stack[:l]
		}
	}

	// if the stack is empty, return true, otherwise false
	return len(stack) == 0
}
