package sha256

import (
	"encoding/hex"
	"testing"
)

func TestHash(t *testing.T) {
	testCases := []struct {
		in       string
		expected string
	}{
		{"hello world", "b94d27b9934d3e08a52e52d7da7dabfac484efe37a5380ee9088f7ace2efcde9"},
		{"", "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855"},
		{"a", "ca978112ca1bbdcafac231b39a23dc4da786eff8147c4e72b9807785afee48bb"},
		{"The quick brown fox jumps over the lazy dog", "d7a8fbb307d7809469ca9abcb0082e4f8d5651e46d3cdb762d02d0bf37c9e592"},
		{"The quick brown fox jumps over the lazy dog.", "ef537f25c895bfa782526529a9b63d97aa631564d5d789c2b765448c8635fb6c"},
		{"Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.", "2d8c2f6d978ca21712b5f6de36c9d31fa8e96a4fa5d8ff8b0188dfb9e7c171bb"},
	}
	for _, tc := range testCases {
		res := Hash([]byte(tc.in))
		result := hex.EncodeToString(res[:])
		if result != tc.expected {
			t.Fatalf("Hash(%s) = %s, expected %s", tc.in, result, tc.expected)
		}
	}
}

func BenchmarkHash(b *testing.B) {
	testCases := []struct {
		name string
		in   string
	}{
		{"hello world", "hello world"},
		{"empty", ""},
		{"a", "a"},
		{"The quick brown fox jumps over the lazy dog", "The quick brown fox jumps over the lazy dog"},
		{"The quick brown fox jumps over the lazy dog.", "The quick brown fox jumps over the lazy dog."},
		{"Lorem ipsum", "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum."},
	}
	for _, testCase := range testCases {
		b.Run(testCase.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Hash([]byte(testCase.in))
			}
		})
	}
}
