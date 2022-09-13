// binomialcoefficient_test.go
// description: Returns C(n, k)
// author: Akshay Dubey (https://github.com/itsAkshayDubey)
// see binomialcoefficient.go

package math_test

import (
	"testing"

	"github.com/TheAlgorithms/Go/math"
)

func TestC(t *testing.T) {
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
		result, error := math.C(test.n, test.k)
		t.Log(test.n, " ", result, " ", error)
		if result != test.expected {
			t.Errorf("Wrong result! Expected:%v, Returned:%v, Error:%v", test.expected, result, error.Error())
		}
	}
}
func BenchmarkBinomialCoefficient(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = math.C(65536, 65536)
	}
}
