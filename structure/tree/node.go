package tree

import "github.com/TheAlgorithms/Go/constraints"

type Node struct {
	key   int
	left  *Node
	right *Node
}

// NewNode Returns a new pointer to an empty Node
func NewNode(val int) *Node {
	return &Node{key: val, left: nil, right: nil}
}

type Color byte

const (
	Black Color = iota
	Red
)

// Node of Red-Black Tree
type RBNode[T constraints.Ordered] struct {
	key    T
	parent *RBNode[T]
	left   *RBNode[T]
	right  *RBNode[T]
	color  Color
}
