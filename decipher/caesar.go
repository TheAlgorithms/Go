package decipher

import "github.com/TheAlgorithms/Go/cipher"

// Caesar decrypts by left shift of "key" each character of "input"
func Caesar(input string, key int) (string, error) {
	// left shift of "key" is same as right shift of 26-"key"
	return cipher.Caesar(input, 26-key)
}
