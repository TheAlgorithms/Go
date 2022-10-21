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
		t.pushHelper(t.Root, key)
	}
}

// Delete removes the node of val
func (t *BinarySearch[T]) Delete(val T) bool {
	node, ok := t.Get(val)
	if !ok {
		return false
	}
	t.deleteHelper(node)
	return true
}

func (t *BinarySearch[T]) pushHelper(x *Node[T], val T) {
	z := &Node[T]{Key: val, Left: nil, Right: nil}
	var y *Node[T]
	for !t.isNil(x) {
		y = x
		if val < x.Key {
			x = x.Left
		} else {
			x = x.Right
		}
	}
	z.Parent = y
	if t.isNil(y) {
		t.Root = z
	} else if val < y.Key {
		y.Left = z
	} else {
		y.Right = z
	}
}

func (t *BinarySearch[T]) deleteHelper(z *Node[T]) {
	switch {
	case z.Left == nil:
		t.transplant(z, z.Right)
	case z.Right == nil:
		t.transplant(z, z.Left)
	default:
		y := t.minimum(z.Right)
		if y.Parent != z {
			t.transplant(y, y.Right)
			y.Right = z.Right
			y.Right.Parent = y
		}

		t.transplant(z, y)
		y.Left = z.Left
		y.Left.Parent = y
	}
}