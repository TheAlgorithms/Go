// luhn_test.go
// description: Test for Luhn algorithm
// author(s) [red_byte](https://github.com/i-redbyte)
// see luhn.go

package checksum

import "testing"

func TestLuhnAlgorithm(t *testing.T) {
	tests := []struct {
		name string
		s    []rune
		want bool
	}{
		{"check 4242424242424242", []rune("4242424242424242"), true},
		{"check 6200000000000005", []rune("6200000000000005"), true},
		{"check 5534200028533164", []rune("5534200028533164"), true},
		{"check 36227206271667", []rune("36227206271667"), true},
		{"check 471629309440", []rune("471629309440"), false},
		{"check 1111", []rune("1111"), false},
		{"check 12345674", []rune("12345674"), true},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := LuhnAlgorithm(test.s); got != test.want {
				t.Errorf("LuhnAlgorithm() = %v, want %v", got, test.want)
			}
		})
	}
}

func BenchmarkBruteForceFactorial(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LuhnAlgorithm([]rune("4242424242424242"))
	}
}
