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
	return &Node[T]{key: val, left: nil, right: nil}
}

type BSTree[T constraints.Ordered] struct {
	root *Node[T]
}

// Insert a value in the BSTree
func (bt *BSTree[T]) Insert(val T) *Node[T] {
	return bt.insertHelper(bt.root, val)
}

// Depth returns the calculated depth of a binary search tree
func (bt *BSTree[T]) Depth() int {
	return bt.calculateDepth(bt.root, 0)
}

// InOrderSuccessor Goes to the left
func (bt *BSTree[T]) InOrderSuccessor(root *Node[T]) *Node[T] {
	cur := root
	for cur.left != nil {
		cur = cur.left
	}
	return cur
}

func (bt *BSTree[T]) Delete(val T) *Node[T] {
	return bt.deleteHelper(bt.root, val)
}

// Traverses the tree in the following order left --> root --> right
func (bt *BSTree[T]) InOrder() []T {
	traversal := make([]T, 0)
	bt.inOrderRecursive(bt.root, &traversal)
	return traversal
}

// Traverses the tree in the following order root --> left --> right
func (bt *BSTree[T]) PreOrder() []T {
	traversal := make([]T, 0)
	bt.preOrderRecursive(bt.root, &traversal)
	return traversal
}

// Traverses the tree in the following order left --> right --> root
func (bt *BSTree[T]) PostOrder() []T {
	traversal := make([]T, 0)
	bt.postOrderRecursive(bt.root, &traversal)
	return traversal
}

func (bt *BSTree[T]) LevelOrder() []T {
	traversal := make([]T, 0)
	bt.levelOrderHelper(bt.root, &traversal)
	return traversal
}

func (bt *BSTree[T]) AccessNodesByLayer() [][]T {
	var res [][]T
	root := bt.root
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
			res[idx] = append(res[idx], n.key)
			if n.left != nil {
				q = append(q, n.left)
			}
			if n.right != nil {
				q = append(q, n.right)
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
	if val < root.key {
		root.left = bt.insertHelper(root.left, val)
	} else {
		root.right = bt.insertHelper(root.right, val)
	}
	return root
}

func (bt *BSTree[T]) calculateDepth(n *Node[T], depth int) int {
	if n == nil {
		return depth
	}
	return max.Int(bt.calculateDepth(n.left, depth+1), bt.calculateDepth(n.right, depth+1))
}

func (bt *BSTree[T]) deleteHelper(root *Node[T], val T) *Node[T] {
	if root == nil {
		return nil
	}
	if val < root.key {
		root.left = bt.deleteHelper(root.left, val)
	} else if val > root.key {
		root.right = bt.deleteHelper(root.right, val)
	} else {
		// this is the node to delete
		// node with one child
		if root.left == nil {
			return root.right
		} else if root.right == nil {
			return root.left
		} else {
			n := root.right
			d := bt.InOrderSuccessor(n)
			d.left = root.left
			return root.right
		}
	}
	return root
}

func (bt *BSTree[T]) inOrderRecursive(n *Node[T], traversal *[]T) {
	if n == nil {
		return
	}

	bt.inOrderRecursive(n.left, traversal)
	*traversal = append(*traversal, n.key)
	bt.inOrderRecursive(n.right, traversal)
}

func (bt *BSTree[T]) preOrderRecursive(n *Node[T], traversal *[]T) {
	if n == nil {
		return
	}

	*traversal = append(*traversal, n.key)
	bt.preOrderRecursive(n.left, traversal)
	bt.preOrderRecursive(n.right, traversal)
}

func (bt *BSTree[T]) postOrderRecursive(n *Node[T], traversal *[]T) {
	if n == nil {
		return
	}

	bt.postOrderRecursive(n.left, traversal)
	bt.postOrderRecursive(n.right, traversal)
	*traversal = append(*traversal, n.key)
}

func (bt *BSTree[T]) levelOrderHelper(root *Node[T], traversal *[]T) {
	var q []*Node[T] // queue
	var n *Node[T]   // temporary node

	q = append(q, root)

	for len(q) != 0 {
		n, q = q[0], q[1:]
		*traversal = append(*traversal, n.key)
		if n.left != nil {
			q = append(q, n.left)
		}
		if n.right != nil {
			q = append(q, n.right)
		}
	}
}
