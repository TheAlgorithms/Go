package cipher

// Caesar encrypts by right shift of "key" each character of "input"
func Caesar(input string, key int) (string, error) {
	// if key is negative value, updates "key" the number which congruent to "key"
	// modulo 26
	key8 := byte(key%26+26) % 26

	var outputBuffer []byte
	// r is a rune, which is the equivalent of uint32.
	for _, r := range input {
		newByte := byte(r)
		if 'A' <= r && r <= 'Z' {
			outputBuffer = append(outputBuffer, 'A'+(newByte-'A'+key8)%26)
		} else if 'a' <= r && r <= 'z' {
			outputBuffer = append(outputBuffer, 'a'+(newByte-'a'+key8)%26)
		} else {
			outputBuffer = append(outputBuffer, newByte)
		}
	}
	return string(outputBuffer), nil
}
