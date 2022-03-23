package cipher

import (
	"fmt"
	"math"
	"strings"
)

// polybius is struct having size, characters, and key
type polybius struct {
	size       int
	characters string
	key        string
}

// NewPolybius returns a pointer to object of polybius.
// If the size of "chars" is longer than "size",
// "chars" are truncated to "size".
func NewPolybius(key string, size int, chars string) (*polybius, error) {
	key = strings.ToUpper(key)
	chars = strings.ToUpper(chars)[:size]
	for idx, ch := range chars {
		if strings.Contains(chars[idx+1:], string(ch)) {
			return nil, fmt.Errorf("\"chars\" contains same character: %c", ch)
		}
	}

	if len(key) != size*size {
		return nil, fmt.Errorf("len(key): %d must be as long as size squared: %d", len(key), size*size)
	}
	return &polybius{size, chars, key}, nil
}

// Encrypt encrypts with polybius encryption
func (p *polybius) Encrypt(text string) (string, error) {
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
func (p *polybius) Decrypt(text string) (string, error) {
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

func (p *polybius) encipher(char rune) (string, error) {
	index := strings.IndexRune(p.key, char)
	if index < 0 {
		return "", fmt.Errorf("%c does not exist in keys", char)
	}
	row := index / p.size
	col := index % p.size
	chars := []rune(p.characters)
	return string([]rune{chars[row], chars[col]}), nil
}

func (p *polybius) decipher(chars []rune) (string, error) {
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
	return string([]rune(p.key)[row*p.size+col]), nil
}
