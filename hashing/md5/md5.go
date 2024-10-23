// md5.go
// description: The MD5 hashing function as defined in RFC 1321.
// author: Simon Waldherr
// ref: https://datatracker.ietf.org/doc/html/rfc1321
// see md5_test.go for testing

package md5

import (
	"encoding/binary"
)

// Constants for MD5
var (
	s = [64]uint32{
		7, 12, 17, 22, 7, 12, 17, 22, 7, 12, 17, 22, 7, 12, 17, 22,
		5, 9, 14, 20, 5, 9, 14, 20, 5, 9, 14, 20, 5, 9, 14, 20,
		4, 11, 16, 23, 4, 11, 16, 23, 4, 11, 16, 23, 4, 11, 16, 23,
		6, 10, 15, 21, 6, 10, 15, 21, 6, 10, 15, 21, 6, 10, 15, 21,
	}

	K = [64]uint32{
		0xd76aa478, 0xe8c7b756, 0x242070db, 0xc1bdceee,
		0xf57c0faf, 0x4787c62a, 0xa8304613, 0xfd469501,
		0x698098d8, 0x8b44f7af, 0xffff5bb1, 0x895cd7be,
		0x6b901122, 0xfd987193, 0xa679438e, 0x49b40821,
		0xf61e2562, 0xc040b340, 0x265e5a51, 0xe9b6c7aa,
		0xd62f105d, 0x02441453, 0xd8a1e681, 0xe7d3fbc8,
		0x21e1cde6, 0xc33707d6, 0xf4d50d87, 0x455a14ed,
		0xa9e3e905, 0xfcefa3f8, 0x676f02d9, 0x8d2a4c8a,
		0xfffa3942, 0x8771f681, 0x6d9d6122, 0xfde5380c,
		0xa4beea44, 0x4bdecfa9, 0xf6bb4b60, 0xbebfbc70,
		0x289b7ec6, 0xeaa127fa, 0xd4ef3085, 0x04881d05,
		0xd9d4d039, 0xe6db99e5, 0x1fa27cf8, 0xc4ac5665,
		0xf4292244, 0x432aff97, 0xab9423a7, 0xfc93a039,
		0x655b59c3, 0x8f0ccc92, 0xffeff47d, 0x85845dd1,
		0x6fa87e4f, 0xfe2ce6e0, 0xa3014314, 0x4e0811a1,
		0xf7537e82, 0xbd3af235, 0x2ad7d2bb, 0xeb86d391,
	}
)

// leftRotate rotates x left by n bits
func leftRotate(x, n uint32) uint32 {
	return (x << n) | (x >> (32 - n))
}

// pad pads the input message so that its length is congruent to 448 modulo 512
func pad(message []byte) []byte {
	originalLength := len(message) * 8
	message = append(message, 0x80)
	for (len(message)*8)%512 != 448 {
		message = append(message, 0x00)
	}

	lengthBytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(lengthBytes, uint64(originalLength))
	message = append(message, lengthBytes...)

	return message
}

// Hash computes the MD5 hash of the input message
func Hash(message []byte) [16]byte {
	message = pad(message)

	// Initialize MD5 state variables
	a0, b0, c0, d0 := uint32(0x67452301), uint32(0xefcdab89), uint32(0x98badcfe), uint32(0x10325476)

	// Process the message in successive 512-bit chunks
	for i := 0; i < len(message); i += 64 {
		chunk := message[i : i+64]
		var M [16]uint32
		for j := 0; j < 16; j++ {
			M[j] = binary.LittleEndian.Uint32(chunk[j*4 : (j+1)*4])
		}

		// Initialize hash value for this chunk
		A, B, C, D := a0, b0, c0, d0

		// Main loop
		for i := 0; i < 64; i++ {
			var F, g uint32
			if i < 16 {
				F = (B & C) | ((^B) & D)
				g = uint32(i)
			} else if i < 32 {
				F = (D & B) | ((^D) & C)
				g = uint32((5*i + 1) % 16)
			} else if i < 48 {
				F = B ^ C ^ D
				g = uint32((3*i + 5) % 16)
			} else {
				F = C ^ (B | (^D))
				g = uint32((7 * i) % 16)
			}
			F = F + A + K[i] + M[g]
			A = D
			D = C
			C = B
			B = B + leftRotate(F, s[i])
		}

		// Add this chunk's hash to result so far
		a0 += A
		b0 += B
		c0 += C
		d0 += D
	}

	// Produce the final hash value (digest)
	var digest [16]byte
	binary.LittleEndian.PutUint32(digest[0:4], a0)
	binary.LittleEndian.PutUint32(digest[4:8], b0)
	binary.LittleEndian.PutUint32(digest[8:12], c0)
	binary.LittleEndian.PutUint32(digest[12:16], d0)

	return digest
}
