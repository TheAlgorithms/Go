package graph

import (
	"math/rand"
	"testing"
	"time"
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
		}, {
			numbersVertex: 1,
			root:          0,
			edges:         []TreeEdge{},
			queries: []Query{
				{
					u:        0,
					v:        0,
					expected: 0,
				},
			},
		}, {
			numbersVertex: 9,
			root:          7,
			edges: []TreeEdge{
				{
					from: 0,
					to:   1,
				},
				{
					from: 0,
					to:   2,
				},
				{
					from: 2,
					to:   3,
				},
				{
					from: 2,
					to:   4,
				},
				{
					from: 2,
					to:   6,
				},
				{
					from: 3,
					to:   5,
				},
				{
					from: 6,
					to:   7,
				},
				{
					from: 6,
					to:   8,
				},
			},
			queries: []Query{
				{
					u:        1,
					v:        0,
					expected: 0,
				},
				{
					u:        8,
					v:        2,
					expected: 6,
				},
				{
					u:        1,
					v:        5,
					expected: 2,
				},
				{
					u:        4,
					v:        7,
					expected: 7,
				},
				{
					u:        0,
					v:        8,
					expected: 6,
				},
			},
		},
	}
	// Test #2:
	//							7
	//						 /
	//						6
	//					/  \
	//				2    8
	//			/ | \
	//		 0  3  4
	//		/   |
	//	 1    5
	for idx, test := range testSuites {
		tree := NewTree(test.numbersVertex, test.root, test.edges)
		LowestCommonAncestor(tree)

		for qi, query := range test.queries {
			actual := tree.GetLCA(query.u, query.v)
			expected := query.expected
			if actual != expected {
				t.Errorf("\nTest #%d:\nQuery #%d: u = %d, v = %d\nExpected %d, but actual %d", idx, qi, query.u, query.v, expected, actual)
			}
		}
	}
}

func generateTree() *Tree {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))

	const MAXVERTEX int = 2000
	var numbersVertex int = rnd.Intn(MAXVERTEX) + 1
	var root int = rnd.Intn(numbersVertex)
	var edges []TreeEdge

	var fullGraph []TreeEdge
	for u := 0; u < numbersVertex; u++ {
		for v := 0; v < numbersVertex; v++ {
			fullGraph = append(fullGraph, TreeEdge{
				from: u,
				to:   v,
			})
		}
	}
	rnd.Shuffle(len(fullGraph), func(i, j int) {
		fullGraph[i], fullGraph[j] = fullGraph[j], fullGraph[i]
	})

	par := make([]int, numbersVertex)
	for u := 0; u < numbersVertex; u++ {
		par[u] = u
	}

	var findp func(int) int
	findp = func(u int) int {
		if u == par[u] {
			return u
		} else {
			par[u] = findp(par[u])
			return par[u]
		}
	}

	join := func(u, v int) bool {
		u, v = findp(u), findp(v)
		if u == v {
			return false
		}
		par[v] = u
		return true
	}

	for _, e := range fullGraph {
		if join(e.from, e.to) == true {
			edges = append(edges, e)
		}
	}

	return NewTree(numbersVertex, root, edges)
}

func generateQuery(tree *Tree) []Query {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	const MAXQUERY = 50
	var queries []Query

	bruteforceLCA := func(u, v int) int {
		for u != v {
			if tree.GetDepth(u) > tree.GetDepth(v) {
				u = tree.GetDad(u)
			} else {
				v = tree.GetDad(v)
			}
		}
		return u
	}

	for q := 1; q <= MAXQUERY; q++ {
		u := rnd.Intn(tree.numbersVertex)
		v := rnd.Intn(tree.numbersVertex)
		queries = append(queries, Query{
			u:        u,
			v:        v,
			expected: bruteforceLCA(u, v),
		})
	}

	return queries
}

// Test with the tree with up to 2000 vertices.
func TestLCAWithLargeTree(t *testing.T) {
	const MAXTEST int = 20

	for test := 1; test <= MAXTEST; test++ {
		tree := generateTree()
		LowestCommonAncestor(tree)

		queries := generateQuery(tree)

		for qi, query := range queries {
			actual := tree.GetLCA(query.u, query.v)
			expected := query.expected
			if actual != expected {
				t.Errorf("\nTest #%d:\nQuery #%d: u = %d, v = %d\nExpected %d, but actual %d", test, qi, query.u, query.v, expected, actual)
			}
		}
	}
}
