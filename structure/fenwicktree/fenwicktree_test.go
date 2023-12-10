package fenwicktree_test

import (
	"github.com/TheAlgorithms/Go/structure/fenwicktree"
	"testing"
)

type query struct {
	queryType  string
	firstIndex int // firstIndex and lastIndex are same for point queries
	lastIndex  int
}

type update struct {
	pos   int
	value int
}

func TestFenwickTree(t *testing.T) {
	var fenwickTreeTestData = []struct {
		description string
		array       []int
		updates     []update
		queries     []query
		expected    []int
	}{
		{
			description: "test empty array",
			array:       []int{},
			queries:     []query{{"point", 1, 1}},
			expected:    []int{0},
		},
		{
			description: "test array with size 5",
			array:       []int{1, 2, 3, 4, 5},
			queries:     []query{{"range", 1, 5}, {"range", 1, 3}, {"range", 3, 5}},
			expected:    []int{15, 6, 12},
		},
		{
			description: "test array with size 5, single index updates and range queries",
			array:       []int{1, 2, 3, 4, 5},
			updates:     []update{{pos: 2, value: 2}, {pos: 3, value: 3}},
			queries:     []query{{"range", 1, 5}, {"range", 1, 3}, {"range", 3, 5}},
			expected:    []int{20, 11, 15},
		},
		{
			description: "test array with size 5, single index updates and point queries",
			array:       []int{1, 2, 3, 4, 5},
			updates:     []update{{pos: 2, value: 2}, {pos: 3, value: 3}},
			queries:     []query{{"point", 3, 3}, {"point", 1, 1}, {"point", 5, 5}},
			expected:    []int{11, 1, 20},
		},
	}

	for _, test := range fenwickTreeTestData {
		t.Run(test.description, func(t *testing.T) {
			fenwickTree := fenwicktree.NewFenwickTree(test.array)

			for i := 0; i < len(test.updates); i++ {
				fenwickTree.Add(test.updates[i].pos, test.updates[i].value)
			}

			for i := 0; i < len(test.queries); i++ {

				var result int

				if test.queries[i].queryType == "point" {
					result = fenwickTree.PrefixSum(test.queries[i].firstIndex)
				} else {
					result = fenwickTree.RangeSum(test.queries[i].firstIndex, test.queries[i].lastIndex)
				}

				if result != test.expected[i] {
					t.Logf("FAIL: %s", test.description)
					t.Fatalf("Expected result: %d\nFound: %d\n", test.expected[i], result)
				}
			}
		})
	}

}
