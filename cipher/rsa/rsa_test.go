// rsa_test.go
// description: Test for RSA Encrypt and Decrypt algorithms
// author(s) [Taj](https://github.com/tjgurwara99)
// see rsa.go

package rsa

import (
	"testing"

	"github.com/TheAlgorithms/Go/math/gcd"
	"github.com/TheAlgorithms/Go/math/lcm"
	"github.com/TheAlgorithms/Go/math/modular"
)

var rsaTestData = []struct {
	description string
	input       string
}{
	{
		"Encrypt letter 'a'",
		"a",
	},
	{
		"Encrypt 'hello world'",
		"hello world",
	},
	{
		"Encrypt full sentence",
		"the quick brown fox jumps over the lazy dog.",
	},
	{
		"Encrypt full sentence from rsacipher.go main function",
		"I think RSA is really great",
	},
	{
		"Encrypt uncommon unicode",
		string("\x82"),
	},
}

func FuzzEncryptDecrypt(f *testing.F) {
	// Both prime numbers
	p := int64(3079)
	q := int64(983)

	n := p * q

	delta := lcm.Lcm(p-1, q-1)

	e := int64(2837) // Coprime with delta

	if gcd.Recursive(e, delta) != 1 {
		f.Fatal("Algorithm failed in preamble stage:\n\tPrime numbers are chosen statically and it shouldn't fail at this stage")
	}

	d, err := modular.Inverse(e, delta)

	if err != nil {
		f.Fatalf("Algorithm failed in preamble stage:\n\tProblem with a modular directory dependency: %v", err)
	}

	for _, test := range rsaTestData {
		f.Add(test.input)
	}

	f.Fuzz(func(t *testing.T, input string) {
		message := []rune(input)
		encrypted, err := Encrypt(message, e, n)
		if err != nil {
			t.Fatalf("Failed to Encrypt test string:\n\tInput:%s\n\tErrMessage: %v", input, err)
		}

		decrypted, err := Decrypt(encrypted, d, n)
		if err != nil {
			t.Fatalf("Failed to Decrypt test message:\n\tInput:%s\n\tErrMessage: %v", input, err)
		}

		if string(message) != decrypted {
			t.Fatalf("Decrypted ciphertext does not match input: Expecting %v, actual %v", []byte(input), []byte(decrypted))
		}
	})

}
