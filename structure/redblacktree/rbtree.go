// rbtree.go
//
// description: A struct representing a red black tree
//
// details:
//
// This file contains the tree with a single field - root node
//
// There are multiple methods, public ones are
// Insert - inserts a value
// Search - searches and returns Node pointer
// Delete - deletes a value
// PrintInorder - prints inorder representation - Val and isRed
// PrintPreorder - prints preorder representation - Val and isRed
//
// Implementation is based on Introduction to Algorithms, Cormen et. al., 3rd edition
//
// Wikipedia: https://en.wikipedia.org/wiki/Red%E2%80%93black_tree
// Geeks-for-geeks series: https://www.geeksforgeeks.org/red-black-tree-set-1-introduction-2/
//
// author: VinWare (https://github.com/VinWare)

package redblacktree

// For printing
import (
	"fmt"
)

type RBTree struct {
	Root *Node
}

func (t *RBTree) Insert(Val int) {
	t.insert(Val)
	t.Root.isRed = false
}

// Internal insert function, refer Intro to Algorithms mentioned above
func (t *RBTree) insert(Val int) {
	var newParent *Node = nil
	var currNode = t.Root

	// Find empty node to act as parent
	for currNode != nil {
		newParent = currNode
		if Val < currNode.Val {
			currNode = currNode.left
		} else {
			currNode = currNode.right
		}
	}
	var insNode = newNode(Val, newParent)
	if newParent == nil {
		// Is root node
		t.Root = insNode
		return
	}
	// Is not root node, so must be left or right child of a node
	if Val < newParent.Val {
		newParent.left = insNode
	} else {

		newParent.right = insNode
	}
	currNode = insNode
	for currNode != t.Root && currNode.parent.isRed {
		// By definition, root must be black (specifically Cormen et. al.).
		// Hence within the loop, currNode.parent is not root
		var parentNode = currNode.parent
		var grandNode = parentNode.parent
		// L* case
		if parentNode == grandNode.left {
			var uncle = grandNode.right

			if uncle != nil && uncle.isRed {
				// Case 1
				currNode = case1(parentNode, uncle, grandNode)
			} else {
				if currNode == parentNode.right {
					// Case 2
					t.leftRotate(parentNode)
					currNode = parentNode
				}

				// Case 3
				t.rightRotate(grandNode)
				currNode.parent.isRed = false
				grandNode.isRed = true
			}
		} else {
			// R* case
			var uncle = grandNode.left

			if uncle != nil && uncle.isRed {
				// Case 1
				currNode = case1(parentNode, uncle, grandNode)
			} else {
				if currNode == parentNode.left {
					// Case 2
					t.rightRotate(parentNode)
					currNode = parentNode
				}
				// Case 3
				t.leftRotate(grandNode)
				currNode.parent.isRed = false
				grandNode.isRed = true
			}
		}
	}
}

// Simplest case, make parent level black
func case1(parentNode, uncleNode, grandNode *Node) *Node {
	parentNode.isRed = false
	uncleNode.isRed = false
	grandNode.isRed = true
	return grandNode
}

// A left rotate similar to binary search tree. Coloring is not handled
func (t *RBTree) leftRotate(parentNode *Node) {
	// Swap nodes and parents
	var currNode = parentNode.right
	var tempNode = currNode.left
	currNode.left = parentNode
	parentNode.right = tempNode
	currNode.parent = parentNode.parent
	parentNode.parent = currNode

	if t.isRoot(currNode) {
		t.Root = currNode
		return
	}
	// Set new parent left or right
	if currNode.parent.left == parentNode {
		currNode.parent.left = currNode
	} else {
		currNode.parent.right = currNode
	}
}

// Right rotate, similar to left rotate above
func (t *RBTree) rightRotate(parentNode *Node) {
	var currNode = parentNode.left
	var tempNode = currNode.right
	currNode.right = parentNode
	parentNode.left = tempNode
	currNode.parent = parentNode.parent
	parentNode.parent = currNode

	if currNode.parent == nil {
		t.Root = currNode
		return
	}
	if currNode.parent.left == parentNode {
		currNode.parent.left = currNode
	} else {
		currNode.parent.right = currNode
	}
}

// Simply return new node
func newNode(Val int, parent *Node) *Node {
	return &Node{Val: Val, left: nil, right: nil, parent: parent, isRed: true}
}

// Public function to print preorder
func (t *RBTree) PrintPreorder() {
	var p = preorder(t.Root)
	fmt.Printf("Length %d\n", len(p))
	for _, a := range p {
		fmt.Printf("%d %v\n", a.Val, a.isRed)
	}
}

// Public function to print inorder
func (t *RBTree) PrintInorder() {
	var p = inorder(t.Root)
	fmt.Printf("Length %d\n", len(p))
	for _, a := range p {
		fmt.Printf("%d %v\n", a.Val, a.isRed)
	}
}

