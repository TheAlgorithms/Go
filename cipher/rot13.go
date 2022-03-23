package cipher

// Rot13 is a special case, which is fixed the shift of 13, of the Caesar cipher
// ref: https://en.wikipedia.org/wiki/ROT13
func Rot13(input string) (string, error) {
	return Caesar(input, 13)
}
