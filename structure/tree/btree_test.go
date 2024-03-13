package tree_test

import (
	bt "github.com/TheAlgorithms/Go/structure/tree"
	"math/rand"
	"testing"
)

func TestBTreeIncreasing(t *testing.T) {
	maxKeysCases := []int{4, 16}
	sizes := []int{100, 0xBA5, 0xF00}
	for _, maxKeys := range maxKeysCases {
		for _, size := range sizes {
			tree := bt.NewBTree[int](maxKeys)
			if tree.Search(0) {
				t.Errorf("Tree expected to contain 0")
			}
			for i := 0; i < size; i++ {
				tree.Insert(i)
			}
			for i := 0; i < size; i++ {
				if !tree.Search(i) {
					t.Errorf("Tree expected to contain %d", i)
				}
			}
			if tree.Search(size + 1) {
				t.Errorf("Tree not expected to contain %d", size+1)
			}

			for i := 0; i < size; i += 5 {
				tree.Delete(i)
			}
			for i := 0; i < size; i++ {
				hasKey := tree.Search(i)
				if i%5 == 0 && hasKey {
					t.Errorf("Tree not expected to contain %d", i)
				} else if i%5 != 0 && !hasKey {
					t.Errorf("Tree expected to contain %d", i)
				}
			}
		}
	}
}

func TestBTreeDecreasing(t *testing.T) {
	maxKeysCases := []int{4, 16}
	sizes := []int{100, 1000}
	for _, maxKeys := range maxKeysCases {
		for _, size := range sizes {
			tree := bt.NewBTree[int](maxKeys)
			if tree.Search(0) {
				t.Errorf("Tree expected to contain 0")
			}
			for i := size - 1; i >= 0; i-- {
				tree.Insert(i)
			}
			for i := 0; i < size; i++ {
				if !tree.Search(i) {
					t.Errorf("Tree expected to contain %d", i)
				}
			}
			if tree.Search(size + 1) {
				t.Errorf("Tree not expected to contain %d", size+1)
			}

			for i := 0; i < size; i += 5 {
				tree.Delete(i)
			}
			for i := 0; i < size; i++ {
				hasKey := tree.Search(i)
				if i%5 == 0 && hasKey {
					t.Errorf("Tree not expected to contain %d", i)
				} else if i%5 != 0 && !hasKey {
					t.Errorf("Tree expected to contain %d", i)
				}
			}
		}
	}
}

func TestBTreeRandom(t *testing.T) {
	maxKeysCases := []int{4, 16}
	sizes := []int{100, 0xBA5, 0xF00}
	for _, maxKeys := range maxKeysCases {
		for _, size := range sizes {
			rnd := rand.New(rand.NewSource(0))
			tree := bt.NewBTree[int](maxKeys)
			nums := rnd.Perm(size)
			if tree.Search(0) {
				t.Errorf("Tree expected to contain 0")
			}
			for i := 0; i < size; i++ {
				tree.Insert(nums[i])
			}
			for i := 0; i < size; i++ {
				if !tree.Search(nums[i]) {
					t.Errorf("Tree expected to contain %d", nums[i])
				}
			}

			for i := 0; i < size; i += 5 {
				tree.Delete(nums[i])
			}
			for i := 0; i < size; i++ {
				hasKey := tree.Search(nums[i])
				if i%5 == 0 && hasKey {
					t.Errorf("Tree not expected to contain %d", i)
				} else if i%5 != 0 && !hasKey {
					t.Errorf("Tree expected to contain %d", i)
				}
			}
		}
	}
}

func TestBTreeDeleteEverything(t *testing.T) {
	tree := bt.NewBTree[int](4)
	size := 128
	for i := 0; i < size; i++ {
		tree.Insert(i)
	}
	for i := 0; i < size; i++ {
		tree.Delete(i)
	}
	tree.Delete(-1)
	tree.Delete(1000)

	for i := 0; i < size; i++ {
		if tree.Search(i) {
			t.Errorf("Tree not expected to contain %d", i)
		}
	}
}
