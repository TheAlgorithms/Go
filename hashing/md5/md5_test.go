// md5_test.go
// description: Tests for the MD5 hashing function as defined in RFC 1321.
// author: Simon Waldherr

package md5

import (
	"encoding/hex"
	"testing"
)

// Helper function to convert hash output to hex string for comparison
func toHexString(hash [16]byte) string {
	return hex.EncodeToString(hash[:])
}

// Test vectors for MD5 (from RFC 1321 and other known sources)
var tests = []struct {
	input    string
	expected string
}{
	{"", "d41d8cd98f00b204e9800998ecf8427e"},
	{"a", "0cc175b9c0f1b6a831c399e269772661"},
	{"abc", "900150983cd24fb0d6963f7d28e17f72"},
	{"message digest", "f96b697d7cb7938d525a2f31aaf161d0"},
	{"abcdefghijklmnopqrstuvwxyz", "c3fcd3d76192e4007dfb496cca67e13b"},
	{"ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789", "d174ab98d277d9f5a5611c2c9f419d9f"},
	{"12345678901234567890123456789012345678901234567890123456789012345678901234567890", "57edf4a22be3c955ac49da2e2107b67a"},
}

// TestHash verifies that the Hash function produces the correct MD5 hash values
func TestHash(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result := Hash([]byte(tt.input))
			resultHex := toHexString(result)
			if resultHex != tt.expected {
				t.Errorf("MD5(%q) = %s; want %s", tt.input, resultHex, tt.expected)
			}
		})
	}
}
