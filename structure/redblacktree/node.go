// node.go
// description: A struct reporesenting tree node
// details:
// This file contains the node with following fields:
// Val is a value, currently only int is allowed
// left and right and parent are self-explanatory pointers
// isRed stores whether a node is red or not
// author: VinWare (https://github.com/VinWare)
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
