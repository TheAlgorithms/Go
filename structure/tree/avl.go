// Package avl is a Adelson-Velskii and Landis tree implemnation
// avl is self-balancing tree, i.e for all node in a tree, Height difference
// between its Left and Right child will not exceed 1
// more information : https://en.wikipedia.org/wiki/AVL_tree
package tree

// NewAVLTree create a new AVL tree
func NewAVLTree() *AVLNode {
	return nil
}

// Get : return node with given Key
func Get(root *AVLNode, key int) *AVLNode {
	if root == nil {
		return nil
	}
	if root.Key == key {
		return root
	} else if root.Key < key {
		root = root.Right
	} else {
		root = root.Left
	}
	return Get(root, key)
}

// Insert a new item
func Insert(root **AVLNode, key int) {
	if *root == nil {
		*root = &AVLNode{
			Key:    key,
			Height: 1,
		}
		return
	}
	if (*root).Key < key {
		Insert(&(*root).Right, key)
	} else if (*root).Key > key {
		Insert(&(*root).Left, key)
	}

	// update Height
	(*root).Height = height(*root)

	bFactor := balanceFactor(*root)

	if bFactor == 2 { // L
		bFactor = balanceFactor((*root).Left)
		if bFactor == 1 { // LL
			llRotation(root)
		} else if bFactor == -1 { // LR
			lrRotation(root)
		}
	} else if bFactor == -2 { // R
		bFactor = balanceFactor((*root).Right)
		if bFactor == 1 { // RL
			rlRotation(root)
		} else if bFactor == -1 { // RR
			rrRotation(root)
		}
	}
}

// Delete : remove given Key from the tree
func Delete(root **AVLNode, key int) {
	if root == nil {
		return
	}
	if (*root).Key < key {
		Delete(&(*root).Right, key)
	} else if (*root).Key > key {
		Delete(&(*root).Left, key)
	} else {
		// 3 cases
		// 1. No Child
		// 2. With One Child
		// 3. With Two Child
		if (*root).Left == nil && (*root).Right == nil {
			*root = nil
		} else if (*root).Left == nil {
			*root = (*root).Right
		} else if (*root).Right == nil {
			*root = (*root).Left
		} else {
			minVal := min((*root).Right)
			(*root).Key = minVal
			Delete(root, minVal)
		}
		return
	}

	// update Height
	(*root).Height = height(*root)

	bFactor := balanceFactor(*root)

	if bFactor == 2 { // L
		switch balanceFactor((*root).Left) {
		case 1: // LL
			llRotation(root)
		case -1: // LR
			lrRotation(root)
		case 0: //  LL OR LR
			llRotation(root)
		}
	} else if bFactor == -2 { // L
		switch balanceFactor((*root).Right) {
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
	b := (*root).Left
	br := b.Right
	b.Right = *root
	(*root).Left = br
	(*root).Height = height(*root)
	b.Height = height(b)
	*root = b
}
func lrRotation(root **AVLNode) {
	c := (*root).Left.Right
	cl := c.Left
	cr := c.Right

	c.Left = (*root).Left
	c.Right = (*root)
	c.Left.Right = cl

	(*root).Left = cr

	(*root).Height = height(*root)
	c.Left.Height = height(c.Left)
	c.Height = height(c)

	*root = c

}
func rrRotation(root **AVLNode) {
	b := (*root).Right
	bl := b.Left
	b.Left = *root

	(*root).Right = bl
	(*root).Height = height(*root)
	b.Height = height(b)
	*root = b

}
func rlRotation(root **AVLNode) {
	c := (*root).Right.Left
	cl := c.Left
	cr := c.Right

	c.Right = (*root).Right
	c.Right.Left = cr
	c.Left = *root
	(*root).Right = cl

	(*root).Height = height(*root)
	c.Right.Height = height(c.Right)
	c.Height = height(c)
	*root = c
}

// balanceFactor : -ve balance factor means subtree Root is heavy toward Left
// and +ve balance factor means subtree Root is heavy toward Right side
func balanceFactor(root *AVLNode) int {
	var leftHeight, rightHeight int
	if root.Left != nil {
		leftHeight = root.Left.Height
	}
	if root.Right != nil {
		rightHeight = root.Right.Height
	}
	return leftHeight - rightHeight
}

func height(root *AVLNode) int {
	if root == nil {
		return 0
	}
	var leftHeight, rightHeight int
	if root.Left != nil {
		leftHeight = root.Left.Height
	}
	if root.Right != nil {
		rightHeight = root.Right.Height
	}
	max := leftHeight
	if rightHeight > leftHeight {
		max = rightHeight
	}
	return 1 + max
}

func min(root *AVLNode) int {
	if root.Left == nil {
		return root.Key
	}
	return min(root.Left)
}
