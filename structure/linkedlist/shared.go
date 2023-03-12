package linkedlist

// Node Structure representing the linkedlist node.
// This node is shared across different implementations.
type Node[T any] struct {
	Val  T
	Prev *Node[T]
	Next *Node[T]
}

// Create new node.
func NewNode[T any](val T) *Node[T] {
	return &Node[T]{val, nil, nil}
}
