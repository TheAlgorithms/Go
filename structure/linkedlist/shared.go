package linkedlist

// Node Structure representing the linkedlist node.
// This node is shared across different implementations.
type Node struct {
	Val  interface{}
	Prev *Node
	Next *Node
}

// Create new node.
func NewNode(val interface{}) *Node {
	return &Node{val, nil, nil}
}
