// Red-Black Tree is a kind of self-balancing binary search tree.
// Each node stores "color" ("red" or "black"), used to ensure that the tree remains balanced during insertions and deletions.
//
// For more details check out those links below here:
// Programiz article : https://www.programiz.com/dsa/red-black-tree
// Wikipedia article: https://en.wikipedia.org/wiki/Red_black_tree
// authors [guuzaa](https://github.com/guuzaa)
// see rbtree.go

package tree

import "github.com/TheAlgorithms/Go/constraints"

type Color byte

const (
	Red Color = iota
	Black
)

// Verify Interface Compliance
var _ Node[int] = &RBNode[int]{}

// RBNode represents a single node in the RB.
type RBNode[T constraints.Ordered] struct {
	key    T
	parent *RBNode[T]
	left   *RBNode[T]
	right  *RBNode[T]
	color  Color
}

func (n *RBNode[T]) Key() T {
	return n.key
}

func (n *RBNode[T]) Parent() Node[T] {
	return n.parent
}

func (n *RBNode[T]) Left() Node[T] {
	return n.left
}

func (n *RBNode[T]) Right() Node[T] {
	return n.right
}

// RB represents a Red-Black tree.
// By default, _NIL = leaf, a dummy variable.
type RB[T constraints.Ordered] struct {
	Root *RBNode[T]
	_NIL *RBNode[T] // a sentinel value for nil
}

// NewRB creates a new Red-Black Tree
func NewRB[T constraints.Ordered]() *RB[T] {
	leaf := &RBNode[T]{color: Black, left: nil, right: nil}
	leaf.parent = leaf
	return &RB[T]{
		Root: leaf,
		_NIL: leaf,
	}
}

// Empty determines the Red-Black tree is empty
func (t *RB[T]) Empty() bool {
	return t.Root == t._NIL
}

// Push a chain of Node's into the Red-Black Tree
func (t *RB[T]) Push(keys ...T) {
	for _, key := range keys {
		t.pushHelper(t.Root, key)
	}
}

// Delete a node of Red-Black Tree
// Returns false if the node does not exist, otherwise returns true.
func (t *RB[T]) Delete(data T) bool {
	return t.deleteHelper(t.Root, data)
}

// Get a Node from the Red-Black Tree
func (t *RB[T]) Get(key T) (Node[T], bool) {
	return searchTreeHelper[T](t.Root, t._NIL, key)
}

// Has Determines the tree has the node of Key
func (t *RB[T]) Has(key T) bool {
	_, ok := searchTreeHelper[T](t.Root, t._NIL, key)
	return ok
}

// PreOrder Traverses the tree in the following order Root --> Left --> Right
func (t *RB[T]) PreOrder() []T {
	traversal := make([]T, 0)
	preOrderRecursive[T](t.Root, t._NIL, &traversal)
	return traversal
}

// InOrder Traverses the tree in the following order Left --> Root --> Right
func (t *RB[T]) InOrder() []T {
	return inOrderHelper[T](t.Root, t._NIL)
}

// PostOrder traverses the tree in the following order Left --> Right --> Root
func (t *RB[T]) PostOrder() []T {
	traversal := make([]T, 0)
	postOrderRecursive[T](t.Root, t._NIL, &traversal)
	return traversal
}

// LevelOrder returns the level order traversal of the tree
func (t *RB[T]) LevelOrder() []T {
	traversal := make([]T, 0)
	levelOrderHelper[T](t.Root, t._NIL, &traversal)
	return traversal
}

// AccessNodesByLayer accesses nodes layer by layer (2-D array),  instead of printing the results as 1-D array.
func (t *RB[T]) AccessNodesByLayer() [][]T {
	return accessNodeByLayerHelper[T](t.Root, t._NIL)
}

// Depth returns the calculated depth of a Red-Black tree
func (t *RB[T]) Depth() int {
	return calculateDepth[T](t.Root, t._NIL, 0)
}

// Max returns the Max value of the tree
func (t *RB[T]) Max() (T, bool) {
	ret := maximum[T](t.Root, t._NIL)
	if ret == t._NIL {
		var dft T
		return dft, false
	}
	return ret.Key(), true
}

// Min returns the Min value of the tree
func (t *RB[T]) Min() (T, bool) {
	ret := minimum[T](t.Root, t._NIL)
	if ret == t._NIL {
		var dft T
		return dft, false
	}
	return ret.Key(), true
}

// Predecessor returns the Predecessor of the node of Key
// if there is no predecessor, return default value of type T and false
// otherwise return the Key of predecessor and true
func (t *RB[T]) Predecessor(key T) (T, bool) {
	node, ok := searchTreeHelper[T](t.Root, t._NIL, key)
	if !ok {
		var dft T
		return dft, ok
	}
	return predecessorHelper[T](node, t._NIL)
}

// Successor returns the Successor of the node of Key
// if there is no successor, return default value of type T and false
// otherwise return the Key of successor and true
func (t *RB[T]) Successor(key T) (T, bool) {
	node, ok := searchTreeHelper[T](t.Root, t._NIL, key)
	if !ok {
		var dft T
		return dft, ok
	}
	return successorHelper[T](node, t._NIL)
}

