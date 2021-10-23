// Package trie provides Trie data structures in golang.
//
// Wikipedia: https://en.wikipedia.org/wiki/Trie
package trie

// Node represents each node in Trie.
type Node struct {
	children map[rune]*Node // map children nodes
	isLeaf   bool           // current node value
}

// NewNode creates a new Trie node with initialized
// children map.
func NewNode() *Node {
	n := &Node{}
	n.children = make(map[rune]*Node)
	n.isLeaf = false
	return n
}

// insert a single word at a Trie node.
func (n *Node) insert(s string) {
	curr := n
	for _, c := range s {
		next, ok := curr.children[c]
		if !ok {
			next = NewNode()
			curr.children[c] = next
		}
		curr = next
	}
	curr.isLeaf = true
}

// Insert zero, one or more words at a Trie node.
func (n *Node) Insert(s ...string) {
	for _, ss := range s {
		n.insert(ss)
	}
}

// Find  words at a Trie node.
func (n *Node) Find(s string) bool {
	next, ok := n, false
	for _, c := range s {
		next, ok = next.children[c]
		if !ok {
			return false
		}
	}
	return next.isLeaf
}

// Capacity returns the number of nodes in the Trie
func (n *Node) Capacity() int {
	r := 0
	for _, c := range n.children {
		r += c.Capacity()
	}
	return 1 + r
}

// Size returns the number of words in the Trie
func (n *Node) Size() int {
	r := 0
	for _, c := range n.children {
		r += c.Size()
	}
	if n.isLeaf {
		r++
	}
	return r
}

// remove lazily a word from the Trie node, no node is actually removed.
func (n *Node) remove(s string) {
	if len(s) == 0 {
		return
	}
	next, ok := n, false
	for _, c := range s {
		next, ok = next.children[c]
		if !ok {
			// word cannot be found - we're done !
			return
		}
	}
	next.isLeaf = false
}

// Remove zero, one or more words lazily from the Trie, no node is actually removed.
func (n *Node) Remove(s ...string) {
	for _, ss := range s {
		n.remove(ss)
	}
}

// Compact will remove unecessay nodes, reducing the capacity, returning true if node n itself should be removed.
func (n *Node) Compact() (remove bool) {

	for r, c := range n.children {
		if c.Compact() {
			delete(n.children, r)
		}
	}
	return !n.isLeaf && len(n.children) == 0
}
