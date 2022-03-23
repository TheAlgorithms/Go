package decipher_test

import (
	"errors"
	"github.com/TheAlgorithms/Go/cipher"
	"github.com/TheAlgorithms/Go/decipher"
	"math/rand"
	"testing"
)

const enAlphabet = "abcdefghijklmnopqrstuvwxyz "

func getTexts() []string {
	return []string{
		"Ilya Sokolov",
		"A slice literal is declared just like an array literal, except you leave out the element count",
		"Go is an open source programming language that makes it easy to build simple, reliable, and efficient software.",
		"Go’s treatment of errors as values has served us well over the last decade. Although the standard library’s support for errors has been minimal—just the errors.New and fmt.Errorf functions, which produce errors that contain only a message—the built-in error interface allows Go programmers to add whatever information they desire. All it requires is a type that implements an Error method:",
		"А тут для примера русский текст",
	}
}

func getRandomString() string {
	enRunes := []rune(enAlphabet)
	b := make([]rune, rand.Intn(100))
	for i := range b {
		b[i] = enRunes[rand.Intn(len(enRunes))]
	}
	return string(b)
}
func TestTransposition(t *testing.T) {
	for _, s := range getTexts() {
		keyWord := getRandomString()
		encrypt, errEncrypt := cipher.Transposition([]rune(s), keyWord)
		if errEncrypt != nil &&
			!errors.Is(errEncrypt, cipher.NoTextToEncryptError) &&
			!errors.Is(errEncrypt, cipher.KeyMissingError) {
			t.Error("Unexpected error ", errEncrypt)
		}
		if errEncrypt != nil {
			t.Error(errEncrypt)
		}
		decrypt, errDecrypt := decipher.Transposition([]rune(encrypt), keyWord)
		if errDecrypt != nil &&
			!errors.Is(errDecrypt, decipher.NoTextToDecryptError) &&
			!errors.Is(errDecrypt, decipher.KeyMissingError) {
			t.Error("Unexpected error ", errDecrypt)
		}
		if errDecrypt != nil {
			t.Error(errDecrypt)
		}
		if encrypt == decrypt {
			t.Error("String ", s, " not encrypted")
		}
		if encrypt == s {
			t.Error("String ", s, " not encrypted")
		}
	}
}
