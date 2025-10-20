
package math

import (
	"math"
	"testing"
)

func TestNeville(t *testing.T) {
	testCases := []struct {
		name     string
		x        []float64
		y        []float64
		target   float64
		expected float64
		hasError bool
	}{
		{
			name:     "Linear Interpolation",
			x:        []float64{0, 2},
			y:        []float64{1, 5},
			target:   1,
			expected: 3.0,
			hasError: false,
		},
		{
			name:     "Quadratic Interpolation",
			x:        []float64{0, 1, 3},
			y:        []float64{0, 1, 9},
			target:   2,
			expected: 4.0,
			hasError: false,
		},
		{
			name:     "Negative Numbers Interpolation",
			x:        []float64{-1, 0, 2},
			y:        []float64{4, 1, 1},
			target:   1,
			expected: 0.0,
			hasError: false,
		},
		{
			name:     "Mismatched Lengths",
			x:        []float64{1, 2},
			y:        []float64{1},
			target:   1.5,
			expected: 0,
			hasError: true,
		},
		{
			name:     "Empty Slices",
			x:        []float64{},
			y:        []float64{},
			target:   1,
			expected: 0,
			hasError: true,
		},
		{
			name:     "Duplicate X-Coordinates",
			x:        []float64{1, 2, 1},
			y:        []float64{5, 8, 3},
			target:   1.5,
			expected: 0,
			hasError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := Neville(tc.x, tc.y, tc.target)

			if tc.hasError {
				if err == nil {
					t.Errorf("Expected an error but got none")
				}
			} else {
				if err != nil {
					t.Errorf("Expected no error but got: %v", err)
				}
				if math.Abs(result-tc.expected) > 1e-9 {
					t.Errorf("Expected %f, but got %f", tc.expected, result)
				}
			}
		})
	}
}