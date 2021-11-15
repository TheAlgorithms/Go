// Package base64
// base64.go
// description: The base64 encodage algorithm as defined in the RFC4648 standard.
// author: [Paul Leydier] (https://github.com/paul-leydier)
// ref: https://datatracker.ietf.org/doc/html/rfc4648#section-4
// ref: https://en.wikipedia.org/wiki/Base64
// see base64_test.go
package base64

const Alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"

// Encode encodes the received input bytes into a base64 string.
// The implementation follows the RFC4648 standard, which is documented
// at https://datatracker.ietf.org/doc/html/rfc4648#section-4
func Encode(input []byte) (encoded string) {
	// If not 24 bits (3 bytes) multiple, pad with 0 value bytes, and with "=" for the output
	var padding string
	for i := len(input) % 3; i > 0 && i < 3; i++ {
		var zeroByte byte
		input = append(input, zeroByte)
		padding += "="
	}

	// encode 24 bits per 24 bits (3 bytes per 3 bytes)
	for i := 0; i < len(input); i += 3 {
		// select 3 8-bit input groups, and re-arrange them into 4 6-bit groups
		group := [4]byte{
			input[i] >> 2,
			(input[i]<<6)>>2 + input[i+1]>>4,
			(input[i+1]<<4)>>2 + input[i+2]>>6,
			(input[i+2] << 2) >> 2,
		}

		// translate each groups into a char using the static map
		for _, b := range group {
			encoded = encoded + string(Alphabet[int(b)])
		}
	}

	// Apply the output padding
	encoded = encoded[:len(encoded)-len(padding)] + padding[:]

	return
}
