// luhn_test.go
// description: Test for Luhn algorithm
// author(s) [red_byte](https://github.com/i-redbyte)
// see luhn.go

package checksum_test

import (
	"testing"

	"github.com/TheAlgorithms/Go/checksum"
)

func TestLuhn(t *testing.T) {
	tests := []struct {
		name string
		s    []byte
		want bool
	}{
		{"check 4242424242424242", []byte("4242424242424242"), true},
		{"check 6200000000000005", []byte("6200000000000005"), true},
		{"check 5534200028533164", []byte("5534200028533164"), true},
		{"check 36227206271667", []byte("36227206271667"), true},
		{"check 471629309440", []byte("471629309440"), false},
		{"check 1111", []byte("1111"), false},
		{"check 12345674", []byte("12345674"), true},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := checksum.Luhn(test.s); got != test.want {
				t.Errorf("LuhnAlgorithm() = %v, want %v", got, test.want)
			}
		})
	}
}

func BenchmarkBruteForceFactorial(b *testing.B) {
	for i := 0; i < b.N; i++ {
		checksum.Luhn([]byte("4242424242424242"))
	}
}
