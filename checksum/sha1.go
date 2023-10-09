// sha1.go
// description: SHA1 algorithm
// details: is a Secure Hash Algorithm 1 is a cryptographic algorithm which takes an input and produces a 160-bit (20-byte) hash value. [source](https://www.google.com)
// See more [SHA1](https://en.wikipedia.org/wiki/SHA-1)
// author(s) [red_byte](https://github.com/i-redbyte)
// see sha1_test.go

// Package checksum describes algorithms for finding various checksums
package checksum

// Importing necessary existing libraries to support SHA1
import (
	"crypto/sha1"
	"encoding/hex"
)

// SHA1 validates the provided data using the SHA1 algorithm.
func Sha1(s string) string {

	// Create a new SHA-1 hash object
	sha1Hash := sha1.New()

	// Write the data to the hash object
	_, err := sha1Hash.Write([]byte(s))
	if err != nil {
		return ""
	}

	// Get the SHA-1 checksum
	checksum := sha1Hash.Sum(nil)

	// Convert the checksum to a hex string
	checksumString := hex.EncodeToString(checksum)

	// Returns the 20 byte hash value
	return checksumString
}
