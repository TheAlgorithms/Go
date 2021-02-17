// Package rot13 is a simple letter substitution cipher that replaces a letter with the 13th letter after it in the alphabet.
// ref: https://en.wikipedia.org/wiki/ROT13
package rot13

import (
	"github.com/tjgurwara99/Go/ciphers/caesar"
)

// rot13 is a special case, which is fixed the shift of 13, of the Caesar cipher
func rot13(input string) string {
	return caesar.Encrypt(input, 13)
}
