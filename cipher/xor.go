package cipher

// Xor encrypts with Xor encryption after converting each character to byte
// The returned value might not be readable because there is no guarantee
// which is within the ASCII range
// If using other type such as string, []int, or some other types,
// add the statements for converting the type to []byte.
func Xor(key byte, plaintext []byte) ([]byte, error) {
	var cipherText = make([]byte, len(plaintext))
	for i, ch := range plaintext {
		cipherText[i] = key ^ ch
	}
	return cipherText, nil
}
