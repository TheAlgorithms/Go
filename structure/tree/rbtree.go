// Red-Black Tree is a kind of self-balancing binary search tree.
// Each node stores "color" ("red" or "black"), used to ensure that the tree remains balanced during insertions and deletions.
//
// For more details check out those links below here:
// Programiz article : https://www.programiz.com/dsa/red-black-tree
// Wikipedia article: https://en.wikipedia.org/wiki/Red_black_tree
// authors [guuzaa](https://github.com/guuzaa)
// see rbtree.go

package tree

import (
	"fmt"

	"github.com/TheAlgorithms/Go/constraints"
)

type RBTree[T constraints.Ordered] struct {
	root *RBNode[T]
	NIL  *RBNode[T] // NIL denotes the leaf node of Red-Black Tree
}

// Create a new Red-Black Tree
func NewRBTree[T constraints.Ordered]() *RBTree[T] {
	leaf := &RBNode[T]{color: Black, left: nil, right: nil}
	return &RBTree[T]{
		root: leaf,
		NIL:  leaf,
	}
}

// Determine node is a leaf node
func (r *RBTree[T]) isNil(node *RBNode[T]) bool {
	return node == r.NIL
}

// Insert a key into the tree
func (r *RBTree[T]) Insert(key T) {
	node := &RBNode[T]{
		key:    key,
		left:   r.NIL,
		right:  r.NIL,
		parent: nil,
		color:  Red,
	}

	var y *RBNode[T]
	x := r.root

	for !r.isNil(x) {
		y = x
		if node.key < x.key {
			x = x.left
		} else {
			x = x.right
		}

	}

	node.parent = y
	if y == nil {
		r.root = node
	} else if node.key < y.key {
		y.left = node
	} else {
		y.right = node
	}

	if node.parent == nil {
		node.color = Black
		return
	}

	if node.parent.parent == nil {
		return
	}

	r.insertFix(node)
}

// Traverse the tree via inorder.
// Because of the Red-Black tree is a binary search tree, the sequence of inorder traversal is sorted ascendingly.
func (r *RBTree[T]) InOrder() []T {
	return r.inOrderHelper(r.root)
}

// Delete a node of Red-Black Tree
// Returns false if the node does not exist, otherwise returns true.
func (r *RBTree[T]) Delete(data T) bool {
	return r.deleteNodeHelper(r.root, data)
}

// Return the max value of the tree
func (r *RBTree[T]) Max() (T, bool) {
	ret := r.maximum(r.root)
	if r.isNil(ret) {
		return r.NIL.key, false
	}

	return ret.key, true
}

// Return the min value of the tree
func (r *RBTree[T]) Min() (T, bool) {
	ret := r.minimum(r.root)
	if r.isNil(ret) {
		return r.NIL.key, false
	}

	return ret.key, true
}

// Determine the tree has the node of key
func (r *RBTree[T]) Has(key T) bool {
	_, ok := r.searchTreeHelper(r.root, key)
	return ok
}

// Return the predecessor of the node of key
// if there is no predecessor, return default value of type T and false
// otherwise return the key of predecessor and true
func (r *RBTree[T]) Predecessor(key T) (T, bool) {
	node, ok := r.searchTreeHelper(r.root, key)
	if !ok {
		return r.NIL.key, ok
	}

	return r.predecessorHelper(node)
}

// Return the successor of the node of key
// if there is no successor, return default value of type T and false
// otherwise return the key of predecessor and true
func (r *RBTree[T]) Successor(key T) (T, bool) {
	node, ok := r.searchTreeHelper(r.root, key)
	if !ok {
		return r.NIL.key, ok
	}

	return r.successorHelper(node)
}

func (r *RBTree[T]) leftRotate(x *RBNode[T]) {
	y := x.right
	x.right = y.left

	if !r.isNil(y.left) {
		y.left.parent = x
	}

	y.parent = x.parent
	if x.parent == nil {
		r.root = y
	} else if x == x.parent.left {
		x.parent.left = y
	} else {
		x.parent.right = y
	}

	y.left = x
	x.parent = y
}

func (r *RBTree[T]) rightRotate(x *RBNode[T]) {
	y := x.left
	x.left = y.right
	if !r.isNil(y.right) {
		y.right.parent = x
	}
	y.parent = x.parent
	if x.parent == nil {
		r.root = y
	} else if x == y.parent.right {
		y.parent.right = y
	} else {
		y.parent.left = y
	}

	y.right = x
	x.parent = y
}

func (r *RBTree[T]) insertFix(k *RBNode[T]) {
	var u *RBNode[T]
	for k.parent.color == Red {
		if k.parent == k.parent.parent.right {
			u = k.parent.parent.left
			if u != nil && u.color == Red {
				u.color = Black
				k.parent.color = Black
				k.parent.parent.color = Red
				k = k.parent.parent
			} else {
				if k == k.parent.left {
					k = k.parent
					r.rightRotate(k)
				}
				k.parent.color = Black
				k.parent.parent.color = Red
				r.leftRotate(k.parent.parent)
			}
		} else {
			u = k.parent.parent.right
			if u != nil && u.color == Red {
				u.color = Black
				k.parent.color = Black
				k.parent.parent.color = Red
				k = k.parent.parent
			} else {
				if k == k.parent.right {
					k = k.parent
					r.leftRotate(k)
				}
				k.parent.color = Black
				k.parent.parent.color = Red
				r.rightRotate(k.parent.parent)
			}
		}
		if k == r.root {
			break
		}
	}

	r.root.color = Black
}

