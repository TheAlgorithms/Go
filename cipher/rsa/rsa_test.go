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

func testPrecondition(t *testing.T) (int64, int64, int64) {
	// Both prime numbers
	t.Helper()
	p := int64(61)
	q := int64(53)
	n := p * q
	delta := lcm.Lcm(p-1, q-1)
	e := int64(17) // Coprime with delta
	if gcd.Recursive(e, delta) != 1 {
		t.Fatal("something went wrong: prime numbers are chosen statically and it shouldn't fail at this stage")
	}
	d, err := modular.Inverse(e, delta)
	if err != nil {
		t.Fatalf("something went wrong: problem with %q: %v", "modular.Inverse", err)
	}
	return e, d, n
}

func TestEncryptDecrypt(t *testing.T) {
	for _, test := range rsaTestData {
		t.Run(test.description, func(t *testing.T) {
			e, d, n := testPrecondition(t)
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

func FuzzRsa(f *testing.F) {
	for _, rsaTestInput := range rsaTestData {
		f.Add(rsaTestInput.input)
	}
	f.Fuzz(func(t *testing.T, input string) {
		e, d, n := testPrecondition(t)
		encrypted, err := Encrypt([]rune(input), e, n)
		if err != nil {
			t.Fatalf("failed to encrypt string: %v", err)
		}
		decrypted, err := Decrypt(encrypted, d, n)
		if err != nil {
			t.Fatalf("failed to decrypt string: %v", err)
		}
		if decrypted != input {
			t.Fatalf("expected: %q, got: %q", input, decrypted)
		}
	})
}
