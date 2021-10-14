package math

import (
	"testing"
)

func getTestsForPhi() []struct {
	n        int64
	expected int64
} {
	var tests = []struct {
		n        int64
		expected int64
	}{
		{4, 2},
		{5, 4},
		{7, 6},
		{10, 4},
		{999, 648},
		{1000, 400},
		{1000000, 400000},
		{999999, 466560},
		{999999999999878, 473684210526240},
	}
	return tests
}

func TestPhi(t *testing.T) {

	tests := getTestsForPhi()
	for _, test := range tests {
		result := Phi(test.n)
		t.Log(test.n, " ", result)
		if result != test.expected {
			t.Errorf("Wrong result! Expected:%v, returned:%v ", test.expected, result)
		}
	}
}

func BenchmarkPhi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Phi(65536)
	}
}
