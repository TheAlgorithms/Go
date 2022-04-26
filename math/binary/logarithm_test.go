// logarithm_test.go
// description: Test for finding the exponent of n = 2**x using bitwise operations (logarithm in base 2 of n)
// author(s) [red_byte](https://github.com/i-redbyte)
// see logarithm.go

package binary

import (
	"math"
	"testing"
)

func TestLogBase2(t *testing.T) {
	tests := []struct {
		name string
		n    uint32
		want uint32
	}{
		{"log2(1) = 0", 1, 0},
		{"log2(2) = 1", 2, 1},
		{"log2(4) = 2", 4, 2},
		{"log2(8) = 3", 8, 3},
		{"log2(16) = 4", 16, 4},
		{"log2(32) = 5", 32, 5},
		{"log2(64) = 6", 64, 6},
		{"log2(128) = 7", 128, 7},
		{"log2(256) = 8", 256, 8},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := LogBase2(test.n); got != test.want {
				t.Errorf("LogBase2() = %v, want %v", got, test.want)
			}
		})
	}
}

func BenchmarkBitwiseLogBase2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LogBase2(1024)
	}
}

func BenchmarkMathPAckageLogBase2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		math.Log2(1024)
	}
}
