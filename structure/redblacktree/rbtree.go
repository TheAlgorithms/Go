package redblacktree

import (
	"fmt"
)

type RBTree struct {
	Root *Node
}

func (t *RBTree) Insert(Val int) {
	t.intInsert(Val)
	t.Root.isRed = false
}

func (t *RBTree) intInsert(Val int) {
	var newParent *Node = nil
	var currNode = t.Root
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
		t.Root = insNode
		return
	}
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

func case1(parentNode, uncleNode, grandNode *Node) *Node {
	parentNode.isRed = false
	uncleNode.isRed = false
	grandNode.isRed = true
	return grandNode
}

func (t *RBTree) leftRotate(parentNode *Node) {
	var currNode = parentNode.right
	var tempNode = currNode.left
	currNode.left = parentNode
	parentNode.right = tempNode
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

func newNode(Val int, parent *Node) *Node {
	return &Node{Val: Val, left: nil, right: nil, parent: parent, isRed: true}
}
func (t *RBTree) PrintPreorder() {
	// var currNode = t.Root
	preorder(t.Root)
}
func (t *RBTree) PrintInorder() {
	inorder(t.Root)
}
func inorder(currNode *Node) {
	if currNode == nil {
		return
	}

	inorder(currNode.left)
	fmt.Printf("%d %v\n", currNode.Val, currNode.isRed)
	inorder(currNode.right)
}
func preorder(currNode *Node) {
	if currNode == nil {
		return
	}
	fmt.Printf("%d %v\n", currNode.Val, currNode.isRed)
	preorder(currNode.left)
	preorder(currNode.right)
}

func (t *RBTree) Search(Val int) *Node {
	return intSearch(t.Root, Val)
}

func intSearch(node *Node, Val int) *Node {
	if node == nil {
		return nil
	}
	if node.Val == Val {
		return node
	}
	if Val < node.Val {
		return intSearch(node.left, Val)
	}
	return intSearch(node.right, Val)
}

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

func (t *RBTree) isRoot(curr *Node) bool {
	return curr != nil && curr.parent == nil
}
func (t *RBTree) setRoot(curr *Node) {
	t.Root = curr
	if t.Root != nil {
		t.Root.isRed = false
	}
}

func (t *RBTree) Delete(val int) {
	var curr = t.Search(val)
	t.delete(curr)
}
func (t *RBTree) delete(curr *Node) {
	if curr == nil {
		return
	}
	var checkIsRed = curr.isRed
	var next *Node
	if curr.left == nil {
		next = curr.right
		t.replace(curr, curr.right)
	} else if curr.right == nil {
		next = curr.left
		t.replace(curr, curr.left)
	} else {
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
		return
	}
	for !t.isRoot(curr) && !curr.isRed {
		if curr == curr.parent.left {
			var sibling = curr.parent.right
			if sibling == nil {
				return
			}
			if sibling.isRed {
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
