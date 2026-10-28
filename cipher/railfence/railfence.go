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
	runes := []rune(text)
	if rails <= 1 || rails >= len(runes) {
		return text
	}

	rows := make([][]rune, rails)

	dirDown := false
	row := 0
	for _, char := range runes {
		rows[row] = append(rows[row], char)
		if row == 0 || row == rails-1 {
			dirDown = !dirDown
		}
		if dirDown {
			row++
		} else {
			row--
		}
	}

	var result strings.Builder
	for _, r := range rows {
		result.WriteString(string(r))
	}

	return result.String()
}

func Decrypt(cipherText string, rails int) string {
	runes := []rune(cipherText)
	if rails <= 1 || rails >= len(runes) {
		return cipherText
	}

	// Placeholder for the decrypted message
	decrypted := make([]rune, len(runes))

	// Calculate the zigzag pattern and place characters accordingly
	index := 0
	for rail := 0; rail < rails; rail++ {
		position := rail
		down := true // Direction flag
		for position < len(runes) {
			decrypted[position] = runes[index]
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
