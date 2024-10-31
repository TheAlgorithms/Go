// railfence.go
// description: Rail Fence Cipher
// details: The rail fence cipher is a an encryption algorithm that uses a rail fence pattern to encode a message. it is a type of transposition cipher that rearranges the characters of the plaintext to form the ciphertext.
// time complexity: O(n)
// space complexity: O(n)
// ref: https://en.wikipedia.org/wiki/Rail_fence_cipher
package railfence

import (
	"strings"
)

func Encrypt(text string, rails int) string {
	if rails == 1 {
		return text
	}

	// Create a matrix for the rail fence pattern
	matrix := make([][]rune, rails)
	for i := range matrix {
		matrix[i] = make([]rune, len(text))
	}

	// Fill the matrix
	dirDown := false
	row, col := 0, 0
	for _, char := range text {
		if row == 0 || row == rails-1 {
			dirDown = !dirDown
		}
		matrix[row][col] = char
		col++
		if dirDown {
			row++
		} else {
			row--
		}
	}
	var result strings.Builder
	for _, line := range matrix {
		for _, char := range line {
			if char != 0 {
				result.WriteRune(char)
			}
		}
	}

	return result.String()
}
func Decrypt(cipherText string, rails int) string {
	if rails == 1 || rails >= len(cipherText) {
		return cipherText
	}

	// Placeholder for the decrypted message
	decrypted := make([]rune, len(cipherText))

	// Calculate the zigzag pattern and place characters accordingly
	index := 0
	for rail := 0; rail < rails; rail++ {
		position := rail
		down := true // Direction flag
		for position < len(cipherText) {
			decrypted[position] = rune(cipherText[index])
			index++

			// Determine the next position based on the current rail and direction
			if rail == 0 || rail == rails-1 {
				position += 2 * (rails - 1)
			} else if down {
				position += 2 * (rails - 1 - rail)
				down = false
			} else {
				position += 2 * rail
				down = true
			}
		}
	}

	return string(decrypted)
}
