// sha1_test.go
// description: Tests for the SHA-1 hashing function as defined in RFC 3174.
// author: Simon Waldherr

package sha1

import (
	"encoding/hex"
	"testing"
)

// Helper function to convert hash output to hex string for comparison
func toHexString(hash [20]byte) string {
	return hex.EncodeToString(hash[:])
}

// Test vectors for SHA-1 (from RFC 3174 and other known sources)
var tests = []struct {
	input    string
	expected string
}{
	{"", "da39a3ee5e6b4b0d3255bfef95601890afd80709"},
	{"a", "86f7e437faa5a7fce15d1ddcb9eaeaea377667b8"},
	{"abc", "a9993e364706816aba3e25717850c26c9cd0d89d"},
	{"message digest", "c12252ceda8be8994d5fa0290a47231c1d16aae3"},
	{"abcdefghijklmnopqrstuvwxyz", "32d10c7b8cf96570ca04ce37f2a19d84240d3a89"},
	{"ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789", "761c457bf73b14d27e9e9265c46f4b4dda11f940"},
	{"1234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890", "fecfd28bbc9345891a66d7c1b8ff46e60192d284"},
}

// TestHash verifies that the Hash function produces the correct SHA-1 hash values
func TestHash(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result := Hash([]byte(tt.input))
			resultHex := toHexString(result)
			if resultHex != tt.expected {
				t.Errorf("SHA-1(%q) = %s; want %s", tt.input, resultHex, tt.expected)
			}
		})
	}
}
