// Binary search tree.
//
// For more details check out those links below here:
// Wikipedia article: https://en.wikipedia.org/wiki/Binary_search_tree
// see bstree.go

package tree

import (
	"github.com/TheAlgorithms/Go/constraints"
	"github.com/TheAlgorithms/Go/math/max"
)

// NewNode Returns a new pointer to an empty Node
func NewNode[T constraints.Ordered](val T) *Node[T] {
	return &Node[T]{Key: val, Left: nil, Right: nil}
}

type BSTree[T constraints.Ordered] struct {
	Root *Node[T]
}

// Insert a value in the BSTree
func (bt *BSTree[T]) Insert(val T) *Node[T] {
	return bt.insertHelper(bt.Root, val)
}

// Depth returns the calculated depth of a binary search tree
func (bt *BSTree[T]) Depth() int {
	return bt.calculateDepth(bt.Root, 0)
}

// InOrderSuccessor Goes to the Left
func (bt *BSTree[T]) InOrderSuccessor(root *Node[T]) *Node[T] {
	cur := root
	for cur.Left != nil {
		cur = cur.Left
	}
	return cur
}

// // Delete removes the node of val
func (bt *BSTree[T]) Delete(val T) *Node[T] {
	return bt.deleteHelper(bt.Root, val)
}

// Traverses the tree in the following order Left --> Root --> Right
func (bt *BSTree[T]) InOrder() []T {
	traversal := make([]T, 0)
	bt.inOrderRecursive(bt.Root, &traversal)
	return traversal
}

// Traverses the tree in the following order Root --> Left --> Right
func (bt *BSTree[T]) PreOrder() []T {
	traversal := make([]T, 0)
	bt.preOrderRecursive(bt.Root, &traversal)
	return traversal
}

// Traverses the tree in the following order Left --> Right --> Root
func (bt *BSTree[T]) PostOrder() []T {
	traversal := make([]T, 0)
	bt.postOrderRecursive(bt.Root, &traversal)
	return traversal
}

// LevelOrder returns the level order traversal of the tree
func (bt *BSTree[T]) LevelOrder() []T {
	traversal := make([]T, 0)
	bt.levelOrderHelper(bt.Root, &traversal)
	return traversal
}

// AccessNodesByLayer Function that access nodes layer by layer instead of printing the results as one line.
func (bt *BSTree[T]) AccessNodesByLayer() [][]T {
	var res [][]T
	root := bt.Root
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

func (bt *BSTree[T]) insertHelper(root *Node[T], val T) *Node[T] {
	if root == nil {
		return NewNode(val)
	}
	if val < root.Key {
		root.Left = bt.insertHelper(root.Left, val)
	} else {
		root.Right = bt.insertHelper(root.Right, val)
	}
	return root
}

func (bt *BSTree[T]) calculateDepth(n *Node[T], depth int) int {
	if n == nil {
		return depth
	}
	return max.Int(bt.calculateDepth(n.Left, depth+1), bt.calculateDepth(n.Right, depth+1))
}

func (bt *BSTree[T]) deleteHelper(root *Node[T], val T) *Node[T] {
	if root == nil {
		return nil
	}
	if val < root.Key {
		root.Left = bt.deleteHelper(root.Left, val)
	} else if val > root.Key {
		root.Right = bt.deleteHelper(root.Right, val)
	} else {
		// this is the node to delete
		// node with one child
		if root.Left == nil {
			return root.Right
		} else if root.Right == nil {
			return root.Left
		} else {
			n := root.Right
			d := bt.InOrderSuccessor(n)
			d.Left = root.Left
			return root.Right
		}
	}
	return root
}

func (bt *BSTree[T]) inOrderRecursive(n *Node[T], traversal *[]T) {
	if n == nil {
		return
	}

	bt.inOrderRecursive(n.Left, traversal)
	*traversal = append(*traversal, n.Key)
	bt.inOrderRecursive(n.Right, traversal)
}

func (bt *BSTree[T]) preOrderRecursive(n *Node[T], traversal *[]T) {
	if n == nil {
		return
	}

	*traversal = append(*traversal, n.Key)
	bt.preOrderRecursive(n.Left, traversal)
	bt.preOrderRecursive(n.Right, traversal)
}

func (bt *BSTree[T]) postOrderRecursive(n *Node[T], traversal *[]T) {
	if n == nil {
		return
	}

	bt.postOrderRecursive(n.Left, traversal)
	bt.postOrderRecursive(n.Right, traversal)
	*traversal = append(*traversal, n.Key)
}

func (bt *BSTree[T]) levelOrderHelper(root *Node[T], traversal *[]T) {
	var q []*Node[T] // queue
	var n *Node[T]   // temporary node

	q = append(q, root)

	for len(q) != 0 {
		n, q = q[0], q[1:]
		*traversal = append(*traversal, n.Key)
		if n.Left != nil {
			q = append(q, n.Left)
		}
		if n.Right != nil {
			q = append(q, n.Right)
		}
	}
}
