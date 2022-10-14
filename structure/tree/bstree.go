// Binary search tree.
//
// For more details check out those links below here:
// Wikipedia article: https://en.wikipedia.org/wiki/Binary_search_tree
// see bstree.go

package tree

import (
	"github.com/TheAlgorithms/Go/constraints"
)

type BSTree[T constraints.Ordered] struct {
	*binaryTree[T]
}

// NewBSTree create a novel Binary-Search tree
func NewBSTree[T constraints.Ordered]() *BSTree[T] {
	return &BSTree[T]{
		binaryTree: &binaryTree[T]{
			Root: nil,
			NIL:  nil,
		},
	}
}

// Insert a chain of Node's into the BSTree
func (t *BSTree[T]) Insert(keys ...T) {
	for _, key := range keys {
		t.Root = t.insertHelper(t.Root, key)
	}
}

// // Delete removes the node of val
func (t *BSTree[T]) Delete(val T) *Node[T] {
	return t.deleteHelper(t.Root, val)
}

// AccessNodesByLayer  accesses nodes layer by layer instead of printing the results as one line.
func (t *BSTree[T]) AccessNodesByLayer() [][]T {
	var res [][]T
	root := t.Root
	if root == nil {
		return res
	}
	var q []*Node[T]
	var n *Node[T]
	var idx = 0
	q = append(q, root)

	for len(q) != 0 {
		res = append(res, []T{})
		qLen := len(q)
		for i := 0; i < qLen; i++ {
			n, q = q[0], q[1:]
			res[idx] = append(res[idx], n.Key)
			if n.Left != nil {
				q = append(q, n.Left)
			}
			if n.Right != nil {
				q = append(q, n.Right)
			}
		}
		idx++
	}
	return res
}

func (t *BSTree[T]) insertHelper(root *Node[T], val T) *Node[T] {
	if root == nil {
		return &Node[T]{Key: val, Left: nil, Right: nil}
	}
	if val < root.Key {
		root.Left = t.insertHelper(root.Left, val)
	} else {
		root.Right = t.insertHelper(root.Right, val)
	}
	return root
}

func (t *BSTree[T]) deleteHelper(root *Node[T], val T) *Node[T] {
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
