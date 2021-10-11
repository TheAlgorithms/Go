package segmenttree

import (
	"testing"
)

//Query interval
type query struct {
	firstIndex int
	lastIndex  int
}

func TestSegmentTree(t *testing.T) {
	var segmentTreeTestData = []struct {
		description string
		array       []int
		queries     []query
		expected    []int
	}{
		{
			description: "test empty array",
			array:       []int{},
			queries:     []query{{0, 0}},
			expected:    []int{0},
		},
		{
			description: "test array with size 5",
			array:       []int{1, 2, 3, 4, 5},
			queries:     []query{{0, 5}, {0, 2}, {2, 4}},
			expected:    []int{15, 6, 12},
		},
	}
	for _, test := range segmentTreeTestData {
		t.Run(test.description, func(t *testing.T) {
			segmentTree := NewSegmentTree(test.array)

			for i := 0; i < len(test.queries); i++ {
				result := segmentTree.Query(1, 0, len(test.array)-1, test.queries[i].firstIndex, test.queries[i].lastIndex)

				if result != test.expected[i] {
					t.Logf("FAIL: %s", test.description)
					t.Fatalf("Expected result: %d\nFound: %d\n", test.expected[i], result)
				}
			}

		})
	}
}
