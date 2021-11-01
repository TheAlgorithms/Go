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
}

func TestEncryptDecrypt(t *testing.T) {
	// Both prime numbers
	p := int64(61)
	q := int64(53)

	n := p * q

	delta := lcm.Lcm(p-1, q-1)

	e := int64(17) // Coprime with delta

	if gcd.Recursive(e, delta) != 1 {
		t.Fatal("Algorithm failed in preamble stage:\n\tPrime numbers are chosen statically and it shouldn't fail at this stage")
	}

	d, err := modular.Inverse(e, delta)

	if err != nil {
		t.Fatalf("Algorithm failed in preamble stage:\n\tProblem with a modular directory dependency: %v", err)
	}

	for _, test := range rsaTestData {
		t.Run(test.description, func(t *testing.T) {

			message := []rune(test.input)
			encrypted, err := Encrypt(message, e, n)
			if err != nil {
				t.Fatalf("Failed to Encrypt test string:\n\tDescription: %v\n\tErrMessage: %v", test.description, err)
			}

			decrypted, err := Decrypt(encrypted, d, n)
			if err != nil {
				t.Fatalf("Failed to Decrypt test message:\n\tDescription: %v\n\tErrMessage: %v", test.description, err)
			}

			if actual := test.input; actual != decrypted {
				t.Logf("FAIL: %s", test.description)
				t.Fatalf("Expecting %v, actual %v", decrypted, actual)
			}
		})
	}
}
