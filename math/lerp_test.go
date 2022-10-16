package math_test

import (
	"testing"

	algmath "github.com/TheAlgorithms/Go/math"
)

func TestLerp(t *testing.T) {
	tests := []struct {
		name       string
		testValues []float64
		answer     float64
	}{
		{"Lerp(1,1,1)", []float64{1, 1, 1}, 1},
		{"Lerp(0,1,1)", []float64{0, 1, 1}, 1},
		{"Lerp(0,1,0.5)", []float64{0, 1, 0.5}, 0.5},
		{"Lerp(0,1,0.1)", []float64{0, 1, 0.1}, 0.1},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := algmath.Lerp(test.testValues[0], test.testValues[1], test.testValues[2])
			if got != test.answer {
				t.Errorf("Lerp(%f,%f,%f) = %v, want %v", got, test.testValues[0],
					test.testValues[1], test.testValues[2], test.answer)
			}
		})
	}
}
