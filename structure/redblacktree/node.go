package redblacktree

// A representation of a tree node
// Consists of Val (currently only int allowed)
// left, right and parent are pointers to the relevant nodes
// isRed is a representation of the color, made as bool
type RBNode struct {
	Val    int
	left   *RBNode
	right  *RBNode
	parent *RBNode
	isRed  bool
}

// A function for better naming wherever it is used
func (n *RBNode) setBlack() {
	n.isRed = false
}

// Similar to setBlack but for red
func (n *RBNode) setRed() {
	n.isRed = true
}
