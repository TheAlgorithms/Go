package rsa

import (
	"testing"
)

func FuzzRSA2RoundTrip(f *testing.F) {
	// Add seed corpus
	f.Add("Hello")
	f.Add("World")
	f.Add("A")
	f.Add("Test123")
	f.Add("")
	f.Add("Special chars: !@#")
	f.Add("Unicode: αβγ")

	f.Fuzz(func(t *testing.T, message string) {
		rsaInstance := New()

		if rsaInstance == nil {
			t.Fatal("Failed to create RSA instance")
		}

		encrypted := rsaInstance.EncryptString(message)
		decrypted := rsaInstance.DecryptString(encrypted)

		if decrypted != message {
			t.Errorf("RSA2 roundtrip failed:\n  Original: %q\n  Decrypted: %q", message, decrypted)
		}
	})
}
