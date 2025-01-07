// sha1.go
// description: The SHA-1 hashing function as defined in RFC 3174.
// author: Simon Waldherr
// ref: https://datatracker.ietf.org/doc/html/rfc3174
// see sha1_test.go for testing

package sha1

import (
	"encoding/binary" // Used for interacting with uint at the byte level
)

// Constants for SHA-1
const (
	h0 uint32 = 0x67452301
	h1 uint32 = 0xEFCDAB89
	h2 uint32 = 0x98BADCFE
	h3 uint32 = 0x10325476
	h4 uint32 = 0xC3D2E1F0
)

// pad pads the input message so that its length is congruent to 448 modulo 512
func pad(message []byte) []byte {
	originalLength := len(message) * 8
	message = append(message, 0x80)
	for (len(message)*8)%512 != 448 {
		message = append(message, 0x00)
	}

	lengthBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(lengthBytes, uint64(originalLength))
	message = append(message, lengthBytes...)

	return message
}

// leftRotate rotates x left by n bits
func leftRotate(x, n uint32) uint32 {
	return (x << n) | (x >> (32 - n))
}

// Hash computes the SHA-1 hash of the input message
func Hash(message []byte) [20]byte {
	message = pad(message)

	// Initialize variables
	a, b, c, d, e := h0, h1, h2, h3, h4

	// Process the message in successive 512-bit chunks
	for i := 0; i < len(message); i += 64 {
		var w [80]uint32
		chunk := message[i : i+64]

		// Break chunk into sixteen 32-bit big-endian words
		for j := 0; j < 16; j++ {
			w[j] = binary.BigEndian.Uint32(chunk[j*4 : (j+1)*4])
		}

		// Extend the sixteen 32-bit words into eighty 32-bit words
		for j := 16; j < 80; j++ {
			w[j] = leftRotate(w[j-3]^w[j-8]^w[j-14]^w[j-16], 1)
		}

		// Initialize hash value for this chunk
		A, B, C, D, E := a, b, c, d, e

		// Main loop
		for j := 0; j < 80; j++ {
			var f, k uint32
			switch {
			case j < 20:
				f = (B & C) | ((^B) & D)
				k = 0x5A827999
			case j < 40:
				f = B ^ C ^ D
				k = 0x6ED9EBA1
			case j < 60:
				f = (B & C) | (B & D) | (C & D)
				k = 0x8F1BBCDC
			default:
				f = B ^ C ^ D
				k = 0xCA62C1D6
			}

			temp := leftRotate(A, 5) + f + E + k + w[j]
			E = D
			D = C
			C = leftRotate(B, 30)
			B = A
			A = temp
		}

		// Add this chunk's hash to result so far
		a += A
		b += B
		c += C
		d += D
		e += E
	}

	// Produce the final hash value (digest)
	var digest [20]byte
	binary.BigEndian.PutUint32(digest[0:4], a)
	binary.BigEndian.PutUint32(digest[4:8], b)
	binary.BigEndian.PutUint32(digest[8:12], c)
	binary.BigEndian.PutUint32(digest[12:16], d)
	binary.BigEndian.PutUint32(digest[16:20], e)

	return digest
}
