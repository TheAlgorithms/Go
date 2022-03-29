package math_test

import (
	algmath "github.com/TheAlgorithms/Go/math"
	stdmath "math"
	"testing"
)

func TestSin(t *testing.T) {
	tests := []struct {
		name string
		n    float64
		want float64
	}{
		{"sin(0)", 0, 0},
		{"sin(3π/2)", (3 * stdmath.Pi) / 2, -1},
		{"sin(π/2)", stdmath.Pi / 2, 1},
		{"sin(π/6)", stdmath.Pi / 6, 0.5},
		{"sin(90)", 90, 0.893996663600558},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := algmath.Sin(test.n)
			if stdmath.Abs(got-test.want) >= epsilon {
				t.Errorf("Sin() = %v, want %v", got, test.want)
				t.Errorf("MATH Sin() = %v", stdmath.Sin(test.n))
			}
		})
	}
}

func BenchmarkSin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		algmath.Sin(180)
	}
}

// BenchmarkMathSin is slower because the standard library `math.Sin` calculates a more accurate value.
func BenchmarkMathSin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		stdmath.Sin(180)
	}
}
