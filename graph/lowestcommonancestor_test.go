package graph

import (
	"testing"
)

type Query struct {
	u        int
	v        int
	expected int
}

func TestLCA(t *testing.T) {
	testSuites := []struct {
		numbersVertex int
		root          int
		edges         []TreeEdge
		queries       []Query
	}{
		{
			numbersVertex: 2,
			root:          0,
			edges: []TreeEdge{
				{
					from: 0,
					to:   1,
				},
			},
			queries: []Query{
				{
					u:        0,
					v:        0,
					expected: 0,
				},
				{
					u:        0,
					v:        1,
					expected: 0,
				},
				{
					u:        1,
					v:        1,
					expected: 1,
				},
			},
		},
	}

	for idx, test := range testSuites {
		tree := NewTree(test.numbersVertex, test.root, test.edges)
		tree.BuildLCA()

		for _, query := range test.queries {
			actual := tree.GetLCA(query.u, query.v)
			expected := query.expected
			if actual != expected {
				t.Errorf("Test #%d: Expect %d, but actual %d", idx, expected, actual)
			}
		}
	}
}

func TestLCAWithLargeTree(t *testing.T) {

}
