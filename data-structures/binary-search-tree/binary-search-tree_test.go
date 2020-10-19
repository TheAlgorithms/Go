package binarysearchtree

import (
	"fmt"
	"testing"
)

func TestBinarySearchTree(t *testing.T) {
	bt := &btree{nil}
	inorder(bt.root)
	bt.root = insert(bt.root, 30)

	bt.root = insert(bt.root, 20)
	bt.root = insert(bt.root, 15)
	bt.root = insert(bt.root, 10)
	bt.root = insert(bt.root, 12)
	bt.root = insert(bt.root, 9)
	bt.root = insert(bt.root, 11)
	bt.root = insert(bt.root, 17)
	fmt.Print(bt.depth(), "\n")
	inorder(bt.root)
	fmt.Print("\n")
	bt.root = bst_delete(bt.root, 10)
	inorder(bt.root)
	fmt.Print("\n")
	bt.root = bst_delete(bt.root, 30)
	inorder(bt.root)
	fmt.Print("\n")
	bt.root = bst_delete(bt.root, 15)
	inorder(bt.root)
	fmt.Print("\n")
	bt.root = bst_delete(bt.root, 20)
	inorder(bt.root)
	fmt.Print("\n")
	fmt.Print(bt.depth(), "\n")
}
