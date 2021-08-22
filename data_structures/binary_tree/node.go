package binarytree

// Node structure for Nodes
type Node struct {
	val   int
	left  *Node
	right *Node
}

// NewNode Returns a new pointer to an empty Node
func NewNode(val int) *Node {
	return &Node{val, nil, nil}
}
