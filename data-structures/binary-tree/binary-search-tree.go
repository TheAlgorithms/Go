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

type btree struct {
	root *node
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
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
			return root.right
		} else if root.right == nil {
			return root.left
		} else {
			n := root.right
			d := inorderSuccessor(n)
			d.left = root.left
			return root.right
		}
	}
	return root
}

// helper function for t.depth
func _calculate_depth(n *node, depth int) int {
	if n == nil {
		return depth
	}
	return max(_calculate_depth(n.left, depth+1), _calculate_depth(n.right, depth+1))
}

func (t *btree) depth() int {
	return _calculate_depth(t.root, 0)
}


func main() {
	t := &btree{nil}
	inorder(t.root)
	t.root = insert(t.root, 30)

	t.root = insert(t.root, 20)
	t.root = insert(t.root, 15)
	t.root = insert(t.root, 10)
	t.root = insert(t.root, 12)
	t.root = insert(t.root, 9)
	t.root = insert(t.root, 11)
	t.root = insert(t.root, 17)
	fmt.Print(t.depth(), "\n")
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
	fmt.Print(t.depth(), "\n")
}

