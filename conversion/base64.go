// base64.go
// description: The base64 encoding algorithm as defined in the RFC4648 standard.
// author: [Paul Leydier] (https://github.com/paul-leydier)
// time complexity: O(n)
// space complexity: O(n)
// ref: https://datatracker.ietf.org/doc/html/rfc4648#section-4
// ref: https://en.wikipedia.org/wiki/Base64
// see base64_test.go

package conversion

import (
	"strings" // Used for efficient string builder (more efficient than simply appending strings)
)

const Alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"

// Base64Encode encodes the received input bytes slice into a base64 string.
// The implementation follows the RFC4648 standard, which is documented
// at https://datatracker.ietf.org/doc/html/rfc4648#section-4
func Base64Encode(input []byte) string {
	var sb strings.Builder
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
		// the literal 0x3F corresponds to the byte "0011 1111"
		// the operation "byte & 0x3F" masks the two left-most bits
		group := [4]byte{
			input[i] >> 2,
			(input[i]<<4)&0x3F + input[i+1]>>4,
			(input[i+1]<<2)&0x3F + input[i+2]>>6,
			input[i+2] & 0x3F,
		}

		// translate each group into a char using the static map
		for _, b := range group {
			sb.WriteString(string(Alphabet[int(b)]))
		}
	}
	encoded := sb.String()

	// Apply the output padding
	encoded = encoded[:len(encoded)-len(padding)] + padding[:]

	return encoded
}

// Base64Decode decodes the received input base64 string into a byte slice.
// The implementation follows the RFC4648 standard, which is documented
// at https://datatracker.ietf.org/doc/html/rfc4648#section-4
func Base64Decode(input string) []byte {
	padding := strings.Count(input, "=") // Number of bytes which will be ignored
	var decoded []byte

	// select 4 6-bit input groups, and re-arrange them into 3 8-bit groups
	for i := 0; i < len(input); i += 4 {
		// translate each group into a byte using the static map
		byteInput := [4]byte{
			byte(strings.IndexByte(Alphabet, input[i])),
			byte(strings.IndexByte(Alphabet, input[i+1])),
			byte(strings.IndexByte(Alphabet, input[i+2])),
			byte(strings.IndexByte(Alphabet, input[i+3])),
		}

		group := [3]byte{
			byteInput[0]<<2 + byteInput[1]>>4,
			byteInput[1]<<4 + byteInput[2]>>2,
			byteInput[2]<<6 + byteInput[3],
		}

		decoded = append(decoded, group[:]...)
	}

	return decoded[:len(decoded)-padding]
}
