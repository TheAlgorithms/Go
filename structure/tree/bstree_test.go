package tree_test

import (
	"math/rand"
	"sort"
	"testing"
	"time"

	bt "github.com/TheAlgorithms/Go/structure/tree"
)

func TestPush(t *testing.T) {
	bst := bt.NewBinarySearch[int]()

	bst.Push(90)
	bst.Push(80)
	bst.Push(100)

	if bst.Root.Key() != 90 {
		t.Errorf("Root should have value = 90")
	}

	if bst.Root.Left().Key() != 80 {
		t.Errorf("Left child should have value = 80")
	}

	if bst.Root.Right().Key() != 100 {
		t.Errorf("Right child should have value = 100")
	}

	if bst.Depth() != 2 {
		t.Errorf("tree should have depth = 1")
	}
}

func TestDelete(t *testing.T) {
	t.Run("Delete a node with no child", func(t *testing.T) {
		bst := bt.NewBinarySearch[int]()

		bst.Push(90)
		bst.Push(80)
		bst.Push(80)
		bst.Push(100)

		if !bst.Delete(100) {
			t.Errorf("There is a node, whose value is 100")
		}

		if bst.Delete(105) {
			t.Errorf("There is no node, whose value is 105")
		}

		root := bst.Root
		if root.Key() != 90 {
			t.Errorf("Root should have value = 90")
		}

		if root.Left().Key() != 80 {
			t.Errorf("Left child should have value = 80")
		}

		if root.Right().(*bt.BSNode[int]) != nil {
			t.Errorf("Right child should have value = nil")
		}

		if bst.Depth() != 2 {
			t.Errorf("Depth should have value = 2")
		}

		bst.Delete(80)

		if root.Key() != 90 {
			t.Errorf("Root should have value = 90")
		}

		if root.Left().(*bt.BSNode[int]) != nil {
			t.Errorf("Left child should have value = nil")
		}

		if bst.Depth() != 1 {
			t.Errorf("Depth should have value = 1")
		}
	})

	t.Run("Delete a node with one child", func(t *testing.T) {
		bst := bt.NewBinarySearch[int]()

		bst.Push(90)
		bst.Push(80)
		bst.Push(100)
		bst.Push(70)

		if bst.Delete(102) {
			t.Errorf("There is no node, whose value is 102")
		}

		if !bst.Delete(80) {
			t.Errorf("There is a node, whose value is 80")
		}

		root := bst.Root
		if root.Key() != 90 {
			t.Errorf("Root should have value = 90")
		}

		if root.Right().Key() != 100 {
			t.Errorf("Right child should have value = 100")
		}

		if root.Left().Key() != 70 {
			t.Errorf("Left child should have value = 70")
		}

		if bst.Depth() != 2 {
			t.Errorf("Depth should have value = 2")
		}
	})

	t.Run("Delete a node with two children", func(t *testing.T) {
		bst := bt.NewBinarySearch[int]()

		bst.Push(90)
		bst.Push(80)
		bst.Push(100)
		bst.Push(70)
		bst.Push(85)

		if !bst.Delete(80) {
			t.Errorf("There is a node, whose value is 80")
		}

		if bst.Delete(102) {
			t.Errorf("There is no node, whose value is 102")
		}

		root := bst.Root
		if root.Key() != 90 {
			t.Errorf("Root should have value = 90")
		}

		if root.Left().Key() != 85 {
			t.Errorf("Left child should have value = 85")
		}

		if root.Right().Key() != 100 {
			t.Errorf("Right child should have value = 100")
		}

		if bst.Depth() != 3 {
			t.Errorf("Depth should have value = 3")
		}
	})

	t.Run("Random Test", func(t *testing.T) {
		tests := []int{100, 500, 1000, 10_000}
		for _, n := range tests {
			rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
			tree := bt.NewBinarySearch[int]()
			nums := rnd.Perm(n)
			tree.Push(nums...)

			rets := tree.InOrder()
			if !sort.IntsAreSorted(rets) {
				t.Error("Error with Push")
			}

			if res, ok := tree.Min(); !ok || res != rets[0] {
				t.Errorf("Error with Min, get %d, want: %d", res, rets[0])
			}

			if res, ok := tree.Max(); !ok || res != rets[n-1] {
				t.Errorf("Error with Max, get %d, want: %d", res, rets[n-1])
			}

			for i := 0; i < n-1; i++ {
				if ret, ok := tree.Successor(rets[0]); ret != rets[1] || !ok {
					t.Error("Error with Successor")
				}

				if ret, ok := tree.Predecessor(rets[1]); ret != rets[0] || !ok {
					t.Error("Error with Predecessor")
				}

				ok := tree.Delete(nums[i])
				rets = tree.InOrder()
				if !ok || !sort.IntsAreSorted(rets) {
					t.Errorf("Error With Delete")
				}
			}
		}
	})
}
