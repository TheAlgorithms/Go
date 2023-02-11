package tree_test

import (
	bt "github.com/TheAlgorithms/Go/structure/tree"
	"math/rand"
	"sort"
	"testing"
	"time"
)

func TestAVLPush(t *testing.T) {
	t.Run("LLRotation-Test", func(t *testing.T) {
		tree := bt.NewAVL[int]()
		tree.Push(5, 4, 3)

		root := tree.Root
		if root.Key() != 4 {
			t.Errorf("Root should have value = 4, not %v", root.Key())
		}
		if root.Height() != 2 {
			t.Errorf("Height of Root should be = 2, not %d", root.Height())
		}

		if root.Left().Key() != 3 {
			t.Errorf("Left child should have value = 3")
		}
		if root.Left().(*bt.AVLNode[int]).Height() != 1 {
			t.Errorf("Height of Left child should be 1")
		}

		if root.Right().Key() != 5 {
			t.Errorf("Right child should have value = 5")
		}
		if root.Right().(*bt.AVLNode[int]).Height() != 1 {
			t.Errorf("Height of Right should be 1")
		}
	})

	t.Run("LRRotation-Test", func(t *testing.T) {
		tree := bt.NewAVL[int]()
		tree.Push(5, 3, 4)

		root := tree.Root
		if root.Key() != 4 {
			t.Errorf("Root should have value = 4, not %v", root.Key())
		}
		if root.Height() != 2 {
			t.Errorf("Height of Root should be = 2, not %d", root.Height())
		}

		if root.Left().Key() != 3 {
			t.Errorf("Left child should have value = 3")
		}
		if root.Left().(*bt.AVLNode[int]).Height() != 1 {
			t.Errorf("Height of Left child should be 1")
		}

		if root.Right().Key() != 5 {
			t.Errorf("Right child should have value = 5")
		}
		if root.Right().(*bt.AVLNode[int]).Height() != 1 {
			t.Errorf("Height of Right should be 1")
		}
	})

	t.Run("RRRotation-Test", func(t *testing.T) {
		tree := bt.NewAVL[int]()
		tree.Push(3)
		tree.Push(4)
		tree.Push(5)

		root := tree.Root
		if root.Key() != 4 {
			t.Errorf("Root should have value = 4, not %v", root.Key())
		}
		if root.Height() != 2 {
			t.Errorf("Height of Root should be = 2, not %d", root.Height())
		}

		if root.Left().Key() != 3 {
			t.Errorf("Left child should have value = 3")
		}
		if root.Left().(*bt.AVLNode[int]).Height() != 1 {
			t.Errorf("Height of Left child should be 1")
		}

		if root.Right().Key() != 5 {
			t.Errorf("Right child should have value = 5")
		}
		if root.Right().(*bt.AVLNode[int]).Height() != 1 {
			t.Errorf("Height of Right should be 1")
		}
	})
	t.Run("RLRotation-Test", func(t *testing.T) {
		tree := bt.NewAVL[int]()
		tree.Push(3)
		tree.Push(5)
		tree.Push(4)

		root := tree.Root
		if root.Key() != 4 {
			t.Errorf("Root should have value = 4")
		}
		if root.Height() != 2 {
			t.Errorf("Height of Root should be = 2")
		}

		if root.Left().Key() != 3 {
			t.Errorf("Left child should have value = 3")
		}
		if root.Left().(*bt.AVLNode[int]).Height() != 1 {
			t.Errorf("Height of Left child should be 1")
		}

		if root.Right().Key() != 5 {
			t.Errorf("Right child should have value = 5")
		}
		if root.Right().(*bt.AVLNode[int]).Height() != 1 {
			t.Errorf("Height of Right should be 1")
		}
	})
}

