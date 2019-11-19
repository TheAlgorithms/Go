package main

import "fmt"

// definition of a bst node
type node struct {
	val   int
	left  *node
	right *node
}

// definition of a node
type btree struct {
	root *node
}

// allocating a new node
func newNode(val int) *node {
	return &node{val, nil, nil}
}

// insert nodes into a binary search tree
func insert(root *node, val int) *node {
	if root == nil {
		return newNode(val)
	}
	if val < root.val {
		root.left = insert(root.left, val)
	} else {
		root.right = insert(root.right, val)
	}
	return root
}

// inorder traversal algorithm
// Copies the elements of the bst to the array in sorted order
func inorderCopy(n *node, array []int, index *int) {
	if n != nil {
		inorderCopy(n.left, array, index)
		array[*index] = n.val
		*index++
		inorderCopy(n.right, array, index)
	}
}

func treesort(array []int, tree *btree) {
	// build the binary search tree
	for _, element := range array {
		tree.root = insert(tree.root, element)
	}
	index := 0
	// perform inorder traversal to get the elements in sorted order
	inorderCopy(tree.root, array, &index)
}

// tester
func main() {
	tree := &btree{nil}
	numbers := []int{5, 4, 3, 2, 1, -1, 0}
	fmt.Println("numbers : ", numbers)
	treesort(numbers, tree)
	fmt.Println("numbers : ", numbers)
}
