// xorsearch_test.go
// description: Test for Find a missing number in a sequence
// author(s) [red_byte](https://github.com/i-redbyte)
// see xorsearch.go

package binary

import (
	"testing"
)

func TestXorSearchMissingNumber(t *testing.T) {
	var tests = []struct {
		name    string
		a       []int
		missing int
	}{
		{"[3, 0, 1]/2", []int{3, 0, 1}, 2},
		{"[0, 1, 3, 4, 5, 6, 7]/2", []int{0, 1, 3, 4, 5, 6, 7}, 2},
		{"[0, 2, 3, 4, 5, 6, 7, 8, 9]/1", []int{0, 2, 3, 4, 5, 6, 7, 8, 9}, 1},
		{"[0, 10, 9, 7, 2, 1,  4, 3, 5, 6]/8", []int{0, 10, 9, 7, 2, 1, 4, 3, 5, 6}, 8},
	}

	for _, tv := range tests {
		t.Run(tv.name, func(t *testing.T) {
			result := XorSearchMissingNumber(tv.a)
			if result != tv.missing {
				t.Errorf("Wrong result! Expected:%d, returned:%d ", tv.missing, result)
			}
		})
	}

}

func BenchmarkTestXorSearchMissingNumber(b *testing.B) {
	for i := 0; i < b.N; i++ {
		XorSearchMissingNumber([]int{0, 10, 9, 7, 2, 1, 4, 3, 5, 6})
	}
}
