package decipher_test

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
			"Basic caesar decryption with letter 'a'",
			"a",
			3,
			"x",
		},
		{
			"Basic caesar decryption wrap around alphabet on letter 'z'",
			"z",
			3,
			"w",
		},
		{
			"Decrypt a simple string with caesar encryiption",
			"hello",
			3,
			"ebiil",
		},
		{
			"Decrypt a simple string with key 13",
			"hello",
			13,
			"uryyb",
		},
		{
			"Decrypt a simple string with key -13",
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
			"Decrypt sentence with key 10",
			"Dro Aesmu Lbygx Pyh Tewzc yfob dro Vkji Nyq.",
			10,
			"The Quick Brown Fox Jumps over the Lazy Dog.",
		},
	}

	for _, test := range caesarTestData {
		t.Run(test.description, func(t *testing.T) {
			actual, err := decipher.Caesar(test.input, test.key)
			if err != nil {
				t.Fatalf("FAIL: no errors expected: %s", err)
			}
			if actual != test.expected {
				t.Fatalf("FAIL: with input string '%s' and key '%d' was expected '%s' but actual was '%s'",
					test.input, test.key, test.expected, actual)
			}
		})
	}
}

func ExampleCaesar() {
	const (
		key   = 10
		input = "The Quick Brown Fox Jumps over the Lazy Dog."
	)

	encryptedText, _ := cipher.Caesar(input, key)

	decryptedText, _ := decipher.Caesar(encryptedText, key)
	fmt.Printf("Decrypt=> key: %d, input: %s, decryptedText: %s\n", key, encryptedText, decryptedText)

	// Output:
	// Decrypt=> key: 10, input: Dro Aesmu Lbygx Pyh Tewzc yfob dro Vkji Nyq., decryptedText: The Quick Brown Fox Jumps over the Lazy Dog.
}
