package redblacktree

type Node struct {
	Val    int
	left   *Node
	right  *Node
	parent *Node
	isRed  bool
}

func (n *Node) setBlack() {
	n.isRed = false
}

func (n *Node) setRed() {
	n.isRed = true
}
