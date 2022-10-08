package tree

import "github.com/TheAlgorithms/Go/constraints"

// Node of a binary-search tree
type Node[T constraints.Ordered] struct {
	key   T
	left  *Node[T]
	right *Node[T]
}

// AVLNode of a AVL tree
type AVLNode struct {
	key         int
	height      int
	left, right *AVLNode
}

type Color byte

const (
	Black Color = iota
	Red
)

// RBNode: Node of Red-Black Tree
type RBNode[T constraints.Ordered] struct {
	key    T
	parent *RBNode[T]
	left   *RBNode[T]
	right  *RBNode[T]
	color  Color
}
