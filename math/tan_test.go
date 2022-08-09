package math_test

import (
	algmath "github.com/TheAlgorithms/Go/math"
	stdmath "math"
	"testing"
)

const epsilon = 0.001

func TestTan(t *testing.T) {
	tests := []struct {
		name string
		n    float64
		want float64
	}{
		{"tan(0)", 0, 0},
		{"tan(π/4)", stdmath.Pi / 4, 1},
		{"tan(1)", 1, 0.540},
		{"tan(π)", stdmath.Pi, 0},
		{"tan(2π)", (2*stdmath.Pi), 0},
		{"tan(π/6)", stdmath.Pi / 6, 0.57735026919}
		{"tan(270", 270, -0.17883906379}
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := algmath.Cos(test.n)
			if stdmath.Abs(got-test.want) >= epsilon {
				t.Errorf("Tan() = %v, want %v", got, test.want)
				t.Errorf("MATH Tan() = %v", stdmath.Tan(test.n))
			}
		})
	}
}

func BenchmarkTan(b *testing.B) {
	for i := 0; i < b.N; i++ {
		algmath.Tan(180)
	}
}

// BenchmarkMathCos is slower because the standard library `math.Tan` calculates a more accurate value.
func BenchmarkMathTan(b *testing.B) {
	for i := 0; i < b.N; i++ {
		stdmath.Tan(180)
	}
}