package main

import "fmt"

type node struct {
	value       int
	left, right *node
}

func (n *node) recurPreorder() {
	if n == nil {
		return
	}
	fmt.Print(n.value)
	n.left.recurPreorder()
	n.right.recurPreorder()
}

func (n *node) recurInorder() {
	if n == nil {
		return
	}
	n.left.recurInorder()
	fmt.Print(n.value)
	n.right.recurInorder()
}

func (n *node) recurPostorder() {
	if n == nil {
		return
	}
	n.left.recurPostorder()
	n.right.recurPostorder()
	fmt.Print(n.value)
}

func main() {
	tree := &node{1,
		&node{2,
			&node{4,
				&node{7, nil, nil},
				nil},
			&node{5, nil, nil}},
		&node{3,
			&node{6,
				&node{8, nil, nil},
				&node{9, nil, nil}},
			nil}}
	fmt.Print("preorder:    ")
	tree.recurPreorder()
	fmt.Println()
	fmt.Print("inorder:     ")
	tree.recurInorder()
	fmt.Println()
	fmt.Print("postorder:   ")
	tree.recurPostorder()
	fmt.Println()
}
