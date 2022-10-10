package tree

import "github.com/TheAlgorithms/Go/constraints"

// Node of a binary-search tree
type Node[T constraints.Ordered] struct {
	Key   T
	Left  *Node[T]
	Right *Node[T]
}

// AVLNode of a AVL tree
type AVLNode struct {
	Key         int
	Height      int
	Left, Right *AVLNode
}

type Color byte

const (
	Black Color = iota
	Red
)

// RBNode: Node of Red-Black Tree
type RBNode[T constraints.Ordered] struct {
	Key    T
	Parent *RBNode[T]
	Left   *RBNode[T]
	Right  *RBNode[T]
	Color  Color
}
