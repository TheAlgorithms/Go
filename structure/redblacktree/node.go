package redblacktree

// A representation of a tree node
// Consists of Val (currently only int allowed)
// left, right and parent are pointers to the relevant nodes
// isRed is a representation of the color, made as bool
type Node struct {
	Val    int
	left   *Node
	right  *Node
	parent *Node
	isRed  bool
}

// A function for better naming wherever it is used
func (n *Node) setBlack() {
	n.isRed = false
}

// Similar to setBlack but for red
func (n *Node) setRed() {
	n.isRed = true
}
