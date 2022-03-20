package math

import (
	"math"
	"testing"
)

const epsilonCos = 0.001

func TestCos(t *testing.T) {
	tests := []struct {
		name string
		n    float64
		want float64
	}{
		{"cos(0)", 0, 1},
		{"cos(90)", 90, -0.447},
		{"cos(180)", 180, -0.598},
		{"cos(1)", 1, 0.540},
		{"cos(π)", math.Pi, -1},
		{"cos(π/2)", math.Pi / 2, 0},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := Cos(test.n)
			if math.Abs(got-test.want) >= epsilonCos {
				t.Errorf("Cos() = %v, want %v", got, test.want)
				t.Errorf("MATH Cos() = %v", math.Cos(test.n))
			}
		})
	}
}

func BenchmarkCos(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Cos(180)
	}
}

// BenchmarkMathCos is slower because the standard library `math.Cos` calculates a more accurate value.
func BenchmarkMathCos(b *testing.B) {
	for i := 0; i < b.N; i++ {
		math.Cos(180)
	}
}
