package graph

import (
	"testing"
)

func TestUnionFind(t *testing.T) {
	u := NewUnionFind(10) // Creating a Union-Find data structure with 10 elements

	//union operations
	u.Union(0, 1)
	u.Union(2, 3)
	u.Union(4, 5)
	u.Union(6, 7)

	// Testing the parent of specific elements
	t.Run("Test Find", func(t *testing.T) {
		if u.Find(0) != u.Find(1) || u.Find(2) != u.Find(3) || u.Find(4) != u.Find(5) || u.Find(6) != u.Find(7) {
			t.Error("Union operation not functioning correctly")
		}
	})

	u.Union(1, 5) // Additional union operation
	u.Union(3, 7) // Additional union operation

	// Testing the parent of specific elements after more union operations
	t.Run("Test Find after Union", func(t *testing.T) {
		if u.Find(0) != u.Find(5) || u.Find(1) != u.Find(4) || u.Find(2) != u.Find(7) || u.Find(3) != u.Find(6) {
			t.Error("Union operation not functioning correctly")
		}
	})

	u.Union(3, 7) // Repeated union operation

	// Testing that repeated union operations are idempotent
	t.Run("Test Find after repeated Union", func(t *testing.T) {
		if u.Find(2) != u.Find(6) || u.Find(2) != u.Find(7) || u.Find(3) != u.Find(6) || u.Find(3) != u.Find(7) {
			t.Error("Union operation not functioning correctly")
		}
	})
}
