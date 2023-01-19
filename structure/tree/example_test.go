package tree_test

import (
	"fmt"
	"github.com/TheAlgorithms/Go/constraints"
	"testing"

	bt "github.com/TheAlgorithms/Go/structure/tree"
)

type TestTree[T constraints.Ordered] interface {
	Push(...T)
	Delete(T) bool
	Get(T) (bt.Node[T], bool)
	Empty() bool
	Has(T) bool
	Depth() int
	Max() (T, bool)
	Min() (T, bool)
	Predecessor(T) (T, bool)
	Successor(T) (T, bool)
	PreOrder() []T
	InOrder() []T
	PostOrder() []T
	LevelOrder() []T
	AccessNodesByLayer() [][]T
}

// BinarySearch, AVL and RB have completed the `TestTree` interface.
var (
	_ TestTree[int] = (*bt.BinarySearch[int])(nil)
	_ TestTree[int] = (*bt.AVL[int])(nil)
	_ TestTree[int] = (*bt.RB[int])(nil)
)

var tree TestTree[int]

func TestBinarySearch(t *testing.T) {
	tree = bt.NewBinarySearch[int]()

	if tree.Empty() {
		t.Log("Binary Search Tree is empty now.")
	}

	tree.Push(1, 4, 10)
	tree.Push(-8)
	nums := []int{87, 18, 10, -34}
	tree.Push(nums...)
	tree.Push(4) // duplicate key, dismiss it

	if tree.Has(4) {
		t.Logf("There is a node of 4")
	}

	if n, ok := tree.Get(10); ok {
		t.Logf("node of 10: %T", n)
	}

	if ret, ok := tree.Min(); ok {
		t.Logf("tree.Min() = %v", ret)
	}

	if ret, ok := tree.Max(); ok {
		t.Logf("tree.Max() = %v", ret)
	}

	if ret, ok := tree.Predecessor(1); ok {
		t.Logf("tree.Preducessor(1) = %v", ret)
	}

	if ret, ok := tree.Successor(18); ok {
		t.Logf("tree.Successor(18) = %v", ret)
	}

	fmt.Println(tree.InOrder())
	fmt.Println(tree.AccessNodesByLayer())

	tree.Delete(18)
	fmt.Println("Delete 18")
	fmt.Println(tree.InOrder())
}

func TestAVL(t *testing.T) {
	tree = bt.NewAVL[int]()

	if tree.Empty() {
		t.Log("AVL Tree is empty now.")
	}

	tree.Push(1, 4, 10)
	tree.Push(-8)
	nums := []int{87, 18, 10, -34}
	tree.Push(nums...)
	tree.Push(4) // duplicate key, dismiss it

	if tree.Has(4) {
		t.Logf("There is a node of 4")
	}

	if n, ok := tree.Get(10); ok {
		t.Logf("node of 10: %T", n)
	}

	if ret, ok := tree.Min(); ok {
		t.Logf("tree.Min() = %v", ret)
	}

	if ret, ok := tree.Max(); ok {
		t.Logf("tree.Max() = %v", ret)
	}

	if ret, ok := tree.Predecessor(1); ok {
		t.Logf("tree.Preducessor(1) = %v", ret)
	}

	if ret, ok := tree.Successor(18); ok {
		t.Logf("tree.Successor(18) = %v", ret)
	}

	fmt.Println(tree.InOrder())
	fmt.Println(tree.AccessNodesByLayer())

	tree.Delete(18)
	fmt.Println("Delete 18")
	fmt.Println(tree.InOrder())
}

func TestRB(t *testing.T) {
	tree = bt.NewRB[int]()

	if tree.Empty() {
		t.Log("RB Tree is empty now.")
	}

	tree.Push(1, 4, 10)
	tree.Push(-8)
	nums := []int{87, 18, 10, -34}
	tree.Push(nums...)
	tree.Push(4) // duplicate key, dismiss it

	if tree.Has(4) {
		t.Logf("There is a node of 4")
	}

	if n, ok := tree.Get(10); ok {
		t.Logf("node of 10: %T", n)
	}

	if ret, ok := tree.Min(); ok {
		t.Logf("tree.Min() = %v", ret)
	}

	if ret, ok := tree.Max(); ok {
		t.Logf("tree.Max() = %v", ret)
	}

	if ret, ok := tree.Predecessor(1); ok {
		t.Logf("tree.Preducessor(1) = %v", ret)
	}

	if ret, ok := tree.Successor(18); ok {
		t.Logf("tree.Successor(18) = %v", ret)
	}

	fmt.Println(tree.InOrder())
	fmt.Println(tree.AccessNodesByLayer())
}
