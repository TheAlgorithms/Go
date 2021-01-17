// Package caesar is the shift cipher
// ref: https://en.wikipedia.org/wiki/Caesar_cipher
package caesar

// Caesar is struct for Caesar cipher
type Caesar struct {
}

// NewCaesar returns a pointer to object of Caesar
func NewCaesar() *Caesar {
	return &Caesar{}
}

// Encrypt encrypts by right shift of "key" each character of "input"
func (c Caesar) Encrypt(input string, key int) string {
	// if key is negative value,
	// updates "key" the number which congruents to "key" modulo 26
	key = (key%26 + 26) % 26

	outputBuffer := []byte{}
	for _, r := range input {
		newByte := byte(r)
		if 'A' <= newByte && newByte <= 'Z' {
			outputBuffer = append(outputBuffer, byte(int('A')+int(int(newByte-'A')+key)%26))
		} else if 'a' <= newByte && newByte <= 'z' {
			outputBuffer = append(outputBuffer, byte(int('a')+int(int(newByte-'a')+key)%26))
		} else {
			outputBuffer = append(outputBuffer, newByte)
		}
	}
	return string(outputBuffer)
}

// Decrypt decrypts by left shift of "key" each character of "input"
func (c Caesar) Decrypt(input string, key int) string {
	// left shift of "key" is same as right shift of 26-"key"
	return c.Encrypt(input, 26-key)
}
