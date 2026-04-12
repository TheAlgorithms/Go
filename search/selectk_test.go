package search

import (
	"testing"
)

func TestSelectK(t *testing.T) {
	tests := []struct {
		data     []int
		k        int
		expected int
		err      error
		name     string
	}{
		{[]int{1, 2, 3, 4, 5}, 1, 5, nil, "sorted data"},
		{[]int{5, 4, 3, 2, 1}, 2, 4, nil, "reversed data"},
		{[]int{3, 1, 2, 5, 4}, 3, 3, nil, "random data"},
		{[]int{3, 2, 1, 5, 4}, 10, -1, ErrNotFound, " absent data"},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			elem, err := SelectK(tc.data, tc.k)
			if err != tc.err {
				t.Errorf("name:%v SelectK() = _, %v, want err: %v", tc.name, err, tc.err)
			}
			if elem != tc.expected {
				t.Errorf("name:%v SelectK() = %v,_ , want: %v", tc.name, elem, tc.expected)
			}
		})
	}
}

func BenchmarkSelectK(b *testing.B) {
	testCase := generateBenchmarkTestCase()
	testCase = append(testCase, 1) // make sure len(testCase)/2+1 is valid
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, _ = SelectK(testCase, len(testCase)/2)
	}
}
