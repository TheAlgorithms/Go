// transposition_test.go
// description: Transposition cipher
// author(s) [red_byte](https://github.com/i-redbyte)
// see transposition.go

package transposition

import (
	"math/rand"
	"strings"
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
	rusRune := []rune(enAlphabet)
	b := make([]rune, rand.Intn(100))
	for i := range b {
		b[i] = rusRune[rand.Intn(len(rusRune))]
	}
	return string(b)
}

func TestEncrypt(t *testing.T) {
	fn := func(text string, keyWord string) (bool, error) {
		encrypt, err := Encrypt([]rune(text), keyWord)
		return text == encrypt, err
	}
	for _, s := range getTexts() {
		if check, err := fn(s, getRandomString()); check || err != nil {
			t.Error("String ", s, " not encrypted")
		}
	}
	if _, err := fn(getRandomString(), ""); err == nil {
		t.Error("Error! empty string encryption")
	}
}

func TestDecrypt(t *testing.T) {
	for _, s := range getTexts() {
		keyWord := getRandomString()
		encrypt, err := Encrypt([]rune(s), keyWord)
		decrypt, _ := Decrypt([]rune(encrypt), keyWord)
		if err != nil {
			t.Error(err)
		}
		if encrypt == decrypt {
			t.Error("String ", s, " not encrypted")
		}
		if encrypt == s {
			t.Error("String ", s, " not encrypted")
		}
	}
}

func TestEncryptDecrypt(t *testing.T) {
	text := "Test text for checking the algorithm"
	key1 := "testKey"
	key2 := "Test Key2"
	encrypt, _ := Encrypt([]rune(text), key1)
	decrypt, _ := Decrypt([]rune(encrypt), key1)
	if strings.Contains(decrypt, text) == false {
		t.Error("The string was not decrypted correctly")
	}
	decrypt, _ = Decrypt([]rune(encrypt), key2)
	if strings.Contains(decrypt, text) == true {
		t.Error("The string was decrypted with a different key")
	}
}
