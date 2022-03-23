package cipher_test

import (
	"fmt"
	"github.com/TheAlgorithms/Go/cipher"
	"github.com/TheAlgorithms/Go/decipher"
	"testing"
)

func TestCaesar(t *testing.T) {
	var caesarTestData = []struct {
		description string
		input       string
		key         int
		expected    string
	}{
		{
			"Basic caesar encryption with letter 'a'",
			"a",
			3,
			"d",
		},
		{
			"Basic caesar encryption wrap around alphabet on letter 'z'",
			"z",
			3,
			"c",
		},
		{
			"Encrypt a simple string with caesar encryiption",
			"hello",
			3,
			"khoor",
		},
		{
			"Encrypt a simple string with key 13",
			"hello",
			13,
			"uryyb",
		},
		{
			"Encrypt a simple string with key -13",
			"hello",
			-13,
			"uryyb",
		},
		{
			"With key of 26 output should be the same as the input",
			"no change",
			26,
			"no change",
		},
		{
			"Encrypt sentence with key 10",
			"the quick brown fox jumps over the lazy dog.",
			10,
			"dro aesmu lbygx pyh tewzc yfob dro vkji nyq.",
		},
		{
			"Encrypt sentence with key 10",
			"The Quick Brown Fox Jumps over the Lazy Dog.",
			10,
			"Dro Aesmu Lbygx Pyh Tewzc yfob dro Vkji Nyq.",
		},
	}
	for _, test := range caesarTestData {
		t.Run(test.description, func(t *testing.T) {
			actual, err := cipher.Caesar(test.input, test.key)
			if err != nil {
				t.Logf("FAIL: No errors expected: %s", err)
				t.Fatalf("FAIL: %s", test.description)
			}
			if actual != test.expected {
				t.Logf("FAIL: %s", test.description)
				t.Fatalf("With input string '%s' and key '%d' was expecting '%s' but actual was '%s'",
					test.input, test.key, test.expected, actual)
			}
		})
	}
}

func Example() {
	const (
		key   = 10
		input = "The Quick Brown Fox Jumps over the Lazy Dog."
	)

	encryptedText, _ := cipher.Caesar(input, key)
	fmt.Printf("Encrypt=> key: %d, input: %s, encryptedText: %s\n", key, input, encryptedText)

	decryptedText, _ := decipher.Caesar(encryptedText, key)
	fmt.Printf("Decrypt=> key: %d, input: %s, decryptedText: %s\n", key, encryptedText, decryptedText)

	// Output:
	// Encrypt=> key: 10, input: The Quick Brown Fox Jumps over the Lazy Dog., encryptedText: Dro Aesmu Lbygx Pyh Tewzc yfob dro Vkji Nyq.
	// Decrypt=> key: 10, input: Dro Aesmu Lbygx Pyh Tewzc yfob dro Vkji Nyq., decryptedText: The Quick Brown Fox Jumps over the Lazy Dog.
}
