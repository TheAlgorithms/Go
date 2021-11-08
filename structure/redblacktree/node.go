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

// Checks whether two nodes can be substituted for each other
// Only value and color is checked as node pointers can be set during actual substitution
func (n *RBNode) EquivalentNode(m *RBNode) bool {
	return n.Val == m.Val && n.isRed == m.isRed
}

//
func CreateNode(Val int, isRed bool) *RBNode {
	return &RBNode{Val: Val, isRed: isRed}
}
