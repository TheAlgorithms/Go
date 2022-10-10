// Red-Black Tree is a kind of self-balancing binary search tree.
// Each node stores "Color" ("red" or "black"), used to ensure that the tree remains balanced during insertions and deletions.
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
	Root *RBNode[T]
	NIL  *RBNode[T] // NIL denotes the leaf node of Red-Black Tree
}

// Create a new Red-Black Tree
func NewRBTree[T constraints.Ordered]() *RBTree[T] {
	leaf := &RBNode[T]{Color: Black, Left: nil, Right: nil}
	return &RBTree[T]{
		Root: leaf,
		NIL:  leaf,
	}
}

// Determine node is a leaf node
func (r *RBTree[T]) isNil(node *RBNode[T]) bool {
	return node == r.NIL
}

// Insert a Key into the tree
func (r *RBTree[T]) Insert(key T) {
	node := &RBNode[T]{
		Key:    key,
		Left:   r.NIL,
		Right:  r.NIL,
		Parent: nil,
		Color:  Red,
	}

	var y *RBNode[T]
	x := r.Root

	for !r.isNil(x) {
		y = x
		if node.Key < x.Key {
			x = x.Left
		} else {
			x = x.Right
		}

	}

	node.Parent = y
	if y == nil {
		r.Root = node
	} else if node.Key < y.Key {
		y.Left = node
	} else {
		y.Right = node
	}

	if node.Parent == nil {
		node.Color = Black
		return
	}

	if node.Parent.Parent == nil {
		return
	}

	r.insertFix(node)
}

// Traverse the tree via inorder.
// Because of the Red-Black tree is a binary search tree, the sequence of inorder traversal is sorted ascendingly.
func (r *RBTree[T]) InOrder() []T {
	return r.inOrderHelper(r.Root)
}

// Delete a node of Red-Black Tree
// Returns false if the node does not exist, otherwise returns true.
func (r *RBTree[T]) Delete(data T) bool {
	return r.deleteNodeHelper(r.Root, data)
}

// Return the max value of the tree
func (r *RBTree[T]) Max() (T, bool) {
	ret := r.maximum(r.Root)
	if r.isNil(ret) {
		return r.NIL.Key, false
	}

	return ret.Key, true
}

// Return the min value of the tree
func (r *RBTree[T]) Min() (T, bool) {
	ret := r.minimum(r.Root)
	if r.isNil(ret) {
		return r.NIL.Key, false
	}

	return ret.Key, true
}

// Determine the tree has the node of Key
func (r *RBTree[T]) Has(key T) bool {
	_, ok := r.searchTreeHelper(r.Root, key)
	return ok
}

// Return the predecessor of the node of Key
// if there is no predecessor, return default value of type T and false
// otherwise return the Key of predecessor and true
func (r *RBTree[T]) Predecessor(key T) (T, bool) {
	node, ok := r.searchTreeHelper(r.Root, key)
	if !ok {
		return r.NIL.Key, ok
	}

	return r.predecessorHelper(node)
}

// Return the successor of the node of Key
// if there is no successor, return default value of type T and false
// otherwise return the Key of predecessor and true
func (r *RBTree[T]) Successor(key T) (T, bool) {
	node, ok := r.searchTreeHelper(r.Root, key)
	if !ok {
		return r.NIL.Key, ok
	}

	return r.successorHelper(node)
}

func (r *RBTree[T]) leftRotate(x *RBNode[T]) {
	y := x.Right
	x.Right = y.Left

	if !r.isNil(y.Left) {
		y.Left.Parent = x
	}

	y.Parent = x.Parent
	if x.Parent == nil {
		r.Root = y
	} else if x == x.Parent.Left {
		x.Parent.Left = y
	} else {
		x.Parent.Right = y
	}

	y.Left = x
	x.Parent = y
}

func (r *RBTree[T]) rightRotate(x *RBNode[T]) {
	y := x.Left
	x.Left = y.Right
	if !r.isNil(y.Right) {
		y.Right.Parent = x
	}
	y.Parent = x.Parent
	if x.Parent == nil {
		r.Root = y
	} else if x == y.Parent.Right {
		y.Parent.Right = y
	} else {
		y.Parent.Left = y
	}

	y.Right = x
	x.Parent = y
}

func (r *RBTree[T]) insertFix(k *RBNode[T]) {
	var u *RBNode[T]
	for k.Parent.Color == Red {
		if k.Parent == k.Parent.Parent.Right {
			u = k.Parent.Parent.Left
			if u != nil && u.Color == Red {
				u.Color = Black
				k.Parent.Color = Black
				k.Parent.Parent.Color = Red
				k = k.Parent.Parent
			} else {
				if k == k.Parent.Left {
					k = k.Parent
					r.rightRotate(k)
				}
				k.Parent.Color = Black
				k.Parent.Parent.Color = Red
				r.leftRotate(k.Parent.Parent)
			}
		} else {
			u = k.Parent.Parent.Right
			if u != nil && u.Color == Red {
				u.Color = Black
				k.Parent.Color = Black
				k.Parent.Parent.Color = Red
				k = k.Parent.Parent
			} else {
				if k == k.Parent.Right {
					k = k.Parent
					r.leftRotate(k)
				}
				k.Parent.Color = Black
				k.Parent.Parent.Color = Red
				r.rightRotate(k.Parent.Parent)
			}
		}
		if k == r.Root {
			break
		}
	}

	r.Root.Color = Black
}

