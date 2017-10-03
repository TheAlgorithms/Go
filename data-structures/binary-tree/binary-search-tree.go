// Binary search tree
// https://en.wikipedia.org/wiki/Binary_search_tree

package binarySearchTree

// package main

import "fmt"

type node struct {
	val   int
	left  *node
	right *node
}

type tree struct {
	root *node
}

func newNode(val int) *node {
	return &node{val, nil, nil}
}

func inorder(n *node) {
	if n != nil {
		inorder(n.left)
		fmt.Print(n.val, " ")
		inorder(n.right)
	}
}

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

func inorderSuccessor(root *node) *node {
	cur := root
	for cur.left != nil {
		cur = cur.left
	}
	return cur
}

func bst_delete(root *node, val int) *node {
	if root == nil {
		return nil
	}
	if val < root.val {
		root.left = bst_delete(root.left, val)
	} else if val > root.val {
		root.right = bst_delete(root.right, val)
	} else {
		// this is the node to delete

		// node with one child
		if root.left == nil {
			root = root.right
		} else if root.right == nil {
			root = root.left
		} else {
			// node with two children
			d := inorderSuccessor(root)
			root.val = d.val
			root.right = bst_delete(root.right, d.val)
		}
	}
	return root
}

/*
func main() {
	t := &tree{nil}
	inorder(t.root)
	t.root = insert(t.root, 10)
	t.root = insert(t.root, 20)
	t.root = insert(t.root, 15)
	t.root = insert(t.root, 30)
	inorder(t.root)
	fmt.Print("\n")
	t.root = bst_delete(t.root, 10)
	inorder(t.root)
	fmt.Print("\n")
	t.root = bst_delete(t.root, 30)
	inorder(t.root)
	fmt.Print("\n")
	t.root = bst_delete(t.root, 15)
	inorder(t.root)
	fmt.Print("\n")
	t.root = bst_delete(t.root, 20)
	inorder(t.root)
	fmt.Print("\n")
}
*/
