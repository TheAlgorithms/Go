package rbtree

import (
	"fmt"

	"github.com/TheAlgorithms/Go/constraints"
)

type Color byte

const (
	Black Color = iota
	Red
)

type Node[T constraints.Ordered] struct {
	data   T
	parent *Node[T]
	left   *Node[T]
	right  *Node[T]
	color  Color
}

type RBTree[T constraints.Ordered] struct {
	root *Node[T]
	NIL  *Node[T]
}

func NewRBTree[T constraints.Ordered]() *RBTree[T] {
	nilNode := &Node[T]{color: Black, left: nil, right: nil}
	return &RBTree[T]{
		root: nilNode,
		NIL:  nilNode,
	}
}

func (r *RBTree[T]) isNil(node *Node[T]) bool {
	return node == r.NIL
}

func (r *RBTree[T]) Insert(key T) {
	node := &Node[T]{
		data:   key,
		left:   r.NIL,
		right:  r.NIL,
		parent: nil,
		color:  Red,
	}

	var y *Node[T]
	x := r.root

	for !r.isNil(x) {
		y = x
		if node.data < x.data {
			x = x.left
		} else {
			x = x.right
		}

	}

	node.parent = y
	if y == nil {
		r.root = node
	} else if node.data < y.data {
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

func (r *RBTree[T]) InOrder() []T {
	return r.inOrderHelper(r.root)
}

func (r *RBTree[T]) Delete(data T) bool {
	return r.deleteNodeHelper(r.root, data)
}

func (r *RBTree[T]) Max() (T, bool) {
	ret := r.maximum(r.root)
	if r.isNil(ret) {
		return r.NIL.data, false
	}

	return ret.data, true
}

func (r *RBTree[T]) Min() (T, bool) {
	ret := r.minimum(r.root)
	if r.isNil(ret) {
		return r.NIL.data, false
	}

	return ret.data, true
}

func (r *RBTree[T]) Has(key T) bool {
	_, ok := r.searchTreeHelper(r.root, key)
	return ok
}

func (r *RBTree[T]) Predecessor(key T) (T, bool) {
	node, ok := r.searchTreeHelper(r.root, key)
	if !ok {
		return r.NIL.data, ok
	}

	return r.predecessorHelper(node)
}

func (r *RBTree[T]) Successor(key T) (T, bool) {
	node, ok := r.searchTreeHelper(r.root, key)
	if !ok {
		return r.NIL.data, ok
	}

	return r.successorHelper(node)
}

func (r *RBTree[T]) leftRotate(x *Node[T]) {
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

func (r *RBTree[T]) rightRotate(x *Node[T]) {
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

func (r *RBTree[T]) insertFix(k *Node[T]) {
	var u *Node[T]
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

func (r *RBTree[T]) inOrderHelper(node *Node[T]) []T {
	stack := []*Node[T]{}
	ret := []T{}

	for !r.isNil(node) || len(stack) > 0 {
		for !r.isNil(node) {
			stack = append(stack, node)
			node = node.left
		}

		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		ret = append(ret, node.data)
		node = node.right
	}

	return ret
}

func (r *RBTree[T]) searchTreeHelper(node *Node[T], key T) (*Node[T], bool) {
	if node == nil || r.isNil(node) {
		return node, false
	}

	if key == node.data {
		return node, true
	}

	if key < node.data {
		return r.searchTreeHelper(node.left, key)
	}
	return r.searchTreeHelper(node.right, key)
}

func (r *RBTree[T]) deleteNodeHelper(node *Node[T], key T) bool {
	z := r.NIL
	for !r.isNil(node) {
		switch {
		case node.data == key:
			z = node
			fallthrough
		case node.data <= key:
			node = node.right
		case node.data > key:
			node = node.left
		}
	}

	if r.isNil(z) {
		fmt.Println("Key not found in the tree")
		return false
	}

	var x *Node[T]
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

func (r *RBTree[T]) deleteFix(x *Node[T]) {
	var s *Node[T]
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

func (r *RBTree[T]) minimum(node *Node[T]) *Node[T] {
	if r.isNil(node) {
		return node
	}

	for !r.isNil(node.left) {
		node = node.left
	}
	return node
}

func (r *RBTree[T]) maximum(node *Node[T]) *Node[T] {
	if r.isNil(node) {
		return node
	}

	for !r.isNil(node.right) {
		node = node.right
	}

	return node
}

func (r *RBTree[T]) rbTransplant(u, v *Node[T]) {
	if u.parent == nil {
		r.root = v
	} else if u == u.parent.left {
		u.parent.left = v
	} else {
		u.parent.right = v
	}

	v.parent = u.parent
}

func (r *RBTree[T]) predecessorHelper(node *Node[T]) (T, bool) {
	if !r.isNil(node.left) {
		return r.maximum(node.left).data, true
	}

	p := node.parent
	for p != nil && !r.isNil(p) && node == p.left {
		node = p
		p = p.parent
	}

	if p == nil {
		return r.NIL.data, false
	}
	return p.data, true
}

func (r *RBTree[T]) successorHelper(node *Node[T]) (T, bool) {
	if !r.isNil(node.right) {
		return r.minimum(node.right).data, true
	}

	p := node.parent
	for p != nil && !r.isNil(p) && node == p.right {
		node = p
		p = p.parent
	}

	if p == nil {
		return r.NIL.data, false
	}
	return p.data, true
}
