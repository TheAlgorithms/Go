package sorts

import "testing"

//BEGIN TESTS
func TestMerge(t *testing.T) {
	for _, test := range sortTests {
		actual := Mergesort(test.input)
		pos, sorted := compareSlices(actual, test.expected)
		if !sorted {
			if pos == -1 {
				t.Errorf("test %s failed due to slice length changing", test.name)
			}
			t.Errorf("test %s failed at index %d", test.name, pos)
		}
	}
}

//END TESTS

//BEGIN BENCHMARKS
func BenchmarkMerge(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range sortTests {
			Mergesort(test.input)
		}
	}
}

//END BENCHMARKS

func compareSlices(a []int, b []int) (int, bool) {
	if len(a) != len(b) {
		return -1, false
	}
	for pos := range a {
		if a[pos] != b[pos] {
			return pos, false
		}
	}
	return -1, true
}
