package tree_test

import (
	"fmt"
	"testing"

	bt "github.com/TheAlgorithms/Go/structure/tree"
)

// BinarySearch, AVL and RB have completed the `bt.Tree` interface.
var tree bt.Tree[int]

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
