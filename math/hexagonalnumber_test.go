// hexagonalnumber_test.go
// description: Returns nth hexagonal number
// author: Akshay Dubey (https://github.com/itsAkshayDubey)
// see hexagonalnumber.go

package math_test

import (
	"testing"

	"github.com/TheAlgorithms/Go/math"
)

func TestHexagonalNumber(t *testing.T) {
	var tests = []struct {
		name          string
		n             int
		expectedValue int
		expectedError error
	}{
		{"n = 7", 7, 91, nil},
		{"n = 22", 22, 946, nil},
		{"n = -1", -1, 0, math.ErrPosArgsOnly},
		{"n = 0", 0, 0, math.ErrNonZeroArgsOnly},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := math.HexagonalNumber(test.n)
			if result != test.expectedValue || test.expectedError != err {
				t.Errorf("expected error: %s, got: %s; expected value: %v, got: %v", test.expectedError, err, test.expectedValue, result)
			}
		})
	}
}
func BenchmarkHexagonalNumber(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = math.HexagonalNumber(65536)
	}
}
