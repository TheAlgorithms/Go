// sha1_test.go
// description: Test for SHA1 algorithm
// author(s) [red_byte](https://github.com/i-redbyte)
// see sha1.go

package checksum_test

import (
	"testing"

	"github.com/TheAlgorithms/Go/checksum"
)

func TestSHA1Checksum(t *testing.T) {
	tests := []struct {
		input        string
		expectedHash string
	}{
		{"Hello, SHA-1!", "f322e078fef4f49da1618d3793d3272a91f0488c"},
		{"", "da39a3ee5e6b4b0d3255bfef95601890afd80709"},
		{"OpenAI", "a19ee5a9fdc191092796cd9bfa97c55fe5631695"},
		{"Nikhil", "9e7f55eff8611cf542c232efcab1bb09f8037633"},
		{"😎", "5bdba8421c123fefa9325e58d73ef02f2da7cd51"},
		{"@As%.{-+*`~NNNnks!A*}", "52a31dae50da392c801d1551b1e38c46e38fb214"},
		{"0x25", "ee0694e25ddbe98f179fd000b09c7be79790cec0"},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			hash := checksum.Sha1(test.input)
			if hash != test.expectedHash {
				t.Errorf("Expected hash %s, but got %s", test.expectedHash, hash)
			}
		})
	}
}

func BenchmarkSHA1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		checksum.Sha1("THe @lGoR|ThM$ - GO")
	}
}
