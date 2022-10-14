// Binary-Search tree is the tree, with the key of each internal node
// being greater than keys in the respective node's left subtree
// and less than the ones in its right subtree.

// For more details check out those links below here:
// Wikipedia article: https://en.wikipedia.org/wiki/Binary_search_tree
// authors [guuzaa](https://github.com/guuzaa)
package tree

import (
	"github.com/TheAlgorithms/Go/constraints"
	"github.com/TheAlgorithms/Go/math/max"
)

type Color byte

const (
	Red Color = iota
	Black
)

// Node of a binary tree
type Node[T constraints.Ordered] struct {
	Key    T
	Parent *Node[T] // for Red-Black Tree
	Left   *Node[T]
	Right  *Node[T]
	Color  Color // for Red-Black Tree
	Height int   // for AVL Tree
}

// binaryTree is a base-struct for BSTree, AVLTree, RBTree, etc.
// Note: to avoid instantiation, we make the base struct un-exported.
type binaryTree[T constraints.Ordered] struct {
	Root *Node[T]
	NIL  *Node[T] // NIL denotes the leaf node of Red-Black Tree
}

// Get a Node from the binary-search Tree
func (t *binaryTree[T]) Get(key T) (*Node[T], bool) {
	return t.searchTreeHelper(t.Root, key)
}

// Determines the tree has the node of Key
func (t *binaryTree[T]) Has(key T) bool {
	_, ok := t.searchTreeHelper(t.Root, key)
	return ok
}

// Traverses the tree in the following order Root --> Left --> Right
func (t *binaryTree[T]) PreOrder() []T {
	traversal := make([]T, 0)
	t.preOrderRecursive(t.Root, &traversal)
	return traversal
}

// Traverses the tree in the following order Left --> Root --> Right
func (t *binaryTree[T]) InOrder() []T {
	return t.inOrderHelper(t.Root)
}

// Traverses the tree in the following order Left --> Right --> Root
func (t *binaryTree[T]) PostOrder() []T {
	traversal := make([]T, 0)
	t.postOrderRecursive(t.Root, &traversal)
	return traversal
}

// Depth returns the calculated depth of a binary search tree
func (t *binaryTree[T]) Depth() int {
	return t.calculateDepth(t.Root, 0)
}

// Returns the Max value of the tree
func (t *binaryTree[T]) Max() (T, bool) {
	ret := t.maximum(t.Root)
	if t.isNil(ret) {
		return t.NIL.Key, false
	}

	return ret.Key, true
}

// Return the Min value of the tree
func (t *binaryTree[T]) Min() (T, bool) {
	ret := t.minimum(t.Root)
	if t.isNil(ret) {
		return t.NIL.Key, false
	}

	return ret.Key, true
}

// LevelOrder returns the level order traversal of the tree
func (t *binaryTree[T]) LevelOrder() []T {
	traversal := make([]T, 0)
	t.levelOrderHelper(t.Root, &traversal)
	return traversal
}

// Determines node is a leaf node
func (t *binaryTree[T]) isNil(node *Node[T]) bool {
	return node == t.NIL
}

func (t *binaryTree[T]) searchTreeHelper(node *Node[T], key T) (*Node[T], bool) {
	if node == nil || t.isNil(node) {
		return node, false
	}

	if key == node.Key {
		return node, true
	}

	if key < node.Key {
		return t.searchTreeHelper(node.Left, key)
	}
	return t.searchTreeHelper(node.Right, key)
}

// The iterative inorder;
// The recursive way is similar to the preOrderRecursive
func (t *binaryTree[T]) inOrderHelper(node *Node[T]) []T {
	var stack []*Node[T]
	var ret []T

	for !t.isNil(node) || len(stack) > 0 {
		for !t.isNil(node) {
			stack = append(stack, node)
			node = node.Left
		}

		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		ret = append(ret, node.Key)
		node = node.Right
	}

	return ret
}

func (t *binaryTree[T]) preOrderRecursive(n *Node[T], traversal *[]T) {
	if t.isNil(n) {
		return
	}

	*traversal = append(*traversal, n.Key)
	t.preOrderRecursive(n.Left, traversal)
	t.preOrderRecursive(n.Right, traversal)
}

func (t *binaryTree[T]) postOrderRecursive(n *Node[T], traversal *[]T) {
	if t.isNil(n) {
		return
	}

	t.postOrderRecursive(n.Left, traversal)
	t.postOrderRecursive(n.Right, traversal)
	*traversal = append(*traversal, n.Key)
}

func (t *binaryTree[T]) calculateDepth(n *Node[T], depth int) int {
	if t.isNil(n) {
		return depth
	}

	return max.Int(t.calculateDepth(n.Left, depth+1), t.calculateDepth(n.Right, depth+1))
}

// Returns the minimum value of node of the tree
func (t *binaryTree[T]) minimum(node *Node[T]) *Node[T] {
	if t.isNil(node) {
		return node
	}

	for !t.isNil(node.Left) {
		node = node.Left
	}
	return node
}

// Returns the maximum value of node of the tree
func (t *binaryTree[T]) maximum(node *Node[T]) *Node[T] {
	if t.isNil(node) {
		return node
	}

	for !t.isNil(node.Right) {
		node = node.Right
	}
	return node
}

func (t *binaryTree[T]) levelOrderHelper(root *Node[T], traversal *[]T) {
	var q []*Node[T] // queue
	var tmp *Node[T]

	q = append(q, root)

	for len(q) != 0 {
		tmp, q = q[0], q[1:]
		*traversal = append(*traversal, tmp.Key)
		if !t.isNil(tmp.Left) {
			q = append(q, tmp.Left)
		}

		if !t.isNil(tmp.Right) {
			q = append(q, tmp.Right)
		}
	}
}
