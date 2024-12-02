// Package polybius is encrypting method with polybius square
// description: Polybius square
// details : The Polybius algorithm is a simple algorithm that is used to encode a message by converting each letter to a pair of numbers.
// time complexity: O(n)
// space complexity: O(n)
// ref: https://en.wikipedia.org/wiki/Polybius_square#Hybrid_Polybius_Playfair_Cipher
package polybius

import (
	"fmt"
	"math"
	"strings"
)

// Polybius is struct having size, characters, and key
type Polybius struct {
	size       int
	characters string
	key        string
}

// NewPolybius returns a pointer to object of Polybius.
// If the size of "chars" is longer than "size",
// "chars" are truncated to "size".
func NewPolybius(key string, size int, chars string) (*Polybius, error) {
	if size < 0 {
		return nil, fmt.Errorf("provided size %d cannot be negative", size)
	}
	key = strings.ToUpper(key)
	if size > len(chars) {
		return nil, fmt.Errorf("provided size %d is too small to use to slice string %q of len %d", size, chars, len(chars))
	}
	for _, r := range chars {
		if (r < 'a' || r > 'z') && (r < 'A' || r > 'Z') {
			return nil, fmt.Errorf("provided string %q should only contain latin characters", chars)
		}
	}
	chars = strings.ToUpper(chars)[:size]
	for i, r := range chars {
		if strings.ContainsRune(chars[i+1:], r) {
			return nil, fmt.Errorf("%q contains same character %q", chars[i+1:], r)
		}
	}

	if len(key) != size*size {
		return nil, fmt.Errorf("len(key): %d must be as long as size squared: %d", len(key), size*size)
	}
	return &Polybius{size, chars, key}, nil
}

// Encrypt encrypts with polybius encryption
func (p *Polybius) Encrypt(text string) (string, error) {
	encryptedText := ""
	for _, char := range strings.ToUpper(text) {
		encryptedChar, err := p.encipher(char)
		if err != nil {
			return "", fmt.Errorf("failed encipher: %w", err)
		}
		encryptedText += encryptedChar
	}
	return encryptedText, nil
}

// Decrypt decrypts with polybius encryption
func (p *Polybius) Decrypt(text string) (string, error) {
	chars := []rune(strings.ToUpper(text))
	decryptedText := ""
	for i := 0; i < len(chars); i += 2 {
		decryptedChar, err := p.decipher(chars[i:int(math.Min(float64(i+2), float64(len(chars))))])
		if err != nil {
			return "", fmt.Errorf("failed decipher: %w", err)
		}
		decryptedText += decryptedChar
	}
	return decryptedText, nil
}

func (p *Polybius) encipher(char rune) (string, error) {
	index := strings.IndexRune(p.key, char)
	if index < 0 {
		return "", fmt.Errorf("%q does not exist in keys", char)
	}
	row := index / p.size
	col := index % p.size
	chars := []rune(p.characters)
	return string([]rune{chars[row], chars[col]}), nil
}

func (p *Polybius) decipher(chars []rune) (string, error) {
	if len(chars) != 2 {
		return "", fmt.Errorf("the size of \"chars\" must be even")
	}
	row := strings.IndexRune(p.characters, chars[0])
	if row < 0 {
		return "", fmt.Errorf("%c does not exist in characters", chars[0])
	}
	col := strings.IndexRune(p.characters, chars[1])
	if col < 0 {
		return "", fmt.Errorf("%c does not exist in characters", chars[1])
	}
	return string(p.key[row*p.size+col]), nil
}
