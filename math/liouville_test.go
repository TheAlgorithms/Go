// liouville_test.go
// description: Returns λ(n)
// author: Akshay Dubey (https://github.com/itsAkshayDubey)
// see liouville.go

package math_test

import (
	"testing"

	"github.com/TheAlgorithms/Go/math"
)

func TestLiouvilleLambda(t *testing.T) {
	var tests = []struct {
		name          string
		n             int
		expectedValue int
		expectedError error
	}{
		{"n = 10", 10, 1, nil},
		{"n = 11", 11, -1, nil},
		{"n = -1", -1, 0, math.ErrPosArgsOnly},
		{"n = 0", 0, 0, math.ErrNonZeroArgsOnly},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := math.LiouvilleLambda(test.n)
			if result != test.expectedValue || test.expectedError != err {
				t.Errorf("expected error: %s, got: %s; expected value: %v, got: %v", test.expectedError, err, test.expectedValue, result)
			}
		})
	}
}
func BenchmarkLiouvilleLambda(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = math.LiouvilleLambda(65536)
	}
}
