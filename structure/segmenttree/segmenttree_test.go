package segmenttree

import (
	"testing"
)

// Query interval
type query struct {
	firstIndex int
	lastIndex  int
}

type update struct {
	firstIndex int
	lastIndex  int
	value      int
}

func TestSegmentTree(t *testing.T) {
	var segmentTreeTestData = []struct {
		description string
		array       []int
		updates     []update
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
		{
			description: "test array with size 5 and updates",
			array:       []int{1, 2, 3, 4, 5},
			updates: []update{{firstIndex: 1, lastIndex: 1, value: 2},
				{firstIndex: 2, lastIndex: 2, value: 3}},
			queries:  []query{{0, 5}, {0, 2}, {2, 4}},
			expected: []int{20, 11, 15},
		},
		{
			description: "test array with size 5 and single index updates",
			array:       []int{1, 2, 3, 4, 5},
			updates: []update{{firstIndex: 1, lastIndex: 1, value: 2},
				{firstIndex: 2, lastIndex: 2, value: 3}},
			queries:  []query{{0, 5}, {0, 2}, {2, 4}},
			expected: []int{20, 11, 15},
		},
		{
			description: "test array with size 5 and range updates",
			array:       []int{1, 2, 3, 4, 5},
			updates: []update{{firstIndex: 0, lastIndex: 4, value: 2},
				{firstIndex: 2, lastIndex: 4, value: 2}},
			queries:  []query{{0, 5}, {0, 2}, {2, 4}},
			expected: []int{31, 14, 24},
		},
		{
			description: "test array with size 11 and range updates",
			array:       []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			updates: []update{{firstIndex: 2, lastIndex: 8, value: 2},
				{firstIndex: 2, lastIndex: 8, value: 2}},
			queries:  []query{{3, 5}, {7, 8}, {4, 5}, {8, 8}},
			expected: []int{15, 10, 10, 5},
		},
	}
	for _, test := range segmentTreeTestData {
		t.Run(test.description, func(t *testing.T) {
			segmentTree := NewSegmentTree(test.array)

			for i := 0; i < len(test.updates); i++ {
				segmentTree.Update(1, 0, len(test.array)-1, test.updates[i].firstIndex, test.updates[i].lastIndex, test.updates[i].value)
			}

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
