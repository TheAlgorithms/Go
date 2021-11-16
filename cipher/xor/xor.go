// Package xor is an encryption algorithm that operates the exclusive disjunction(XOR)
// ref: https://en.wikipedia.org/wiki/XOR_cipher
package xor

import "unsafe"

// Encrypt encrypts with Xor encryption after converting each character to byte
// The returned value might not be readable because there is no guarantee
// which is within the ASCII range
// If using other type such as string, []int, or some other types,
// add the statements for converting the type to []byte.
func Encrypt(key byte, plaintext []byte) []byte {
	cipherText := make([]byte, len(plaintext))
	copy(cipherText, plaintext)

	uint64Plaintext := *(*[]uint64)(unsafe.Pointer(&cipherText))

	uint64Key := uint64(key) +
		uint64(key)<<(8*1) +
		uint64(key)<<(8*2) +
		uint64(key)<<(8*3) +
		uint64(key)<<(8*4) +
		uint64(key)<<(8*5) +
		uint64(key)<<(8*6) +
		uint64(key)<<(8*7)

	n := len(cipherText) / 8

	// less than 64byte of data
	nx := n % 8
	for i := 0; i < nx; i++ {
		uint64Plaintext[i] ^= uint64Key
	}

	// loop once for 8*8=64 bytes
	for i := nx; i < n; i += 8 {
		slice := uint64Plaintext[i : i+8]
		slice[0] ^= uint64Key
		slice[1] ^= uint64Key
		slice[2] ^= uint64Key
		slice[3] ^= uint64Key
		slice[4] ^= uint64Key
		slice[5] ^= uint64Key
		slice[6] ^= uint64Key
		slice[7] ^= uint64Key
	}

	// data with less than 8 bytes at the end
	n = len(cipherText) % 8
	for i := len(cipherText) - n; i < len(cipherText); i++ {
		cipherText[i] ^= key
	}
	return cipherText
}

// Decrypt decrypts with Xor encryption
func Decrypt(key byte, cipherText []byte) []byte {
	return Encrypt(key, cipherText)
}