// Returns inorder representation
func inorder(currNode *Node) []*Node {
	var ret = []*Node{}
	if currNode == nil {
		return ret
	}

	ret = append(ret, inorder(currNode.left)...)
	ret = append(ret, currNode)
	ret = append(ret, inorder(currNode.right)...)
	return ret
}

// Returns preorder representation
func preorder(currNode *Node) []*Node {
	var ret = []*Node{}
	if currNode == nil {
		return ret
	}
	ret = append(ret, currNode)
	ret = append(ret, preorder(currNode.left)...)
	ret = append(ret, preorder(currNode.right)...)
	return ret
}

// Public search function based on value
func (t *RBTree) Search(val int) *Node {
	return search(t.Root, val)
}

// Internal search function, recursive search
func search(node *Node, val int) *Node {
	if node == nil {
		return nil
	}
	if node.Val == val {
		return node
	}
	if val < node.Val {
		return search(node.left, val)
	}
	return search(node.right, val)
}

// Replaces the node curr with repl
func (t *RBTree) replace(curr *Node, repl *Node) {
	if curr == nil {
		return
	}
	if t.isRoot(curr) {
		t.setRoot(repl)
		return
	}

	if curr == curr.parent.left {
		curr.parent.left = repl
	} else {
		curr.parent.right = repl
	}

	if repl != nil {
		repl.parent = curr.parent
	}
}

// Simple function to test if curr is or should be root
func (t *RBTree) isRoot(curr *Node) bool {
	return curr != nil && curr.parent == nil
}

// Set root (coloring it black is important)
func (t *RBTree) setRoot(curr *Node) {
	t.Root = curr
	if t.Root != nil {
		t.Root.isRed = false
	}
}

// Delete function based on value
func (t *RBTree) Delete(val int) {
	var curr = t.Search(val)
	t.delete(curr)
}

// Internal delete based on node
func (t *RBTree) delete(curr *Node) {
	if curr == nil {
		// Didn't find the node
		return
	}

	var checkIsRed = curr.isRed
	var next *Node

	// First replace the relevant child, moving it up the list
	if curr.left == nil {
		// Only right child/no children
		next = curr.right
		t.replace(curr, curr.right)
	} else if curr.right == nil {
		// Only left child
		next = curr.left
		t.replace(curr, curr.left)
	} else {
		// Have to handle both children
		var minNode = t.min(curr.right)
		checkIsRed = minNode.isRed
		next = minNode.right
		if minNode.parent != curr {
			t.replace(minNode, minNode.right)
			minNode.right = curr.right
			minNode.right.parent = minNode
		}
		t.replace(curr, minNode)
		minNode.left = curr.left
		minNode.left.parent = minNode
		minNode.isRed = curr.isRed
	}
	if !checkIsRed {
		// Certain fix up is needed
		t.fix(next)
	}
}

func (t *RBTree) min(curr *Node) *Node {
	if curr == nil {
		return nil
	}
	for curr.left != nil {
		curr = curr.left
	}
	return curr
}

func (t *RBTree) fix(curr *Node) {
	if curr == nil {
		// Nothing to fix
		return
	}
	for !t.isRoot(curr) && !curr.isRed {
		if curr == curr.parent.left {
			// Left child case
			var sibling = curr.parent.right
			if sibling == nil {
				return
			}
			if sibling.isRed {
				// Recolor and rotate
				sibling.setBlack()
				curr.parent.setRed()
				t.leftRotate(curr.parent)
				sibling = curr.parent.right
			}
			if sibling.left != nil && !sibling.left.isRed && sibling.right != nil && !sibling.right.isRed {
				sibling.setRed()
				curr = curr.parent
			} else {
				if sibling.right != nil && !sibling.right.isRed {
					if sibling.left != nil {
						sibling.left.setBlack()
						sibling.setRed()
						t.rightRotate(sibling)
						sibling = curr.parent.right
					}
					curr.parent.setBlack()
					if sibling != nil {
						sibling.isRed = curr.parent.isRed
					}
					if sibling.right != nil {
						sibling.right.setBlack()
					}
					t.leftRotate(curr.parent)
					curr = t.Root
				}
			}
		} else {
			// Right child case, same as above, but left and right are swapped everywhere
			var sibling = curr.parent.left
			if sibling == nil {
				return
			}
			if sibling.isRed {
				sibling.setBlack()
				curr.parent.setRed()
				t.rightRotate(curr.parent)
				sibling = curr.parent.left
			}
			if sibling.left != nil && !sibling.left.isRed && sibling.right != nil && !sibling.right.isRed {
				sibling.setRed()
				curr = curr.parent
			} else {
				if sibling.left != nil && !sibling.left.isRed {
					if sibling.right != nil {
						sibling.right.setBlack()
						sibling.setRed()
						t.leftRotate(sibling)
						sibling = curr.parent.left
					}
					curr.parent.setBlack()
					if sibling != nil {
						sibling.isRed = curr.parent.isRed
					}
					if sibling.left != nil {
						sibling.left.setBlack()
					}
					t.rightRotate(curr.parent)
					curr = t.Root
				}
			}
		}
	}
	curr.setBlack()
}
