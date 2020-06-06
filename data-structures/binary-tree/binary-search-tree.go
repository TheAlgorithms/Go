// Binary search tree
// https://en.wikipedia.org/wiki/Binary_search_tree

package binarySearchTree

//package main

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

func preorder(n *node) {
	if n != nil {
		fmt.Print(n.val, " ")
		preorder(n.left)
		preorder(n.right)
	}
}

func postorder(n *node) {
	if n != nil {
		postorder(n.left)
		postorder(n.right)
		fmt.Print(n.val, " ")
	}
}

func print2DUtil(n *node, lvl int) {
	if n != nil {
		lvl += 10
		print2DUtil(n.right, lvl)
		fmt.Print("\n")

		for i := 10; i < lvl; i++ {
			fmt.Print(" ")
		}
		fmt.Print(n.val, "\n")
		print2DUtil(n.left, lvl)
	}
}

func print2D(n *node) {
	print2DUtil(n, 0)
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

/*
func main() {
	t := &btree{nil}
	inorder(t.root)
	t.root = insert(t.root, 15)
	t.root = insert(t.root, 22)
	t.root = insert(t.root, 25)
	t.root = insert(t.root, 35)
	t.root = insert(t.root, 50)
	t.root = insert(t.root, 10)
	t.root = insert(t.root, 4)
	t.root = insert(t.root, 12)
	t.root = insert(t.root, 44)
	t.root = insert(t.root, 31)
	t.root = insert(t.root, 18)
	t.root = insert(t.root, 24)
	t.root = insert(t.root, 90)
	t.root = insert(t.root, 70)
	t.root = insert(t.root, 66)

	print2D(t.root)

	fmt.Print(t.depth(), "\nInorder\n")
	inorder(t.root)
	fmt.Print("\n")

	fmt.Print("Preorder\n")
	preorder(t.root)
	fmt.Print("\n")
	fmt.Print("Postorder\n")
	postorder(t.root)
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
*/
