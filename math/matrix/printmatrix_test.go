package matrix_test

import (
	"testing"

	"github.com/TheAlgorithms/Go/math/matrix"
)

func TestPrintMatrix(t *testing.T) {
	type args struct {
		matrix interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "1", args: args{matrix: [][]string{
			{"apple", "banana", "cherry"},
			{"dog", "elephant", "fox"},
		}},
		}, {name: "2", args: args{matrix: [][]int{
			{1, 2, 3},
			{4, 5, 6},
		}},
		}, {name: "3", args: args{matrix: [][]float32{
			{12.3, 12.2},
		}},
		}, {name: "4", args: args{matrix: [][]bool{
			{true, false},
			{false, false},
			{true, false},
		}},
		}, {name: "5", args: args{matrix: [][]int16{}}}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			matrix.PrintMatrix(tt.args.matrix)
		})
	}
}