func TestAVLDelete(t *testing.T) {
	t.Run("LLRotation-Test", func(t *testing.T) {
		tree := bt.NewAVL[int]()
		if tree.Delete(5) {
			t.Errorf("There is no node, whose value is 5")
		}

		tree.Push(5)
		tree.Push(4)
		tree.Push(3)
		tree.Push(2)

		if !tree.Delete(5) {
			t.Errorf("There is a node, whose value is 5")
		}

		if tree.Delete(50) {
			t.Errorf("There is no node, whose value is 50")
		}

		root := tree.Root
		if root.Key() != 3 {
			t.Errorf("Root should have value = 3")
		}
		if root.Height() != 2 {
			t.Errorf("Height of Root should be = 2")
		}

		if root.Left().Key() != 2 {
			t.Errorf("Left child should have value = 2")
		}
		if root.Left().(*bt.AVLNode[int]).Height() != 1 {
			t.Errorf("Height of Left child should be 1")
		}

		if root.Right().Key() != 4 {
			t.Errorf("Right child should have value = 5")
		}
		if root.Right().(*bt.AVLNode[int]).Height() != 1 {
			t.Errorf("Height of Right should be 1")
		}
	})

	t.Run("LRRotation-Test", func(t *testing.T) {
		tree := bt.NewAVL[int]()

		tree.Push(10)
		tree.Push(8)
		tree.Push(8)
		tree.Push(6)
		tree.Push(7)

		if !tree.Delete(10) {
			t.Errorf("There is a node, whose value is 10")
		}

		if tree.Delete(5) {
			t.Errorf("There is no node, whose value is 5")
		}

		root := tree.Root
		if root.Key() != 7 {
			t.Errorf("Root should have value = 7")
		}
		if root.Height() != 2 {
			t.Errorf("Height of Root should be = 2")
		}

		if root.Left().Key() != 6 {
			t.Errorf("Left child should have value = 6")
		}
		if root.Left().(*bt.AVLNode[int]).Height() != 1 {
			t.Errorf("Height of Left child should be 1")
		}

		if root.Right().Key() != 8 {
			t.Errorf("Right child should have value = 8")
		}
		if root.Right().(*bt.AVLNode[int]).Height() != 1 {
			t.Errorf("Height of Right should be 1")
		}

	})

	t.Run("RRRotation-Test", func(t *testing.T) {
		tree := bt.NewAVL[int]()

		tree.Push(2)
		tree.Push(3)
		tree.Push(3)
		tree.Push(4)
		tree.Push(5)

		if !tree.Delete(2) {
			t.Errorf("There is a node, whose value is 2")
		}

		if tree.Delete(15) {
			t.Errorf("There is no node, whose value is 15")
		}

		root := tree.Root
		if root.Key() != 4 {
			t.Errorf("Root should have value = 4")
		}
		if root.Height() != 2 {
			t.Errorf("Height of Root should be = 2")
		}

		if root.Left().Key() != 3 {
			t.Errorf("Left child should have value = 3")
		}
		if root.Left().(*bt.AVLNode[int]).Height() != 1 {
			t.Errorf("Height of Left child should be 1")
		}

		if root.Right().Key() != 5 {
			t.Errorf("Right child should have value = 5")
		}
		if root.Right().(*bt.AVLNode[int]).Height() != 1 {
			t.Errorf("Height of Right should be 1")
		}

	})

	t.Run("RLRotation-Test", func(t *testing.T) {
		tree := bt.NewAVL[int]()

		tree.Push(7)
		tree.Push(6)
		tree.Push(6)
		tree.Push(9)
		tree.Push(8)

		tree.Delete(6)

		root := tree.Root
		if root.Key() != 8 {
			t.Errorf("Root should have value = 8")
		}
		if root.Height() != 2 {
			t.Errorf("Height of Root should be = 2")
		}

		if root.Left().Key() != 7 {
			t.Errorf("Left child should have value = 7")
		}
		if root.Left().(*bt.AVLNode[int]).Height() != 1 {
			t.Errorf("Height of Left child should be 1")
		}

		if root.Right().Key() != 9 {
			t.Errorf("Right child should have value = 9")
		}
		if root.Right().(*bt.AVLNode[int]).Height() != 1 {
			t.Errorf("Height of Right should be 1")
		}

	})

	t.Run("Random Test", func(t *testing.T) {
		nums := []int{100, 500, 1000, 10_000}
		for _, n := range nums {
			rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
			tree := bt.NewAVL[int]()
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
