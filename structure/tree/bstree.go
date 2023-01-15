// Binary search tree.
//
// For more details check out those links below here:
// Wikipedia article: https://en.wikipedia.org/wiki/Binary_search_tree
// see bstree.go

package tree

import "github.com/TheAlgorithms/Go/constraints"

// Verify Interface Compliance
var _ Node[int] = &BSNode[int]{}

// BSNode represents a single node in the BinarySearch.
type BSNode[T constraints.Ordered] struct {
	key    T
	parent *BSNode[T]
	left   *BSNode[T]
	right  *BSNode[T]
}

func (n *BSNode[T]) Key() T {
	return n.key
}

func (n *BSNode[T]) Parent() Node[T] {
	return n.parent
}

func (n *BSNode[T]) Left() Node[T] {
	return n.left
}

func (n *BSNode[T]) Right() Node[T] {
	return n.right
}

// BinarySearch represents a Binary-Search tree.
// By default, _NIL = nil.
type BinarySearch[T constraints.Ordered] struct {
	Root *BSNode[T]
	_NIL *BSNode[T] // a sentinel value for nil
}

// NewBinarySearch creates a novel Binary-Search tree
func NewBinarySearch[T constraints.Ordered]() *BinarySearch[T] {
	return &BinarySearch[T]{
		Root: nil,
		_NIL: nil,
	}
}

// Empty determines the Binary-Search tree is empty
func (t *BinarySearch[T]) Empty() bool {
	return t.Root == t._NIL
}

// Push a chain of Node's into the BinarySearch
func (t *BinarySearch[T]) Push(keys ...T) {
	for _, key := range keys {
		t.pushHelper(t.Root, key)
	}
}

// Delete removes the node of val
func (t *BinarySearch[T]) Delete(val T) bool {
	node, ok := t.Get(val)
	if !ok {
		return false
	}
	t.deleteHelper(node.(*BSNode[T]))
	return true
}

// Get a Node from the Binary-Search Tree
func (t *BinarySearch[T]) Get(key T) (Node[T], bool) {
	return searchTreeHelper[T](t.Root, t._NIL, key)
}

// Has Determines the tree has the node of Key
func (t *BinarySearch[T]) Has(key T) bool {
	_, ok := searchTreeHelper[T](t.Root, t._NIL, key)
	return ok
}

// PreOrder Traverses the tree in the following order Root --> Left --> Right
func (t *BinarySearch[T]) PreOrder() []T {
	traversal := make([]T, 0)
	preOrderRecursive[T](t.Root, t._NIL, &traversal)
	return traversal
}

// InOrder Traverses the tree in the following order Left --> Root --> Right
func (t *BinarySearch[T]) InOrder() []T {
	return inOrderHelper[T](t.Root, t._NIL)
}

// PostOrder traverses the tree in the following order Left --> Right --> Root
func (t *BinarySearch[T]) PostOrder() []T {
	traversal := make([]T, 0)
	postOrderRecursive[T](t.Root, t._NIL, &traversal)
	return traversal
}

// LevelOrder returns the level order traversal of the tree
func (t *BinarySearch[T]) LevelOrder() []T {
	traversal := make([]T, 0)
	levelOrderHelper[T](t.Root, t._NIL, &traversal)
	return traversal
}

// AccessNodesByLayer accesses nodes layer by layer (2-D array),  instead of printing the results as 1-D array.
func (t *BinarySearch[T]) AccessNodesByLayer() [][]T {
	return accessNodeByLayerHelper[T](t.Root, t._NIL)
}

// Depth returns the calculated depth of a binary search tree
func (t *BinarySearch[T]) Depth() int {
	return calculateDepth[T](t.Root, t._NIL, 0)
}

// Max returns the Max value of the tree
func (t *BinarySearch[T]) Max() (T, bool) {
	ret := maximum[T](t.Root, t._NIL)
	if ret == t._NIL {
		var dft T
		return dft, false
	}
	return ret.Key(), true
}

// Min returns the Min value of the tree
func (t *BinarySearch[T]) Min() (T, bool) {
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
func (t *BinarySearch[T]) Predecessor(key T) (T, bool) {
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
func (t *BinarySearch[T]) Successor(key T) (T, bool) {
	node, ok := searchTreeHelper[T](t.Root, t._NIL, key)
	if !ok {
		var dft T
		return dft, ok
	}
	return successorHelper[T](node, t._NIL)
}

func (t *BinarySearch[T]) pushHelper(x *BSNode[T], val T) {
	y := t._NIL
	for x != t._NIL {
		y = x
		switch {
		case val < x.Key():
			x = x.left
		case val > x.Key():
			x = x.right
		default:
			return
		}
	}

	z := &BSNode[T]{
		key:    val,
		left:   t._NIL,
		right:  t._NIL,
		parent: y,
	}
	if y == t._NIL {
		t.Root = z
	} else if val < y.key {
		y.left = z
	} else {
		y.right = z
	}
}

func (t *BinarySearch[T]) deleteHelper(z *BSNode[T]) {
	switch {
	case z.left == t._NIL:
		t.transplant(z, z.right)
	case z.right == t._NIL:
		t.transplant(z, z.left)
	default:
		y := minimum[T](z.right, t._NIL).(*BSNode[T])
		if y.parent != z {
			t.transplant(y, y.right)
			y.right = z.right
			y.right.parent = y
		}

		t.transplant(z, y)
		y.left = z.left
		y.left.parent = y
	}
}

func (t *BinarySearch[T]) transplant(u, v *BSNode[T]) {
	switch {
	case u.parent == t._NIL:
		t.Root = v
	case u == u.parent.left:
		u.parent.left = v
	default:
		u.parent.right = v
	}

	if v != t._NIL {
		v.parent = u.parent
	}
}
