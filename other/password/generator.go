// This program generates a password from a list of possible chars
// You must provide a minimum length and a maximum length
// This length is not fixed if you generate multiple passwords for the same range

// Package password contains functions to help generate random passwords
// time complexity: O(n)
// space complexity: O(n)

package password

import (
	"crypto/rand"
	"io"
	"math/big"
)

// Generate returns a newly generated password
func Generate(minLength int, maxLength int) string {
	var chars = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789!@#$%^&*()-_=+,.?/:;{}[]`~")

	length, err := rand.Int(rand.Reader, big.NewInt(int64(maxLength-minLength)))
	if err != nil {
		panic(err) // handle this gracefully
	}
	length.Add(length, big.NewInt(int64(minLength)))

	intLength := int(length.Int64())

	newPassword := make([]byte, intLength)
	randomData := make([]byte, intLength+intLength/4)
	charLen := byte(len(chars))
	maxrb := byte(256 - (256 % len(chars)))
	i := 0
	for {
		if _, err := io.ReadFull(rand.Reader, randomData); err != nil {
			panic(err)
		}
		for _, c := range randomData {
			if c >= maxrb {
				continue
			}
			newPassword[i] = chars[c%charLen]
			i++
			if i == intLength {
				return string(newPassword)
			}
		}
	}
}
