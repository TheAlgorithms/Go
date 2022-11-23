// palindromicnumber_test.go
// author: Akshay Dubey (https://github.com/itsAkshayDubey)
// see palindromicnumber.go

package math_test

import (
	"testing"

	"github.com/TheAlgorithms/Go/math"
)

func TestPalindromicNumber(t *testing.T) {
	var tests = []struct {
		name          string
		n             int
		expectedValue bool
	}{
		{"-1 is not palindromic", -1, false},
		{"0 is palindromic", 0, true},
		{"23456 is not palindromic", 23456, false},
		{"110011 is palindromic", 11011, true},
		{"16461  is palindromic", 16461, true},
		{"1111111111111111111 is palindromic", 1111111111111111111, true},
		{"9223372036302733229  is palindromic", 9223372036302733229, true},
		{"9223372036854775807 is not palindromic", 9223372036854775807, false},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := math.PalindromicNumber(test.n)
			if result != test.expectedValue {
				t.Errorf("expected value: %v, got: %v", test.expectedValue, result)
			}
		})
	}
}
func BenchmarkPalindromicNumber(b *testing.B) {
	for i := 0; i < b.N; i++ {
		math.PalindromicNumber(65536)
	}
}
