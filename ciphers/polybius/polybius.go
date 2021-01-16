package main

import (
	"errors"
	"strings"
)

type polybius struct {
	size       int
	characters string
	key        string
}

func newPolybius(key string, size int, chars string) (*polybius, error) {
	key = strings.ToUpper(key)
	chars = strings.ToUpper(chars)[:size]
	if len(chars) != size {
		return nil, errors.New("chars must be as long as size")
	}
	if len(key) != size*size {
		return nil, errors.New("key must be as long as the size squared")
	}
	return &polybius{size, chars, key}, nil
}

func (p *polybius) encrypt(text string) string {
	chars := []rune(strings.ToUpper(text))
	encryptedText := ""
	for _, char := range chars {
		encryptedText += p.encipher(char)
	}
	return encryptedText
}

func (p *polybius) decrypt(text string) string {
	chars := []rune(strings.ToUpper(text))
	decryptedText := ""
	for i := 0; i < len(chars); i += 2 {
		decryptedText += p.decipher(chars[i : i+2])
	}
	return decryptedText
}

func (p *polybius) encipher(char rune) string {
	index := strings.IndexRune(p.key, char)
	row := index / p.size
	col := index % p.size
	chars := []rune(p.characters)
	return string([]rune{chars[row], chars[col]})
}

func (p *polybius) decipher(chars []rune) string {
	if len(chars) != 2 {
		panic("decipher takes two chars")
	}
	row := strings.IndexRune(p.characters, chars[0])
	col := strings.IndexRune(p.characters, chars[1])
	return string([]rune(p.key)[row*p.size+col])
}
