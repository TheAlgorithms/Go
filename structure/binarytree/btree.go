package binarytree

// BTree Returns a binary tree structure which contains only a root Node
type BTree struct {
	Root *Node
}

// calculateDepth helper function for BTree's depth()
func calculateDepth(n *Node, depth int) int {
	if n == nil {
		return depth
	}
	return Max(calculateDepth(n.left, depth+1), calculateDepth(n.right, depth+1))
}

// Insert a value in the BTree
func Insert(root *Node, val int) *Node {
	if root == nil {
		return NewNode(val)
	}
	if val < root.val {
		root.left = Insert(root.left, val)
	} else {
		root.right = Insert(root.right, val)
	}
	return root
}

// Depth returns the calculated depth of a binary tree
func (t *BTree) Depth() int {
	return calculateDepth(t.Root, 0)
}

// InOrderSuccessor Goes to the left
func InOrderSuccessor(root *Node) *Node {
	cur := root
	for cur.left != nil {
		cur = cur.left
	}
	return cur
}

// BstDelete removes the node
func BstDelete(root *Node, val int) *Node {
	if root == nil {
		return nil
	}
	if val < root.val {
		root.left = BstDelete(root.left, val)
	} else if val > root.val {
		root.right = BstDelete(root.right, val)
	} else {
		// this is the node to delete
		// node with one child
		if root.left == nil {
			return root.right
		} else if root.right == nil {
			return root.left
		} else {
			n := root.right
			d := InOrderSuccessor(n)
			d.left = root.left
			return root.right
		}
	}
	return root
}

// InOrder add's children to a node in order left first then right recursively
func inOrderRecursive(n *Node, traversal *[]int) {
	if n != nil {
		inOrderRecursive(n.left, traversal)
		*traversal = append(*traversal, n.val)
		inOrderRecursive(n.right, traversal)
	}
}

func InOrder(root *Node) []int {
	traversal := make([]int, 0)
	inOrderRecursive(root, &traversal)
	return traversal
}

// PreOrder Preorder
func preOrderRecursive(n *Node, traversal *[]int) {
	if n == nil {
		return
	}
	*traversal = append(*traversal, n.val)
	preOrderRecursive(n.left, traversal)
	preOrderRecursive(n.right, traversal)
}

func PreOrder(root *Node) []int {
	traversal := make([]int, 0)
	preOrderRecursive(root, &traversal)
	return traversal
}

// PostOrder PostOrder
func postOrderRecursive(n *Node, traversal *[]int) {
	if n == nil {
		return
	}
	postOrderRecursive(n.left, traversal)
	postOrderRecursive(n.right, traversal)
	*traversal = append(*traversal, n.val)
}

func PostOrder(root *Node) []int {
	traversal := make([]int, 0)
	postOrderRecursive(root, &traversal)
	return traversal
}

// LevelOrder LevelOrder
func levelOrderRecursive(root *Node, traversal *[]int) {
	var q []*Node // queue
	var n *Node   // temporary node

	q = append(q, root)

	for len(q) != 0 {
		n, q = q[0], q[1:]
		*traversal = append(*traversal, n.val)
		if n.left != nil {
			q = append(q, n.left)
		}
		if n.right != nil {
			q = append(q, n.right)
		}
	}
}

func LevelOrder(root *Node) []int {
	traversal := make([]int, 0)
	levelOrderRecursive(root, &traversal)
	return traversal
}

// AccessNodesByLayer Function that access nodes layer by layer instead of printing the results as one line.
func AccessNodesByLayer(root *Node) [][]int {
	var res [][]int
	if root == nil {
		return res
	}
	var q []*Node
	var n *Node
	var idx = 0
	q = append(q, root)

	for len(q) != 0 {
		res = append(res, []int{})
		qLen := len(q)
		for i := 0; i < qLen; i++ {
			n, q = q[0], q[1:]
			res[idx] = append(res[idx], n.val)
			if n.left != nil {
				q = append(q, n.left)
			}
			if n.right != nil {
				q = append(q, n.right)
			}
		}
		idx++
	}
	return res
}

// Max Function that returns max of two numbers - possibly already declared.
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
