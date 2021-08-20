// Package xor_cipher is an encryption algorithm that operates the exclusive disjunction(XOR)
// ref: https://en.wikipedia.org/wiki/XOR_cipher
package xor_cipher

// Encrypt encrypts with Xor encryption after converting each character to byte
// The returned value might not be readable because there is no guarantee
// which is within the ASCII range
// If using other type such as string, []int, or some other types,
// add the statements for converting the type to []byte.
func Encrypt(key byte, plaintext []byte) []byte {
	cipherText := []byte{}
	for _, ch := range plaintext {
		cipherText = append(cipherText, key^ch)
	}
	return cipherText
}

// Decrypt decrypts with Xor encryption
func Decrypt(key byte, cipherText []byte) []byte {
	plainText := []byte{}
	for _, ch := range cipherText {
		plainText = append(plainText, key^ch)
	}
	return plainText
}