func (r *RBTree[T]) inOrderHelper(node *RBNode[T]) []T {
	stack := []*RBNode[T]{}
	ret := []T{}

	for !r.isNil(node) || len(stack) > 0 {
		for !r.isNil(node) {
			stack = append(stack, node)
			node = node.left
		}

		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		ret = append(ret, node.key)
		node = node.right
	}

	return ret
}

func (r *RBTree[T]) searchTreeHelper(node *RBNode[T], key T) (*RBNode[T], bool) {
	if node == nil || r.isNil(node) {
		return node, false
	}

	if key == node.key {
		return node, true
	}

	if key < node.key {
		return r.searchTreeHelper(node.left, key)
	}
	return r.searchTreeHelper(node.right, key)
}

func (r *RBTree[T]) deleteNodeHelper(node *RBNode[T], key T) bool {
	z := r.NIL
	for !r.isNil(node) {
		switch {
		case node.key == key:
			z = node
			fallthrough
		case node.key <= key:
			node = node.right
		case node.key > key:
			node = node.left
		}
	}

	if r.isNil(z) {
		fmt.Println("key not found in the tree")
		return false
	}

	var x *RBNode[T]
	y := z
	yOriginColor := y.color
	if r.isNil(z.left) {
		x = z.right
		r.rbTransplant(z, z.right)
	} else if r.isNil(z.right) {
		x = z.left
		r.rbTransplant(z, z.left)
	} else {
		y = r.minimum(z.right)
		yOriginColor = y.color
		x = y.right
		if y.parent == z {
			x.parent = y
		} else {
			r.rbTransplant(y, y.right)
			y.right = z.right
			y.right.parent = y
		}

		r.rbTransplant(z, y)
		y.left = z.left
		y.left.parent = y
		y.color = z.color
	}

	if yOriginColor == Black {
		r.deleteFix(x)
	}

	return true
}

func (r *RBTree[T]) deleteFix(x *RBNode[T]) {
	var s *RBNode[T]
	for x != r.root && x.color == Black {
		if x == x.parent.left {
			s = x.parent.right
			if s.color == Red {
				s.color = Black
				x.parent.color = Red
				r.leftRotate(x.parent)
				s = x.parent.right
			}

			if s.left.color == Black && s.right.color == Black {
				s.color = Red
				x = x.parent
			} else {
				if s.right.color == Black {
					s.left.color = Black
					s.color = Red
					r.rightRotate(s)
					s = x.parent.right
				}

				s.color = x.parent.color
				x.parent.color = Black
				s.right.color = Black
				r.leftRotate(x.parent)
				x = r.root
			}
		} else {
			s = x.parent.left
			if s.color == Red {
				s.color = Black
				x.parent.color = Red
				r.rightRotate(x.parent)
				s = x.parent.left
			}

			if s.right.color == Black && s.left.color == Black {
				s.color = Red
				x = x.parent
			} else {
				if s.left.color == Black {
					s.right.color = Black
					s.color = Red
					r.leftRotate(s)
					s = x.parent.left
				}

				s.color = x.parent.color
				x.parent.color = Black
				s.left.color = Black
				r.rightRotate(x.parent)
				x = r.root
			}
		}
	}

	x.color = Black
}

func (r *RBTree[T]) minimum(node *RBNode[T]) *RBNode[T] {
	if r.isNil(node) {
		return node
	}

	for !r.isNil(node.left) {
		node = node.left
	}
	return node
}

func (r *RBTree[T]) maximum(node *RBNode[T]) *RBNode[T] {
	if r.isNil(node) {
		return node
	}

	for !r.isNil(node.right) {
		node = node.right
	}

	return node
}

func (r *RBTree[T]) rbTransplant(u, v *RBNode[T]) {
	if u.parent == nil {
		r.root = v
	} else if u == u.parent.left {
		u.parent.left = v
	} else {
		u.parent.right = v
	}

	v.parent = u.parent
}

func (r *RBTree[T]) predecessorHelper(node *RBNode[T]) (T, bool) {
	if !r.isNil(node.left) {
		return r.maximum(node.left).key, true
	}

	p := node.parent
	for p != nil && !r.isNil(p) && node == p.left {
		node = p
		p = p.parent
	}

	if p == nil {
		return r.NIL.key, false
	}
	return p.key, true
}

func (r *RBTree[T]) successorHelper(node *RBNode[T]) (T, bool) {
	if !r.isNil(node.right) {
		return r.minimum(node.right).key, true
	}

	p := node.parent
	for p != nil && !r.isNil(p) && node == p.right {
		node = p
		p = p.parent
	}

	if p == nil {
		return r.NIL.key, false
	}
	return p.key, true
}
