// binomialcoefficient_test.go
// description: Returns C(n, k)
// author: Akshay Dubey (https://github.com/itsAkshayDubey)
// see binomialcoefficient.go

package math_test

import (
	"errors"
	"strings"
	"testing"

	"github.com/TheAlgorithms/Go/math"
)

func TestCombinations(t *testing.T) {
	var tests = []struct {
		n        int
		k        int
		expected int
	}{
		{5, 2, 10},
		{7, 4, 35},
		{0, 0, 1},
		{-1, 1, -1},
		{1, -1, -1},
		{-1, -1, -1},
	}
	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			var expectedError error = errors.New("arguments must be positive")
			result, error := math.Combinations(test.n, test.k)
			if result != test.expected {
				t.Errorf("Wrong result! Expected:%v, Returned:%v", test.expected, result)
			}
			if result == -1 && error == nil {
				t.Errorf("Wrong result! Expected:%v, Returned:%v", expectedError, error.Error())
			}
			if result != -1 && error != nil {
				t.Errorf("Wrong result! Expected:%v, Returned:%v", nil, error.Error())
			}
			if result == -1 && !strings.EqualFold(expectedError.Error(), error.Error()) {
				t.Errorf("Wrong result! Expected:%v, Returned:%v", expectedError.Error(), error.Error())
			}
		})
	}
}
func BenchmarkCombinations(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = math.Combinations(65536, 65536)
	}
}
