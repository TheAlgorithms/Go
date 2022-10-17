// Binary search tree.
//
// For more details check out those links below here:
// Wikipedia article: https://en.wikipedia.org/wiki/Binary_search_tree
// see bstree.go

package tree

import (
	"github.com/TheAlgorithms/Go/constraints"
)

type BinarySearch[T constraints.Ordered] struct {
	*binaryTree[T]
}

// NewBinarySearch create a novel Binary-Search tree
func NewBinarySearch[T constraints.Ordered]() *BinarySearch[T] {
	return &BinarySearch[T]{
		binaryTree: &binaryTree[T]{
			Root: nil,
			NIL:  nil,
		},
	}
}

// Push a chain of Node's into the BinarySearch
func (t *BinarySearch[T]) Push(keys ...T) {
	for _, key := range keys {
		t.Root = t.pushHelper(t.Root, key)
	}
}

// Delete removes the node of val
func (t *BinarySearch[T]) Delete(val T) *Node[T] {
	return t.deleteHelper(t.Root, val)
}

func (t *BinarySearch[T]) pushHelper(root *Node[T], val T) *Node[T] {
	if root == nil {
		return &Node[T]{Key: val, Left: nil, Right: nil}
	}
	if val < root.Key {
		root.Left = t.pushHelper(root.Left, val)
	} else {
		root.Right = t.pushHelper(root.Right, val)
	}
	return root
}

func (t *BinarySearch[T]) deleteHelper(root *Node[T], val T) *Node[T] {
	if root == nil {
		return nil
	}
	if val < root.Key {
		root.Left = t.deleteHelper(root.Left, val)
	} else if val > root.Key {
		root.Right = t.deleteHelper(root.Right, val)
	} else {
		// this is the node to delete
		// node with one child
		if root.Left == nil {
			return root.Right
		} else if root.Right == nil {
			return root.Left
		} else {
			n := root.Right
			d := t.minimum(n)
			d.Left = root.Left
			return root.Right
		}
	}
	return root
}
