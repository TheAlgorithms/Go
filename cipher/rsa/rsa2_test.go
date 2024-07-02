package rsa_test

import (
	"testing"

	"github.com/TheAlgorithms/Go/cipher/rsa"
)

func TestRSA(t *testing.T) {
	tests := []struct {
		name    string
		message string
	}{
		{
			name:    "Encrypt letter 'a' and decrypt it back",
			message: "a",
		},
		{
			name:    "Encrypt 'Hello, World!' and decrypt it back",
			message: "Hello, World!",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rsa := rsa.New()
			encrypted := rsa.EncryptString(tt.message)
			decrypted := rsa.DecryptString(encrypted)
			if decrypted != tt.message {
				t.Errorf("expected %s, got %s", tt.message, decrypted)
			}
		})
	}
}

func BenchmarkRSAEncryption(b *testing.B) {
	rsa := rsa.New()
	for i := 0; i < b.N; i++ {
		rsa.EncryptString("Hello, World!")
	}
}

func BenchmarkRSADecryption(b *testing.B) {
	rsa := rsa.New()
	encrypted := rsa.EncryptString("Hello, World!")
	for i := 0; i < b.N; i++ {
		rsa.DecryptString(encrypted)
	}
}
