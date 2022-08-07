package math_test

import (
	"github.com/TheAlgorithms/Go/math"
	"testing"
)

func TestMean(t *testing.T) {
	testCases := []struct {
		name       string
		testValues []float64
		average    float64
	}{
		{
			name:       "All 0s",
			testValues: []float64{0, 0, 0, 0, 0},
			average:    0,
		},
		{
			name:       "With integer values",
			testValues: []float64{1, 2, 3, 4, 5},
			average:    3.0,
		},
		{
			name:       "With negative values",
			testValues: []float64{-1, 2, -3, 4, 5},
			average:    1.4,
		},
		{
			name:       "With floating values",
			testValues: []float64{1.1, 2.2, 3.3, 4.4, 5.5},
			average:    3.3,
		},
		{
			name:       "With no values",
			testValues: []float64{},
			average:    0,
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			returnedAverage := math.Mean(test.testValues)
			if returnedAverage != test.average {
				t.Errorf("\nFailed test: %s\ntestValues: %v\naverage: %v\nbut received: %v\n",
					test.name, test.testValues, test.average, returnedAverage)
			}

		})
	}

}
