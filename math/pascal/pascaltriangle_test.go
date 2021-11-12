// pascaltriangle_test.go
// description: Test for Pascal's triangle
// author(s) [red_byte](https://github.com/i-redbyte)
// see pascaltriangle.go

package pascal

import (
	"reflect"
	"testing"
)

func TestGenerateTriangle(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want [][]int
	}{
		{name: "Pascal's three-line triangle", n: 3, want: [][]int{{1}, {1, 1}, {1, 2, 1}}},
		{name: "Pascal's 0-line triangle", n: 0, want: [][]int{}},
		{name: "Pascal's one-line triangle", n: 1, want: [][]int{{1}}},
		{name: "Pascal's 7-line triangle", n: 7, want: [][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}, {1, 4, 6, 4, 1}, {1, 5, 10, 10, 5, 1}, {1, 6, 15, 20, 15, 6, 1}}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := GenerateTriangle(test.n); !reflect.DeepEqual(got, test.want) {
				t.Errorf("GenerateTriangle() = %v, want %v", got, test.want)
			}
		})
	}
}

func BenchmarkGenerateTriangle(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GenerateTriangle(10)
	}
}
