package graph

import (
    "fmt"
	"testing"
)

var testCases = []struct {
    name string
    N int
	constraints [][]int
}{
	{
        "normal test", 2,
		[][]int{[]int{1,0}},
	},
}

func TestTopoSort(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := TopoSort(tc.N, tc.constraints)

            visited := make([]bool, tc.N)
            positions := make([]int, tc.N)
            for i:=0;i<tc.N;i++ {
                positions[actual[i]] = i
                visited[actual[i]] = true
            }
            for i:=0;i<tc.N;i++ {
                if !visited[i] {
                    t.Errorf("nodes not all visited, %v", visited)
                }
            }
            for _, c := range tc.constraints {
                if positions[c[0]]>positions[c[1]] {
                    t.Errorf("%v dun satisfy %v", actual, c)
                }
            }
            fmt.Println(actual)
		})
	}
}