func (r *RBTree[T]) inOrderHelper(node *RBNode[T]) []T {
	stack := []*RBNode[T]{}
	ret := []T{}

	for !r.isNil(node) || len(stack) > 0 {
		for !r.isNil(node) {
			stack = append(stack, node)
			node = node.Left
		}

		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		ret = append(ret, node.Key)
		node = node.Right
	}

	return ret
}

func (r *RBTree[T]) searchTreeHelper(node *RBNode[T], key T) (*RBNode[T], bool) {
	if node == nil || r.isNil(node) {
		return node, false
	}

	if key == node.Key {
		return node, true
	}

	if key < node.Key {
		return r.searchTreeHelper(node.Left, key)
	}
	return r.searchTreeHelper(node.Right, key)
}

func (r *RBTree[T]) deleteNodeHelper(node *RBNode[T], key T) bool {
	z := r.NIL
	for !r.isNil(node) {
		switch {
		case node.Key == key:
			z = node
			fallthrough
		case node.Key <= key:
			node = node.Right
		case node.Key > key:
			node = node.Left
		}
	}

	if r.isNil(z) {
		fmt.Println("Key not found in the tree")
		return false
	}

	var x *RBNode[T]
	y := z
	yOriginColor := y.Color
	if r.isNil(z.Left) {
		x = z.Right
		r.rbTransplant(z, z.Right)
	} else if r.isNil(z.Right) {
		x = z.Left
		r.rbTransplant(z, z.Left)
	} else {
		y = r.minimum(z.Right)
		yOriginColor = y.Color
		x = y.Right
		if y.Parent == z {
			x.Parent = y
		} else {
			r.rbTransplant(y, y.Right)
			y.Right = z.Right
			y.Right.Parent = y
		}

		r.rbTransplant(z, y)
		y.Left = z.Left
		y.Left.Parent = y
		y.Color = z.Color
	}

	if yOriginColor == Black {
		r.deleteFix(x)
	}

	return true
}

func (r *RBTree[T]) deleteFix(x *RBNode[T]) {
	var s *RBNode[T]
	for x != r.Root && x.Color == Black {
		if x == x.Parent.Left {
			s = x.Parent.Right
			if s.Color == Red {
				s.Color = Black
				x.Parent.Color = Red
				r.leftRotate(x.Parent)
				s = x.Parent.Right
			}

			if s.Left.Color == Black && s.Right.Color == Black {
				s.Color = Red
				x = x.Parent
			} else {
				if s.Right.Color == Black {
					s.Left.Color = Black
					s.Color = Red
					r.rightRotate(s)
					s = x.Parent.Right
				}

				s.Color = x.Parent.Color
				x.Parent.Color = Black
				s.Right.Color = Black
				r.leftRotate(x.Parent)
				x = r.Root
			}
		} else {
			s = x.Parent.Left
			if s.Color == Red {
				s.Color = Black
				x.Parent.Color = Red
				r.rightRotate(x.Parent)
				s = x.Parent.Left
			}

			if s.Right.Color == Black && s.Left.Color == Black {
				s.Color = Red
				x = x.Parent
			} else {
				if s.Left.Color == Black {
					s.Right.Color = Black
					s.Color = Red
					r.leftRotate(s)
					s = x.Parent.Left
				}

				s.Color = x.Parent.Color
				x.Parent.Color = Black
				s.Left.Color = Black
				r.rightRotate(x.Parent)
				x = r.Root
			}
		}
	}

	x.Color = Black
}

func (r *RBTree[T]) minimum(node *RBNode[T]) *RBNode[T] {
	if r.isNil(node) {
		return node
	}

	for !r.isNil(node.Left) {
		node = node.Left
	}
	return node
}

func (r *RBTree[T]) maximum(node *RBNode[T]) *RBNode[T] {
	if r.isNil(node) {
		return node
	}

	for !r.isNil(node.Right) {
		node = node.Right
	}

	return node
}

func (r *RBTree[T]) rbTransplant(u, v *RBNode[T]) {
	if u.Parent == nil {
		r.Root = v
	} else if u == u.Parent.Left {
		u.Parent.Left = v
	} else {
		u.Parent.Right = v
	}

	v.Parent = u.Parent
}

func (r *RBTree[T]) predecessorHelper(node *RBNode[T]) (T, bool) {
	if !r.isNil(node.Left) {
		return r.maximum(node.Left).Key, true
	}

	p := node.Parent
	for p != nil && !r.isNil(p) && node == p.Left {
		node = p
		p = p.Parent
	}

	if p == nil {
		return r.NIL.Key, false
	}
	return p.Key, true
}

func (r *RBTree[T]) successorHelper(node *RBNode[T]) (T, bool) {
	if !r.isNil(node.Right) {
		return r.minimum(node.Right).Key, true
	}

	p := node.Parent
	for p != nil && !r.isNil(p) && node == p.Right {
		node = p
		p = p.Parent
	}

	if p == nil {
		return r.NIL.Key, false
	}
	return p.Key, true
}
