package binarytree

import (
	"fmt"
	"testing"
)

func testTree() btree {
	bt := btree{nil}

	//        0
	//   1         2
	// 3   4     5   6
	//                 10

	bt.root = newNode(0)
	bt.root.left = newNode(1)
	bt.root.right = newNode(2)
	bt.root.left.left = newNode(3)
	bt.root.left.right = newNode(4)
	bt.root.right.left = newNode(5)
	bt.root.right.right = newNode(6)
	bt.root.right.right.right = newNode(10)

	return bt
}

func TestDepth(t *testing.T) {
	bt := testTree()
	expectedDepth := 4
	depth := bt.depth()
	if depth != expectedDepth {
		t.Errorf("Incorrect tree depth. Expected %d and got %d\n", expectedDepth, depth)
	}
}

func TestTraverseInOrder(t *testing.T) {
	bt := testTree()
	inorder(bt.root)
	fmt.Print("\n")
}

func TestTraversePreOrder(t *testing.T) {
	bt := testTree()
	preorder(bt.root)
	fmt.Print("\n")
}

func TestTraversePostOrder(t *testing.T) {
	bt := testTree()
	postorder(bt.root)
	fmt.Print("\n")
}

func TestTraverseLevelOrder(t *testing.T) {
	bt := testTree()
	levelorder(bt.root)
	fmt.Print("\n")
}
