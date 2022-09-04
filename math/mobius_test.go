// mobius_test.go
// description: Returns μ(n)
// author: Akshay Dubey (https://github.com/itsAkshayDubey)
// see mobius.go

package math

import (
	"testing"
)

func getTestsForMu() []struct {
	n        int
	expected int
} {
	var tests = []struct {
		n        int
		expected int
	}{
		{-1, 1},
		{0, 1},
		{2, -1},
		{3, -1},
		{95, 1},
		{97, -1},
		{98, 0},
		{99, 0},
		{100, 0},
	}
	return tests
}

func TestMu(t *testing.T) {

	tests := getTestsForMu()
	for _, test := range tests {
		result := Mu(test.n)
		t.Log(test.n, " ", result)
		if result != test.expected {
			t.Errorf("Wrong result! Expected:%v, returned:%v ", test.expected, result)
		}
	}
}

func BenchmarkMu(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Mu(65536)
	}
}
