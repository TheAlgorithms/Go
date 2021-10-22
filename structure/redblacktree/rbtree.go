package redblacktree

import (
	"fmt"

	"github.com/TheAlgorithms/Go/structure/queue"
)

type RBTree struct {
	Root *Node
}

func (t *RBTree) Insert(val int) {
	t.intInsert(val)
	t.Root.isRed = false
}

func (t *RBTree) intInsert(val int) {
	var newParent *Node = nil
	var currNode = t.Root
	for currNode != nil {
		newParent = currNode
		if val < currNode.val {
			currNode = currNode.left
		} else {
			currNode = currNode.right
		}
	}
	var insNode = newNode(val, newParent)
	if newParent == nil {
		t.Root = insNode
		return
	}
	if insNode.parent != nil {
		// fmt.Printf("parent of %d is %d\n", insNode.val, insNode.parent.val)
	}
	// fmt.Println("has parent %s", (insNode.parent != nil))
	if val < newParent.val {
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
		// fmt.Printf("has grandparent %s\n", grandNode != nil)
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

func newNode(val int, parent *Node) *Node {
	return &Node{val: val, left: nil, right: nil, parent: parent, isRed: true}
}
func (t *RBTree) Print() {
	// fmt.Println(t.getArray())
	var arr = t.getArray()
	for _, a := range arr {
		if a == nil {
			// fmt.Println()
			continue
		}
		// fmt.Printf("%d %s | ", a.val, a.isRed)
	}
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
	fmt.Printf("%d %s\n", currNode.val, currNode.isRed)
	inorder(currNode.right)
}
func preorder(currNode *Node) {
	if currNode == nil {
		return
	}
	fmt.Printf("%d %s\n", currNode.val, currNode.isRed)
	preorder(currNode.left)
	preorder(currNode.right)
}
func (t *RBTree) getArray() []*Node {
	var res = []*Node{}
	if t.Root == nil {
		return res
	}
	// var queue = []*Node{}
	// var newQueue queue.Queue
	queue.EnQueue(t.Root)
	queue.EnQueue(nil)
	for !queue.IsEmptyQueue() {
		var curr = queue.DeQueue()
		if curr == nil {
			if !queue.IsEmptyQueue() {
				queue.EnQueue(nil)
				res = append(res, nil)

			}
			// fmt.Println("")
		} else {
			var currNode = curr.(*Node)
			res = append(res, currNode)
			if currNode.left != nil {
				queue.EnQueue(currNode.left)
			}
			if currNode.right != nil {
				queue.EnQueue(currNode.right)
			}
			// fmt.Printf("%d %s | ", currNode.val, currNode.isRed)
		}
	}
	return res
}

func (t *RBTree) Search(val int) *Node {
	return intSearch(t.Root, val)
}

func intSearch(node *Node, val int) *Node {
	if node == nil {
		return nil
	}
	if node.val == val {
		return node
	}
	if val < node.val {
		return intSearch(node.left, val)
	}
	return intSearch(node.right, val)
}
