package redblacktree

type Node struct {
	val    int
	left   *Node
	right  *Node
	parent *Node
	isRed  bool
}
