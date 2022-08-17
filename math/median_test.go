// author(s) [jo3zeph](https://github.com/jo3zeph)
// median_test.go
// see median.go

package math_test

import (
	"testing"

	"github.com/TheAlgorithms/Go/math"
)

func TestMedian(t *testing.T) {
	testCases := []struct {
		name       string
		testValues []float64
		answer     float64
	}{

		{
			name:       "Series of numbers in ascending order",
			testValues: []float64{12, 14, 16, 18, 19},
			answer:     16,
		},

		{
			name:       "Series of numbers in random order",
			testValues: []float64{21, 10, 22, 33, 11, 88},
			answer:     21.5,
		},

		{
			name:       "Series of decimals in random order",
			testValues: []float64{11.2, 32.5, 2.5, 37.8, 21.8, 5.2},
			answer:     16.5,
		},

		{
			testValues: []float64{},
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {

			returnedMedian := math.Median(test.testValues)

			t.Log(test.testValues, " ", returnedMedian)

			if returnedMedian != test.answer {
				t.Errorf("Test failed. Median should have been %v but received %v",
					test.answer, returnedMedian)
			}
		})
	}
}
