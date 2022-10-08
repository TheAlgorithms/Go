// Package avl is a Adelson-Velskii and Landis tree implemnation
// avl is self-balancing tree, i.e for all node in a tree, height difference
// between its left and right child will not exceed 1
// more information : https://en.wikipedia.org/wiki/AVL_tree
package tree

// NewAVLTree create a new AVL tree
func NewAVLTree() *AVLNode {
	return nil
}

// Get : return node with given key
func Get(root *AVLNode, key int) *AVLNode {
	if root == nil {
		return nil
	}
	if root.key == key {
		return root
	} else if root.key < key {
		root = root.right
	} else {
		root = root.left
	}
	return Get(root, key)
}

// Insert a new item
func Insert(root **AVLNode, key int) {
	if *root == nil {
		*root = &AVLNode{
			key:    key,
			height: 1,
		}
		return
	}
	if (*root).key < key {
		Insert(&(*root).right, key)
	} else if (*root).key > key {
		Insert(&(*root).left, key)
	}

	// update height
	(*root).height = height(*root)

	bFactor := balanceFactor(*root)

	if bFactor == 2 { // L
		bFactor = balanceFactor((*root).left)
		if bFactor == 1 { // LL
			llRotation(root)
		} else if bFactor == -1 { // LR
			lrRotation(root)
		}
	} else if bFactor == -2 { // R
		bFactor = balanceFactor((*root).right)
		if bFactor == 1 { // RL
			rlRotation(root)
		} else if bFactor == -1 { // RR
			rrRotation(root)
		}
	}
}

// Delete : remove given key from the tree
func Delete(root **AVLNode, key int) {
	if root == nil {
		return
	}
	if (*root).key < key {
		Delete(&(*root).right, key)
	} else if (*root).key > key {
		Delete(&(*root).left, key)
	} else {
		// 3 cases
		// 1. No Child
		// 2. With One Child
		// 3. With Two Child
		if (*root).left == nil && (*root).right == nil {
			*root = nil
		} else if (*root).left == nil {
			*root = (*root).right
		} else if (*root).right == nil {
			*root = (*root).left
		} else {
			minVal := min((*root).right)
			(*root).key = minVal
			Delete(root, minVal)
		}
		return
	}

	// update height
	(*root).height = height(*root)

	bFactor := balanceFactor(*root)

	if bFactor == 2 { // L
		switch balanceFactor((*root).left) {
		case 1: // LL
			llRotation(root)
		case -1: // LR
			lrRotation(root)
		case 0: //  LL OR LR
			llRotation(root)
		}
	} else if bFactor == -2 { // L
		switch balanceFactor((*root).right) {
		case 1: // RL
			rlRotation(root)
		case -1: // RR
			rrRotation(root)
		case 0: // RL OR RR
			rrRotation(root)
		}
	}
}

// rotations
// 1. LL
// 2. LR
// 3. RR
// 4. RL
func llRotation(root **AVLNode) {
	b := (*root).left
	br := b.right
	b.right = *root
	(*root).left = br
	(*root).height = height(*root)
	b.height = height(b)
	*root = b
}
func lrRotation(root **AVLNode) {
	c := (*root).left.right
	cl := c.left
	cr := c.right

	c.left = (*root).left
	c.right = (*root)
	c.left.right = cl

	(*root).left = cr

	(*root).height = height(*root)
	c.left.height = height(c.left)
	c.height = height(c)

	*root = c

}
func rrRotation(root **AVLNode) {
	b := (*root).right
	bl := b.left
	b.left = *root

	(*root).right = bl
	(*root).height = height(*root)
	b.height = height(b)
	*root = b

}
func rlRotation(root **AVLNode) {
	c := (*root).right.left
	cl := c.left
	cr := c.right

	c.right = (*root).right
	c.right.left = cr
	c.left = *root
	(*root).right = cl

	(*root).height = height(*root)
	c.right.height = height(c.right)
	c.height = height(c)
	*root = c
}

// balanceFactor : -ve balance factor means subtree root is heavy toward left
// and +ve balance factor means subtree root is heavy toward right side
func balanceFactor(root *AVLNode) int {
	var leftHeight, rightHeight int
	if root.left != nil {
		leftHeight = root.left.height
	}
	if root.right != nil {
		rightHeight = root.right.height
	}
	return leftHeight - rightHeight
}

func height(root *AVLNode) int {
	if root == nil {
		return 0
	}
	var leftHeight, rightHeight int
	if root.left != nil {
		leftHeight = root.left.height
	}
	if root.right != nil {
		rightHeight = root.right.height
	}
	max := leftHeight
	if rightHeight > leftHeight {
		max = rightHeight
	}
	return 1 + max
}

func min(root *AVLNode) int {
	if root.left == nil {
		return root.key
	}
	return min(root.left)
}
