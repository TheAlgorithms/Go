package dsu

import (
	"testing"
)

type query interface{}

type findRootQuery struct {
	v           int
	expectedRes int
}

type mergeQuery struct {
	v, u        int
	expectedRes bool
}

type areInTheSameSetQuery struct {
	v, u        int
	expectedRes bool
}

type testDsuInstance struct {
	dsu         DSU
	description string
}

func TestDSU(t *testing.T) {
	var dsuTests = []query{
		findRootQuery{
			v:           1,
			expectedRes: 1,
		},
		findRootQuery{
			v:           2,
			expectedRes: 2,
		},
		areInTheSameSetQuery{
			v:           1,
			u:           2,
			expectedRes: false,
		},
		mergeQuery{
			v:           1,
			u:           2,
			expectedRes: true,
		},
		areInTheSameSetQuery{
			v:           1,
			u:           2,
			expectedRes: true,
		},
		areInTheSameSetQuery{
			v:           1,
			u:           3,
			expectedRes: false,
		},
		mergeQuery{
			v:           2,
			u:           3,
			expectedRes: true,
		},
		areInTheSameSetQuery{
			v:           1,
			u:           3,
			expectedRes: true,
		},
		mergeQuery{
			v:           2,
			u:           3,
			expectedRes: false,
		},
		mergeQuery{
			v:           4,
			u:           5,
			expectedRes: true,
		},
		areInTheSameSetQuery{
			v:           2,
			u:           5,
			expectedRes: false,
		},
		mergeQuery{
			v:           3,
			u:           5,
			expectedRes: true,
		},
		mergeQuery{
			v:           4,
			u:           1,
			expectedRes: false,
		},
		areInTheSameSetQuery{
			v:           4,
			u:           2,
			expectedRes: true,
		},
	}

	normalDsu := NewDsu(6)
	dsuPathCompression := NewDsuPathCompression(6)
	dsuIterateSmallerSet := NewDsuIterateSmallerSet(6)

	dsuInstances := []testDsuInstance{{normalDsu, "normal dsu compare by size"},
		{dsuPathCompression, "dsu with path compression, compare by rank"},
		{dsuIterateSmallerSet, "dsu iterate on smaller set"},
	}

	for _, instance := range dsuInstances {
		t.Run(instance.description, func(t *testing.T) {
			for queryNumber, q := range dsuTests {
				switch i := q.(type) {
				case findRootQuery:
					res := instance.dsu.FindRoot(i.v)
					if res != i.expectedRes {
						t.Logf("FAIL: %s", instance.description)
						t.Fatalf("running query %d FindRoot(%d) Expected result: %d\nFound: %d\n", queryNumber, i.v, i.expectedRes, res)
					}
				case mergeQuery:
					res := instance.dsu.Merge(i.v, i.u)
					if res != i.expectedRes {
						t.Logf("FAIL: %s", instance.description)
						t.Fatalf("running query %d Merge(%d, %d) Expected result: %t\nFound: %t\n", queryNumber, i.v, i.u, i.expectedRes, res)
					}
				case areInTheSameSetQuery:
					res := instance.dsu.AreInTheSameSet(i.v, i.u)
					if res != i.expectedRes {
						t.Logf("FAIL: %s", instance.description)
						t.Fatalf("running query %d AreInTheSameSet(%d, %d) Expected result: %t\nFound: %t\n", queryNumber, i.v, i.u, i.expectedRes, res)
					}
				}
			}
		})
	}

}