func (t *RB[T]) pushHelper(x *RBNode[T], key T) {
	y := t._NIL
	for x != t._NIL {
		y = x
		switch {
		case key < x.Key():
			x = x.left
		case key > x.Key():
			x = x.right
		default:
			return
		}
	}

	node := &RBNode[T]{
		key:    key,
		left:   t._NIL,
		right:  t._NIL,
		parent: y,
		color:  Red,
	}
	if y == t._NIL {
		t.Root = node
	} else if node.key < y.key {
		y.left = node
	} else {
		y.right = node
	}

	if node.parent == t._NIL {
		node.color = Black
		return
	}

	if node.parent.parent == t._NIL {
		return
	}

	t.pushFix(node)
}

func (t *RB[T]) leftRotate(x *RBNode[T]) {
	y := x.right
	x.right = y.left

	if y.left != t._NIL {
		y.left.parent = x
	}

	y.parent = x.parent
	if x.parent == t._NIL {
		t.Root = y
	} else if x == x.parent.left {
		x.parent.left = y
	} else {
		x.parent.right = y
	}

	y.left = x
	x.parent = y
}

func (t *RB[T]) rightRotate(x *RBNode[T]) {
	y := x.left
	x.left = y.right
	if y.right != t._NIL {
		y.right.parent = x
	}

	y.parent = x.parent
	if x.parent == t._NIL {
		t.Root = y
	} else if x == y.parent.right {
		y.parent.right = y
	} else {
		y.parent.left = y
	}

	y.right = x
	x.parent = y
}

func (t *RB[T]) pushFix(k *RBNode[T]) {
	for k.parent.color == Red {
		if k.parent == k.parent.parent.right {
			u := k.parent.parent.left
			if u.color == Red {
				u.color = Black
				k.parent.color = Black
				k.parent.parent.color = Red
				k = k.parent.parent
			} else {
				if k == k.parent.left {
					k = k.parent
					t.rightRotate(k)
				}
				k.parent.color = Black
				k.parent.parent.color = Red
				t.leftRotate(k.parent.parent)
			}
		} else {
			u := k.parent.parent.right
			if u.color == Red {
				u.color = Black
				k.parent.color = Black
				k.parent.parent.color = Red
				k = k.parent.parent
			} else {
				if k == k.parent.right {
					k = k.parent
					t.leftRotate(k)
				}
				k.parent.color = Black
				k.parent.parent.color = Red
				t.rightRotate(k.parent.parent)
			}
		}
		if k == t.Root {
			break
		}
	}

	t.Root.color = Black
}

func (t *RB[T]) deleteHelper(node *RBNode[T], key T) bool {
	z := t._NIL
	for node != t._NIL {
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

	if z == t._NIL {
		return false
	}

	var x *RBNode[T]
	y := z
	yOriginColor := y.color
	if z.left == t._NIL {
		x = z.right
		t.transplant(z, z.right)
	} else if z.right == t._NIL {
		x = z.left
		t.transplant(z, z.left)
	} else {
		y = minimum[T](z.right, t._NIL).(*RBNode[T])
		yOriginColor = y.color
		x = y.right
		if y.parent == z {
			x.parent = y
		} else {
			t.transplant(y, y.right)
			y.right = z.right
			y.right.parent = y
		}

		t.transplant(z, y)
		y.left = z.left
		y.left.parent = y
		y.color = z.color
	}

	if yOriginColor == Black {
		t.deleteFix(x)
	}

	return true
}

func (t *RB[T]) deleteFix(x *RBNode[T]) {
	var s *RBNode[T]
	for x != t.Root && x.color == Black {
		if x == x.parent.left {
			s = x.parent.right
			if s.color == Red {
				s.color = Black
				x.parent.color = Red
				t.leftRotate(x.parent)
				s = x.parent.right
			}

			if s.left.color == Black && s.right.color == Black {
				s.color = Red
				x = x.parent
			} else {
				if s.right.color == Black {
					s.left.color = Black
					s.color = Red
					t.rightRotate(s)
					s = x.parent.right
				}

				s.color = x.parent.color
				x.parent.color = Black
				s.right.color = Black
				t.leftRotate(x.parent)
				x = t.Root
			}
		} else {
			s = x.parent.left
			if s.color == Red {
				s.color = Black
				x.parent.color = Red
				t.rightRotate(x.parent)
				s = x.parent.left
			}

			if s.right.color == Black && s.left.color == Black {
				s.color = Red
				x = x.parent
			} else {
				if s.left.color == Black {
					s.right.color = Black
					s.color = Red
					t.leftRotate(s)
					s = x.parent.left
				}

				s.color = x.parent.color
				x.parent.color = Black
				s.left.color = Black
				t.rightRotate(x.parent)
				x = t.Root
			}
		}
	}

	x.color = Black
}

func (t *RB[T]) transplant(u, v *RBNode[T]) {
	switch {
	case u.parent == t._NIL:
		t.Root = v
	case u == u.parent.left:
		u.parent.left = v
	default:
		u.parent.right = v
	}

	v.parent = u.parent
}
