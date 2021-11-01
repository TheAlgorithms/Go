// package set implements a Set using a golang map.
// This implies that only the types that are accepted as valid map keys can be used as set elements.
// For instance, do not try to Add a slice, or the program will panic.
package set

import (
	"testing"
)

func TestNew(t *testing.T) {
	set := New(1, 2, 3)
	if set.Len() != 3 {
		t.Errorf("expecting 3 elements in the set but got %v: %v", set.Len(), set.GetItems())
	}
	for _, n := range []int{1, 2, 3} {
		if !set.In(n) {
			t.Errorf("expecting %d to be present in the set but was not; set is %v", n, set.GetItems())
		}
	}
	if set.In(5) {
		t.Errorf("expecting 5 not to be present in the set but it was; set is %v", set.GetItems())
	}
}

func TestAdd(t *testing.T) {
	td := []struct {
		name     string
		input    int
		expElems []int
	}{
		{"add new element", 5, []int{1, 2, 3, 5}},
		{"add exiting element", 3, []int{1, 2, 3}},
	}
	for _, tc := range td {
		t.Run(tc.name, func(t *testing.T) {
			s := New(1, 2, 3)
			s.Add(tc.input)
			if s.Len() != len(tc.expElems) {
				t.Errorf("expecting %d elements in the set but got %d: set is %v", len(tc.expElems), s.Len(), s.GetItems())
			}
			for _, n := range tc.expElems {
				if !s.In(n) {
					t.Errorf("expecting %d to be present in the set but was not; set is %v", n, s.GetItems())
				}
			}
		})
	}
}

func TestDelete(t *testing.T) {
	td := []struct {
		name     string
		input    int
		expElems []int
	}{
		{"delete exiting element", 3, []int{1, 2}},
		{"delete not existing element", 5, []int{1, 2, 3}},
	}
	for _, tc := range td {
		t.Run(tc.name, func(t *testing.T) {
			s := New(1, 2, 3)
			s.Delete(tc.input)
			if s.Len() != len(tc.expElems) {
				t.Errorf("expecting %d elements in the set but got %d: set is %v", len(tc.expElems), s.Len(), s.GetItems())
			}
			for _, n := range tc.expElems {
				if !s.In(n) {
					t.Errorf("expecting %d to be present in the set but was not; set is %v", n, s.GetItems())
				}
			}
		})
	}
}

func TestIsSubsetOf(t *testing.T) {
	s1, s2 := New(1, 2, 3), New(1, 2, 3, 4)
	if !s1.IsSubsetOf(s2) {
		t.Errorf("expecting %v to be a subset of %v", s1, s2)
	}
	if s2.IsSubsetOf(s1) {
		t.Errorf("expecting %v not to be a subset of %v", s2, s1)
	}
	if s3 := New(1, 2, 5, 6); s1.IsSubsetOf(s3) {
		t.Errorf("expecting %v not to be a subset of %v", s1, s3)
	}
}

func TestIsSupersetOf(t *testing.T) {
	s1, s2 := New(1, 2, 3), New(1, 2, 3, 4)
	if !s2.IsSupersetOf(s1) {
		t.Errorf("expecting %v to be a superset of %v", s2, s1)
	}
	if s1.IsSupersetOf(s2) {
		t.Errorf("expecting %v not to be a superset of %v", s1, s2)
	}
	if s3 := New(1, 2, 5); s2.IsSupersetOf(s3) {
		t.Errorf("expecting %v not to be a superset of %v", s2, s3)
	}
}

func TestUnion(t *testing.T) {
	td := []struct {
		name   string
		s1     Set
		s2     Set
		expSet Set
	}{
		{"union of different sets", New(1, 2, 3), New(4, 5, 6), New(1, 2, 3, 4, 5, 6)},
		{"union of sets with elements in common", New(1, 2, 3), New(1, 2, 4), New(1, 2, 3, 4)},
		{"union of same sets", New(1, 2, 3), New(1, 2, 3), New(1, 2, 3)},
	}
	for _, tc := range td {
		t.Run(tc.name, func(t *testing.T) {
			s := tc.s1.Union(tc.s2)
			if s.Len() != tc.expSet.Len() {
				t.Errorf("expecting %d elements in the set but got %d: set is %v", tc.expSet.Len(), s.Len(), s.GetItems())
			}
			for _, n := range tc.expSet.GetItems() {
				if !s.In(n) {
					t.Errorf("expecting %d to be present in the set but was not; set is %v", n, s.GetItems())
				}
			}
		})
	}
}

func TestIntersection(t *testing.T) {
	td := []struct {
		name   string
		s1     Set
		s2     Set
		expSet Set
	}{
		{"intersection of different sets", New(0, 1, 2, 3), New(4, 5, 6), New()},
		{"intersection of sets with elements in common", New(1, 2, 3), New(1, 2, 4), New(1, 2)},
		{"intersection of same sets", New(1, 2, 3), New(1, 2, 3), New(1, 2, 3)},
	}
	for _, tc := range td {
		t.Run(tc.name, func(t *testing.T) {
			s := tc.s1.Intersection(tc.s2)
			if s.Len() != tc.expSet.Len() {
				t.Errorf("expecting %d elements in the set but got %d: set is %v", tc.expSet.Len(), s.Len(), s.GetItems())
			}
			for _, n := range tc.expSet.GetItems() {
				if !s.In(n) {
					t.Errorf("expecting %d to be present in the set but was not; set is %v", n, s.GetItems())
				}
			}
		})
	}
}

func TestDifference(t *testing.T) {
	td := []struct {
		name   string
		s1     Set
		s2     Set
		expSet Set
	}{
		{"difference of different sets", New(1, 2, 3), New(4, 5, 6), New(1, 2, 3)},
		{"difference of sets with elements in common", New(1, 2, 3), New(1, 2, 4), New(3)},
		{"difference of same sets", New(1, 2, 3), New(1, 2, 3), New()},
	}
	for _, tc := range td {
		t.Run(tc.name, func(t *testing.T) {
			s := tc.s1.Difference(tc.s2)
			if s.Len() != tc.expSet.Len() {
				t.Errorf("expecting %d elements in the set but got %d: set is %v", tc.expSet.Len(), s.Len(), s.GetItems())
			}
			for _, n := range tc.expSet.GetItems() {
				if !s.In(n) {
					t.Errorf("expecting %d to be present in the set but was not; set is %v", n, s.GetItems())
				}
			}
		})
	}
}

func TestSymmetricDifference(t *testing.T) {
	td := []struct {
		name   string
		s1     Set
		s2     Set
		expSet Set
	}{
		{"symmetric difference of different sets", New(1, 2, 3), New(4, 5, 6), New(1, 2, 3, 4, 5, 6)},
		{"symmetric difference of sets with elements in common", New(1, 2, 3), New(1, 2, 4), New(3, 4)},
		{"symmetric difference of same sets", New(1, 2, 3), New(1, 2, 3), New()},
	}
	for _, tc := range td {
		t.Run(tc.name, func(t *testing.T) {
			s := tc.s1.SymmetricDifference(tc.s2)
			if s.Len() != tc.expSet.Len() {
				t.Errorf("expecting %d elements in the set but got %d: set is %v", tc.expSet.Len(), s.Len(), s.GetItems())
			}
			for _, n := range tc.expSet.GetItems() {
				if !s.In(n) {
					t.Errorf("expecting %d to be present in the set but was not; set is %v", n, s.GetItems())
				}
			}
		})
	}
}
