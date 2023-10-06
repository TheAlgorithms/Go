package disjointsets

import "testing"

func TestDisjointSets(t *testing.T) {
	dsu := NewDisjointSet(5)

	// Initially, all elements are in their own sets
	for i := 0; i < 5; i++ {
		if dsu.Size(i) != 1 {
			t.Errorf("Size of set containing element %d should be 1", i)
		}
	}

	// Union elements 0 and 1
	dsu.Union(0, 1)
	if !dsu.IsInSameSet(0, 1) {
		t.Error("Elements 0 and 1 should be in the same set after union")
	}

	// Size of the set containing 0 or 1 should now be 2
	if dsu.Size(0) != 2 || dsu.Size(1) != 2 {
		t.Error("Size of the set containing elements 0 and 1 should be 2")
	}

	// Union elements 2 and 3
	dsu.Union(2, 3)
	if !dsu.IsInSameSet(2, 3) {
		t.Error("Elements 2 and 3 should be in the same set after union")
	}

	// Size of the set containing 2 or 3 should now be 2
	if dsu.Size(2) != 2 || dsu.Size(3) != 2 {
		t.Error("Size of the set containing elements 2 and 3 should be 2")
	}

	// Elements 0, 1 and 2, 3 should still be in different sets
	if dsu.IsInSameSet(0, 2) || dsu.IsInSameSet(1, 3) {
		t.Error("Elements in different sets should not be in the same set")
	}

	// Union sets containing elements 0, 1 and 2, 3
	dsu.Union(0, 2)
	if !dsu.IsInSameSet(0, 3) || !dsu.IsInSameSet(1, 2) {
		t.Error("Elements 0, 1, and 2, 3 should be in the same set after union")
	}

	// Size of the set containing 0, 1, 2, 3 should now be 4
	if dsu.Size(0) != 4 || dsu.Size(1) != 4 || dsu.Size(2) != 4 || dsu.Size(3) != 4 {
		t.Error("Size of the set containing elements 0, 1, 2, 3 should be 4")
	}
}
