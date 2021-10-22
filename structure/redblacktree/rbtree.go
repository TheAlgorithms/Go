package redblacktree

type RBTree struct {
	Root *Node
}

func (t *RBTree) Insert(val int) {
	t.intInsert(val)
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
		// L* case
		if parentNode == grandNode.left {
			var uncle = grandNode.right

			if uncle.isRed {
				// Case 1
				currNode = case1(parentNode, uncle, grandNode)
			} else {
				if currNode == parentNode.right {
					// Case 2
					leftRotate(parentNode)
					currNode = parentNode
					// var temp = currNode.left
					// currNode.left = parentNode
					// parentNode.right = temp
					// currNode.parent = parentNode.parent
					// parentNode.parent = currNode
				}

				// Case 3
				rightRotate(grandNode)
				// var temp = parentNode.right
				// parentNode.right = grandNode
				// grandNode.left = temp
				// parentNode.parent = grandNode.parent
				// grandNode.parent = parentNode
				currNode.parent.isRed = false
				grandNode.isRed = true
			}
		} else {
			var uncle = grandNode.left

			if uncle.isRed {
				currNode = case1(parentNode, uncle, grandNode)
			} else {
				if currNode == parentNode.left {
					rightRotate(parentNode)
					currNode = parentNode
				}
				leftRotate(grandNode)
				currNode.parent.isRed = false
				grandNode.isRed = true
			}
		}
	}
	t.Root.isRed = false
}

func case1(parentNode, uncleNode, grandNode *Node) *Node {
	parentNode.isRed = false
	uncleNode.isRed = false
	grandNode.isRed = true
	return grandNode
}

func leftRotate(parentNode *Node) {
	var currNode = parentNode.right
	var tempNode = currNode.left
	currNode.left = parentNode
	parentNode.right = tempNode
	currNode.parent = parentNode.parent
	parentNode.parent = currNode

	if currNode.parent.left == parentNode {
		currNode.parent.left = currNode
	} else {
		currNode.parent.right = currNode
	}
}

func rightRotate(parentNode *Node) {
	var currNode = parentNode.left
	var tempNode = currNode.right
	currNode.right = parentNode
	parentNode.left = tempNode
	currNode.parent = parentNode.parent
	parentNode.parent = currNode
	// var temp =
	if currNode.parent.left == parentNode {
		currNode.parent.left = currNode
	} else {
		currNode.parent.right = currNode
	}
}

func newNode(val int, parent *Node) *Node {
	return &Node{val: val, left: nil, right: nil, parent: nil, isRed: true}
}
