package prime

import (
	"reflect"
	"testing"
)

func TestFactorize(t *testing.T) {
	var tests = []struct {
		n        int64
		expected map[int64]int64
	}{
		{4, map[int64]int64{2: 2}},
		{5, map[int64]int64{5: 1}},
		{7, map[int64]int64{7: 1}},
		{10, map[int64]int64{2: 1, 5: 1}},
		{999, map[int64]int64{3: 3, 37: 1}},
		{999999999999878, map[int64]int64{2: 1, 19: 1, 26315789473681: 1}},
	}
	for _, test := range tests {
		result := Factorize(test.n)
		t.Log(test.n, " ", result)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("Wrong result! Expected:%v, returned:%v ", test.expected, result)
		}
	}
}

func BenchmarkFactorize(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Factorize(1000000007)
	}
}
