package decipher

// Xor decrypts with Xor encryption
func Xor(key byte, cipherText []byte) ([]byte, error) {
	var plainText = make([]byte, len(cipherText))
	for i, ch := range cipherText {
		plainText[i] = key ^ ch
	}
	return plainText, nil
}
