// rbtree.go
//
// description: A struct representing a red black tree
//
// details:
//
// A Red Black tree is a self-balancing Binary Search Tree.
// It considers an additional attribute for each node, a color either red or black
// By requiring that no two consecutive nodes in a Depth Traversal can have red color,
// the tree ensures that the heights of all subtrees remain close enough
// to ensure O(n) time search
//
// This file contains the tree with a single field - root node
//
// There are multiple methods, public ones are
// Insert - inserts a value
// Search - searches and returns Node pointer
// Delete - deletes a value
// PrintInOrder - prints inorder representation - Val and isRed
// PrintPreOrder - prints preorder representation - Val and isRed
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
	Root *RBNode
}

func (t *RBTree) Insert(Val int) {
	t.insert(Val)
	t.Root.setBlack()
}

// Simply search the tree for the new parent of the new node
// Similar to ordinary BST
func (t *RBTree) findNewParent(Val int) *RBNode {
	var newParent *RBNode = nil
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
	return newParent
}

// Internal insert function, refer Intro to Algorithms mentioned above
func (t *RBTree) insert(Val int) {
	var newParent = t.findNewParent(Val)
	var insNode = newNode(Val, newParent)
	if newParent == nil {
		// Is root node
		t.setRoot(insNode)
		return
	}
	// Is not root node, so must be left or right child of a node
	if Val < newParent.Val {
		newParent.left = insNode
	} else {
		newParent.right = insNode
	}

	t.rebalanceTree(insNode)
}

func (t *RBTree) rebalanceTree(curr *RBNode) {
	for !t.isRoot(curr) && curr.parent.isRed {
		// As root is black (Cormen style), parent is not root
		// Current node is red and so is parent, so must climb and recolor
		curr = t.climbTreeColoring(curr)
	}
}
func (t *RBTree) climbTreeColoring(curr *RBNode) *RBNode {
	var parent = curr.parent
	var grand = parent.parent
	var uncle = getSibling(parent)
	if uncle != nil && uncle.isRed {
		// Case 1
		return case1(parent, uncle, grand)
	}
	if isTriangle(curr, parent, grand) {
		// 		grand			grand
		//      /					\
		//	   /					 \
		//   parent		OR			parent
		//     \					 /
		//		\					/
		//      curr			curr
		//
		// First rotate the curr and parent to ensure a straight line
		t.rotate(curr, parent)

		// Swap curr and parent
		var tempNode = curr
		curr = parent
		parent = tempNode
	}

	// 			grand		grand
	//   	   /				\
	//		  /				 	 \
	//   	parent		OR		parent
	//     /				 		\
	//	  /							 \
	//   curr						curr
	//
	// Rotate grandparent node
	t.rotate(parent, grand)

	// Ensure that levels are colored to maintain rules
	curr.parent.setBlack() // Since curr is red, it's parent should be black
	grand.setRed()         // Since grand is moved down, it is colored to red
	return curr
}

func isTriangle(curr *RBNode, parent *RBNode, grand *RBNode) bool {
	return grand.left == parent && parent.right == curr || grand.right == parent && parent.left == curr
}

func getSibling(curr *RBNode) *RBNode {
	if curr == curr.parent.right {
		return curr.parent.left
	}
	return curr.parent.right
}

// Simplest case, make parent level black and grandparent level red
func case1(parentNode, uncleNode, grandNode *RBNode) *RBNode {
	parentNode.setBlack()
	uncleNode.setBlack()
	grandNode.setRed()
	return grandNode
}

// A left rotate similar to binary search tree. Colors remain same
func (t *RBTree) leftRotate(parentNode *RBNode) {
	// Swap nodes and parents
	t.rotate(parentNode.right, parentNode)
}

func (t *RBTree) setParents(currNode *RBNode, parentNode *RBNode) {
	var grandNode = parentNode.parent
	currNode.parent = grandNode
	parentNode.parent = currNode

	if t.isRoot(currNode) {
		t.setRoot(currNode)
		return
	}
	// Set new parent left or right
	if grandNode.left == parentNode {
		grandNode.left = currNode
	} else {
		grandNode.right = currNode
	}
}

// Right rotate, similar to left rotate above
func (t *RBTree) rightRotate(parentNode *RBNode) {
	// var currNode = parentNode.left
	t.rotate(parentNode.left, parentNode)
}

func (t *RBTree) rotate(currNode *RBNode, parentNode *RBNode) {
	if parentNode == nil || currNode == nil {
		return
	}
	if parentNode.right == currNode {
		// leftRotate
		var shiftNode = currNode.left
		currNode.left = parentNode
		parentNode.right = shiftNode
	} else {
		// rightRotate
		var shiftNode = currNode.right
		currNode.right = parentNode
		parentNode.left = shiftNode
	}
	t.setParents(currNode, parentNode)
}

// Simply return new node
func newNode(Val int, parent *RBNode) *RBNode {
	return &RBNode{Val: Val, left: nil, right: nil, parent: parent, isRed: true}
}

// Public function to print preorder
func (t *RBTree) PrintPreOrder() {
	var p = preOrder(t.Root)
	fmt.Printf("Length %d\n", len(p))
	for _, a := range p {
		fmt.Printf("%d %v\n", a.Val, a.isRed)
	}
}

// Public function to print inorder
func (t *RBTree) PrintInOrder() {
	var p = inOrder(t.Root)
	fmt.Printf("Length %d\n", len(p))
	for _, a := range p {
		fmt.Printf("%d %v\n", a.Val, a.isRed)
	}
}

// Returns inorder representation
func inOrder(currNode *RBNode) []*RBNode {
	var ret = []*RBNode{}
	if currNode == nil {
		return ret
	}

	ret = append(ret, inOrder(currNode.left)...)
	ret = append(ret, currNode)
	ret = append(ret, inOrder(currNode.right)...)
	return ret
}

// Returns preorder representation
func preOrder(currNode *RBNode) []*RBNode {
	var ret = []*RBNode{}
	if currNode == nil {
		return ret
	}
	ret = append(ret, currNode)
	ret = append(ret, preOrder(currNode.left)...)
	ret = append(ret, preOrder(currNode.right)...)
	return ret
}

// Public search function based on value
func (t *RBTree) Search(val int) *RBNode {
	return SearchFromNode(t.Root, val)
}

// Internal SearchFromNode function, recursive SearchFromNode
func SearchFromNode(node *RBNode, val int) *RBNode {
	if node == nil {
		return nil
	}
	if node.Val == val {
		return node
	}
	if val < node.Val {
		return SearchFromNode(node.left, val)
	}
	return SearchFromNode(node.right, val)
}

// Replaces the node curr with repl
func (t *RBTree) replace(curr *RBNode, repl *RBNode) {
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
func (t *RBTree) isRoot(curr *RBNode) bool {
	return curr != nil && curr.parent == nil
}

// Set root (coloring it black is important)
func (t *RBTree) setRoot(curr *RBNode) {
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
func (t *RBTree) delete(curr *RBNode) {
	if curr == nil {
		// Didn't find the node
		return
	}

	var checkIsRed = curr.isRed
	var next *RBNode

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

func (t *RBTree) min(curr *RBNode) *RBNode {
	if curr == nil {
		return nil
	}
	for curr.left != nil {
		curr = curr.left
	}
	return curr
}

func (t *RBTree) fix(curr *RBNode) {
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
