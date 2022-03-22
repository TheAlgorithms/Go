package math

import (
	"math"
	"testing"
)

func TestSin(t *testing.T) {
	tests := []struct {
		name string
		n    float64
		want float64
	}{
		{"sin(0)", 0, 0},
		{"sin(3π/2)", (3 * math.Pi) / 2, -1},
		{"sin(π/2)", math.Pi / 2, 1},
		{"sin(π/6)", math.Pi / 6, 0.5},
		{"sin(90)", 90, 0.893996663600558},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := Sin(test.n)
			if math.Abs(got-test.want) >= epsilon {
				t.Errorf("Sin() = %v, want %v", got, test.want)
				t.Errorf("MATH Sin() = %v", math.Sin(test.n))
			}
		})
	}
}

func BenchmarkSin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sin(180)
	}
}

// BenchmarkMathSin is slower because the standard library `math.Sin` calculates a more accurate value.
func BenchmarkMathSin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		math.Sin(180)
	}
}
