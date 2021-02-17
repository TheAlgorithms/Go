// The nested brackets problem is a problem that determines if a sequence of
// brackets are properly nested.  A sequence of brackets s is considered properly nested
// if any of the following conditions are true:
// 	- s is empty
// 	- s has the form (U) or [U] or {U} where U is a properly nested string
// 	- s has the form VW where V and W are properly nested strings
// For example, the string "()()[()]" is properly nested but "[(()]" is not.
// The function called isBalanced takes as input a string which is a sequence of brackets and
// returns true if input is nested and false otherwise.
//note that only an even number of brackets can be properly nested

package nestedbrackets

func isBalanced(input string) string {
	if len(input)%2 != 0 {
		return input + "is not balanced."
	}
	if len(input) > 0 {
		var stack []byte
		for i := 0; i < len(input); i++ {
			if input[i] == '(' || input[i] == '{' || input[i] == '[' {
				stack = append(stack, input[i])
			} else {
				if len(stack) > 0 {
					pair := string(stack[len(stack)-1]) + string(input[i])
					stack = stack[:len(stack)-1]

					if pair != "[]" && pair != "{}" && pair != "()" {
						return input + " is not balanced."
					}
				} else {
					return input + " is not balanced."
				}
			}
		}
		if len(stack) == 0 {
			return input + " is balanced."
		}
	}
	return "Please enter a sequence of brackets."
}

// func main() {
// 	reader := bufio.NewReader(os.Stdin)
// 	fmt.Print("Enter sequence of brackets: ")
// 	text, _ := reader.ReadString('\n')

// 	text = strings.TrimSpace(text)
// 	fmt.Println(isBalanced(text))
// }
