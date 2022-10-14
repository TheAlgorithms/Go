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
	*binaryTree[T]
}

// Create a new Red-Black Tree
func NewRBTree[T constraints.Ordered]() *RBTree[T] {
	leaf := &Node[T]{Color: Black, Left: nil, Right: nil}
	return &RBTree[T]{
		binaryTree: &binaryTree[T]{
			Root: leaf,
			NIL:  leaf,
		},
	}
}

// Insert a chain of Node's into the Red-Black Tree
func (t *RBTree[T]) Insert(keys ...T) {
	for _, key := range keys {
		t.insertHelper(t.Root, key)
	}
}

// Delete a node of Red-Black Tree
// Returns false if the node does not exist, otherwise returns true.
func (t *RBTree[T]) Delete(data T) bool {
	return t.deleteHelper(t.Root, data)
}

// Return the Predecessor of the node of Key
// if there is no predecessor, return default value of type T and false
// otherwise return the Key of predecessor and true
func (t *RBTree[T]) Predecessor(key T) (T, bool) {
	node, ok := t.searchTreeHelper(t.Root, key)
	if !ok {
		return t.NIL.Key, ok
	}
	return t.predecessorHelper(node)
}

// Return the Successor of the node of Key
// if there is no successor, return default value of type T and false
// otherwise return the Key of successor and true
func (t *RBTree[T]) Successor(key T) (T, bool) {
	node, ok := t.searchTreeHelper(t.Root, key)
	if !ok {
		return t.NIL.Key, ok
	}

	return t.successorHelper(node)
}

func (t *RBTree[T]) insertHelper(x *Node[T], key T) {
	node := &Node[T]{
		Key:    key,
		Left:   t.NIL,
		Right:  t.NIL,
		Parent: nil,
		Color:  Red,
	}

	var y *Node[T]
	for !t.isNil(x) {
		y = x
		if node.Key < x.Key {
			x = x.Left
		} else {
			x = x.Right
		}

	}

	node.Parent = y
	if y == nil {
		t.Root = node
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

	t.insertFix(node)
}

func (t *RBTree[T]) leftRotate(x *Node[T]) {
	y := x.Right
	x.Right = y.Left

	if !t.isNil(y.Left) {
		y.Left.Parent = x
	}

	y.Parent = x.Parent
	if x.Parent == nil {
		t.Root = y
	} else if x == x.Parent.Left {
		x.Parent.Left = y
	} else {
		x.Parent.Right = y
	}

	y.Left = x
	x.Parent = y
}

func (t *RBTree[T]) rightRotate(x *Node[T]) {
	y := x.Left
	x.Left = y.Right
	if !t.isNil(y.Right) {
		y.Right.Parent = x
	}
	y.Parent = x.Parent
	if x.Parent == nil {
		t.Root = y
	} else if x == y.Parent.Right {
		y.Parent.Right = y
	} else {
		y.Parent.Left = y
	}

	y.Right = x
	x.Parent = y
}

func (t *RBTree[T]) insertFix(k *Node[T]) {
	var u *Node[T]
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
					t.rightRotate(k)
				}
				k.Parent.Color = Black
				k.Parent.Parent.Color = Red
				t.leftRotate(k.Parent.Parent)
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
					t.leftRotate(k)
				}
				k.Parent.Color = Black
				k.Parent.Parent.Color = Red
				t.rightRotate(k.Parent.Parent)
			}
		}
		if k == t.Root {
			break
		}
	}

	t.Root.Color = Black
}

func (t *RBTree[T]) deleteHelper(node *Node[T], key T) bool {
	z := t.NIL
	for !t.isNil(node) {
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

	if t.isNil(z) {
		fmt.Println("Key not found in the tree")
		return false
	}

	var x *Node[T]
	y := z
	yOriginColor := y.Color
	if t.isNil(z.Left) {
		x = z.Right
		t.rbTransplant(z, z.Right)
	} else if t.isNil(z.Right) {
		x = z.Left
		t.rbTransplant(z, z.Left)
	} else {
		y = t.minimum(z.Right)
		yOriginColor = y.Color
		x = y.Right
		if y.Parent == z {
			x.Parent = y
		} else {
			t.rbTransplant(y, y.Right)
			y.Right = z.Right
			y.Right.Parent = y
		}

		t.rbTransplant(z, y)
		y.Left = z.Left
		y.Left.Parent = y
		y.Color = z.Color
	}

	if yOriginColor == Black {
		t.deleteFix(x)
	}

	return true
}

func (t *RBTree[T]) deleteFix(x *Node[T]) {
	var s *Node[T]
	for x != t.Root && x.Color == Black {
		if x == x.Parent.Left {
			s = x.Parent.Right
			if s.Color == Red {
				s.Color = Black
				x.Parent.Color = Red
				t.leftRotate(x.Parent)
				s = x.Parent.Right
			}

			if s.Left.Color == Black && s.Right.Color == Black {
				s.Color = Red
				x = x.Parent
			} else {
				if s.Right.Color == Black {
					s.Left.Color = Black
					s.Color = Red
					t.rightRotate(s)
					s = x.Parent.Right
				}

				s.Color = x.Parent.Color
				x.Parent.Color = Black
				s.Right.Color = Black
				t.leftRotate(x.Parent)
				x = t.Root
			}
		} else {
			s = x.Parent.Left
			if s.Color == Red {
				s.Color = Black
				x.Parent.Color = Red
				t.rightRotate(x.Parent)
				s = x.Parent.Left
			}

			if s.Right.Color == Black && s.Left.Color == Black {
				s.Color = Red
				x = x.Parent
			} else {
				if s.Left.Color == Black {
					s.Right.Color = Black
					s.Color = Red
					t.leftRotate(s)
					s = x.Parent.Left
				}

				s.Color = x.Parent.Color
				x.Parent.Color = Black
				s.Left.Color = Black
				t.rightRotate(x.Parent)
				x = t.Root
			}
		}
	}

	x.Color = Black
}

func (t *RBTree[T]) rbTransplant(u, v *Node[T]) {
	if u.Parent == nil {
		t.Root = v
	} else if u == u.Parent.Left {
		u.Parent.Left = v
	} else {
		u.Parent.Right = v
	}

	v.Parent = u.Parent
}

func (t *RBTree[T]) predecessorHelper(node *Node[T]) (T, bool) {
	if !t.isNil(node.Left) {
		return t.maximum(node.Left).Key, true
	}

	p := node.Parent
	for p != nil && !t.isNil(p) && node == p.Left {
		node = p
		p = p.Parent
	}

	if p == nil {
		return t.NIL.Key, false
	}
	return p.Key, true
}

func (t *RBTree[T]) successorHelper(node *Node[T]) (T, bool) {
	if !t.isNil(node.Right) {
		return t.minimum(node.Right).Key, true
	}

	p := node.Parent
	for p != nil && !t.isNil(p) && node == p.Right {
		node = p
		p = p.Parent
	}

	if p == nil {
		return t.NIL.Key, false
	}
	return p.Key, true
}
