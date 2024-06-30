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
			name:    "Encrypt letter 'a'",
			message: "a",
		},
		{
			name:    "Encrypt 'Hello, World!'",
			message: "Hello, World!",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rsa := rsa.InitRSA()
			encrypted := rsa.EncryptString(tt.message)
			decrypted := rsa.DecryptString(encrypted)
			if decrypted != tt.message {
				t.Errorf("expected %s, got %s", tt.message, decrypted)
			}
		})
	}
}
