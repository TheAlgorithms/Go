// sqrt_test.go
// description: Test for square root calculation
// author(s) [red_byte](https://github.com/i-redbyte)
// see sqrt.go

package binary

import (
	"math"
	"testing"
)

const epsilon = 0.001

func TestSquareRootCalculation(t *testing.T) {
	tests := []struct {
		name   string
		number float32
		want   float64
	}{
		{"sqrt(1)", 1, 1},
		{"sqrt(9)", 9, 3},
		{"sqrt(25)", 25, 5},
		{"sqrt(121)", 121, 11},
		{"sqrt(10000)", 10000, 100},
		{"sqrt(169)", 169, 13},
		{"sqrt(0)", 0, 0},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := Sqrt(test.number)
			delta := math.Abs(test.want - float64(got))
			if delta > epsilon {
				t.Errorf("Sqrt() = %v, want %v delta %v", got, test.want, delta)
			}
		})
	}
}

func BenchmarkSquareRootCalculation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sqrt(225)
	}
}

func BenchmarkMathSqrt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		math.Sqrt(225)
	}
}
