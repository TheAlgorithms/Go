// sha256.go
// description: The sha256 cryptographic hash function as defined in the RFC6234 standard.
// time complexity: O(n)
// space complexity: O(n)
// author: [Paul Leydier] (https://github.com/paul-leydier)
// ref: https://datatracker.ietf.org/doc/html/rfc6234
// ref: https://en.wikipedia.org/wiki/SHA-2
// see sha256_test.go

package sha256

import (
	"encoding/binary" // Used for interacting with uint at the byte level
	"math/bits"       // Used for bits rotation operations
)

var K = [64]uint32{
	0x428a2f98, 0x71374491, 0xb5c0fbcf, 0xe9b5dba5, 0x3956c25b, 0x59f111f1, 0x923f82a4, 0xab1c5ed5,
	0xd807aa98, 0x12835b01, 0x243185be, 0x550c7dc3, 0x72be5d74, 0x80deb1fe, 0x9bdc06a7, 0xc19bf174,
	0xe49b69c1, 0xefbe4786, 0x0fc19dc6, 0x240ca1cc, 0x2de92c6f, 0x4a7484aa, 0x5cb0a9dc, 0x76f988da,
	0x983e5152, 0xa831c66d, 0xb00327c8, 0xbf597fc7, 0xc6e00bf3, 0xd5a79147, 0x06ca6351, 0x14292967,
	0x27b70a85, 0x2e1b2138, 0x4d2c6dfc, 0x53380d13, 0x650a7354, 0x766a0abb, 0x81c2c92e, 0x92722c85,
	0xa2bfe8a1, 0xa81a664b, 0xc24b8b70, 0xc76c51a3, 0xd192e819, 0xd6990624, 0xf40e3585, 0x106aa070,
	0x19a4c116, 0x1e376c08, 0x2748774c, 0x34b0bcb5, 0x391c0cb3, 0x4ed8aa4a, 0x5b9cca4f, 0x682e6ff3,
	0x748f82ee, 0x78a5636f, 0x84c87814, 0x8cc70208, 0x90befffa, 0xa4506ceb, 0xbef9a3f7, 0xc67178f2,
}

const chunkSize = 64

// pad returns a padded version of the input message, such as the padded message's length is a multiple
// of 512 bits.
// The padding methodology is as follows:
// A "1" bit is appended at the end of the input message, followed by m "0" bits such as the length is
// 64 bits short of a 512 bits multiple. The remaining 64 bits are filled with the initial length of the
// message, represented as a 64-bits unsigned integer.
// For more details, see: https://datatracker.ietf.org/doc/html/rfc6234#section-4.1
func pad(message []byte) []byte {
	L := make([]byte, 8)
	binary.BigEndian.PutUint64(L, uint64(len(message)*8))
	message = append(message, 0x80) // "1" bit followed by 7 "0" bits
	for (len(message)+8)%64 != 0 {
		message = append(message, 0x00) // 8 "0" bits
	}
	message = append(message, L...)

	return message
}

// Hash hashes the input message using the sha256 hashing function, and return a 32 byte array.
// The implementation follows the RGC6234 standard, which is documented
// at https://datatracker.ietf.org/doc/html/rfc6234
func Hash(message []byte) [32]byte {
	message = pad(message)

	// Initialize round constants
	h0, h1, h2, h3, h4, h5, h6, h7 := uint32(0x6a09e667), uint32(0xbb67ae85), uint32(0x3c6ef372), uint32(0xa54ff53a),
		uint32(0x510e527f), uint32(0x9b05688c), uint32(0x1f83d9ab), uint32(0x5be0cd19)

	// Iterate through 512-bit chunks
	for chunkStart := 0; chunkStart < len(message); chunkStart += chunkSize {
		// Message schedule
		var w [64]uint32
		for i := 0; i*4 < chunkSize; i++ {
			w[i] = binary.BigEndian.Uint32(message[chunkStart+i*4 : chunkStart+(i+1)*4])
		}

		// Extend the 16 bytes chunk to the whole 64 bytes message schedule
		for i := 16; i < 64; i++ {
			s0 := bits.RotateLeft32(w[i-15], -7) ^ bits.RotateLeft32(w[i-15], -18) ^ (w[i-15] >> 3)
			s1 := bits.RotateLeft32(w[i-2], -17) ^ bits.RotateLeft32(w[i-2], -19) ^ (w[i-2] >> 10)
			w[i] = w[i-16] + s0 + w[i-7] + s1
		}

		// Actual hashing loop
		a, b, c, d, e, f, g, h := h0, h1, h2, h3, h4, h5, h6, h7
		for i := 0; i < 64; i++ {
			S1 := bits.RotateLeft32(e, -6) ^ bits.RotateLeft32(e, -11) ^ bits.RotateLeft32(e, -25)
			ch := (e & f) ^ ((^e) & g)
			tmp1 := h + S1 + ch + K[i] + w[i]
			S0 := bits.RotateLeft32(a, -2) ^ bits.RotateLeft32(a, -13) ^ bits.RotateLeft32(a, -22)
			maj := (a & b) ^ (a & c) ^ (b & c)
			tmp2 := S0 + maj
			h = g
			g = f
			f = e
			e = d + tmp1
			d = c
			c = b
			b = a
			a = tmp1 + tmp2
		}
		h0 += a
		h1 += b
		h2 += c
		h3 += d
		h4 += e
		h5 += f
		h6 += g
		h7 += h
	}

	// Export digest
	digest := [32]byte{}
	binary.BigEndian.PutUint32(digest[:4], h0)
	binary.BigEndian.PutUint32(digest[4:8], h1)
	binary.BigEndian.PutUint32(digest[8:12], h2)
	binary.BigEndian.PutUint32(digest[12:16], h3)
	binary.BigEndian.PutUint32(digest[16:20], h4)
	binary.BigEndian.PutUint32(digest[20:24], h5)
	binary.BigEndian.PutUint32(digest[24:28], h6)
	binary.BigEndian.PutUint32(digest[28:], h7)

	return digest
}
