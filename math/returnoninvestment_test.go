// returnoninvestment_test.go
// description: Tests for Return on Investment (ROI)
// author(s) [InukaWijerathna](https://github.com/InukaWijerathna)
// see returnoninvestment.go

package math_test

import (
	"testing"

	"github.com/TheAlgorithms/Go/math"
)

func TestReturnOnInvestment(t *testing.T) {
	var tests = []struct {
		name          string
		gain          float64
		cost          float64
		expectedValue float64
		expectedError error
	}{
		{"positive ROI", 1000, 500, 100, nil},
		{"zero ROI", 500, 500, 0, nil},
		{"negative ROI", 200, 500, -60, nil},
		{"total loss", 0, 500, -100, nil},
		{"zero cost", 1000, 0, 0, math.ErrNonPositiveCost},
		{"negative cost", 1000, -100, 0, math.ErrNonPositiveCost},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := math.ReturnOnInvestment(test.gain, test.cost)
			if result != test.expectedValue || test.expectedError != err {
				t.Errorf("expected error: %v, got: %v; expected value: %v, got: %v",
					test.expectedError, err, test.expectedValue, result)
			}
		})
	}
}

func BenchmarkReturnOnInvestment(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = math.ReturnOnInvestment(1000, 500)
	}
}
