// Binary-Search tree is the tree, with the key of each internal node
// being greater than keys in the respective node's left subtree
// and less than the ones in its right subtree.

// For more details check out those links below here:
// Wikipedia article: https://en.wikipedia.org/wiki/Binary_search_tree
// authors [guuzaa](https://github.com/guuzaa)
package tree

import (
	"fmt"

	"github.com/TheAlgorithms/Go/constraints"
	"github.com/TheAlgorithms/Go/math/max"
)

type Color byte

const (
	Red Color = iota
	Black
)

type Tree[T constraints.Ordered] interface {
	Push(...T)
	Delete(T) bool
	Get(T) (*Node[T], bool)
	Has(T) bool
	Depth() int
	Max() (T, bool)
	Min() (T, bool)
	Predecessor(T) (T, bool)
	Successor(T) (T, bool)
	PreOrder() []T
	InOrder() []T
	PostOrder() []T
	LevelOrder() []T
	AccessNodesByLayer() [][]T
	Print()
}

// Node of a binary tree
type Node[T constraints.Ordered] struct {
	Key    T
	Parent *Node[T]
	Left   *Node[T]
	Right  *Node[T]
	Color  Color // for Red-Black Tree
	Height int   // for AVL Tree
}

// binaryTree is a base-struct for BinarySearch, AVL, RB, etc.
// Assumption: all keys in the binaryTree are different.
// Note: to avoid instantiation, we make the base struct un-exported.
type binaryTree[T constraints.Ordered] struct {
	Root *Node[T]
	NIL  *Node[T] // NIL denotes the leaf node, which is `nil` in BinarySearch and AVL and a sentinel value in RB
}

// Get a Node from the binary-search Tree
func (t *binaryTree[T]) Get(key T) (*Node[T], bool) {
	return t.searchTreeHelper(t.Root, key)
}

// Determines the tree has the node of Key
func (t *binaryTree[T]) Has(key T) bool {
	_, ok := t.searchTreeHelper(t.Root, key)
	return ok
}

// Traverses the tree in the following order Root --> Left --> Right
func (t *binaryTree[T]) PreOrder() []T {
	traversal := make([]T, 0)
	t.preOrderRecursive(t.Root, &traversal)
	return traversal
}

// Traverses the tree in the following order Left --> Root --> Right
func (t *binaryTree[T]) InOrder() []T {
	return t.inOrderHelper(t.Root)
}

// Traverses the tree in the following order Left --> Right --> Root
func (t *binaryTree[T]) PostOrder() []T {
	traversal := make([]T, 0)
	t.postOrderRecursive(t.Root, &traversal)
	return traversal
}

// Depth returns the calculated depth of a binary search tree
func (t *binaryTree[T]) Depth() int {
	return t.calculateDepth(t.Root, 0)
}

// Returns the Max value of the tree
func (t *binaryTree[T]) Max() (T, bool) {
	ret := t.maximum(t.Root)
	if ret == t.NIL {
		return t.NIL.Key, false
	}

	return ret.Key, true
}

// Return the Min value of the tree
func (t *binaryTree[T]) Min() (T, bool) {
	ret := t.minimum(t.Root)
	if ret == t.NIL {
		return t.NIL.Key, false
	}

	return ret.Key, true
}

// LevelOrder returns the level order traversal of the tree
func (t *binaryTree[T]) LevelOrder() []T {
	traversal := make([]T, 0)
	t.levelOrderHelper(t.Root, &traversal)
	return traversal
}

// AccessNodesByLayer accesses nodes layer by layer (2-D array),  instead of printing the results as 1-D array.
func (t *binaryTree[T]) AccessNodesByLayer() [][]T {
	root := t.Root
	if root == t.NIL {
		return [][]T{}
	}
	var q []*Node[T]
	var n *Node[T]
	var idx = 0
	q = append(q, root)
	var res [][]T

	for len(q) != 0 {
		res = append(res, []T{})
		qLen := len(q)
		for i := 0; i < qLen; i++ {
			n, q = q[0], q[1:]
			res[idx] = append(res[idx], n.Key)
			if n.Left != t.NIL {
				q = append(q, n.Left)
			}
			if n.Right != t.NIL {
				q = append(q, n.Right)
			}
		}
		idx++
	}
	return res
}

// Print the tree horizontally
func (t *binaryTree[T]) Print() {
	t.printHelper(t.Root, "", false)
}

// Return the Predecessor of the node of Key
// if there is no predecessor, return default value of type T and false
// otherwise return the Key of predecessor and true
func (t *binaryTree[T]) Predecessor(key T) (T, bool) {
	node, ok := t.searchTreeHelper(t.Root, key)
	if !ok {
		var dummy T
		return dummy, ok
	}
	return t.predecessorHelper(node)
}

// Return the Successor of the node of Key
// if there is no successor, return default value of type T and false
// otherwise return the Key of successor and true
func (t *binaryTree[T]) Successor(key T) (T, bool) {
	node, ok := t.searchTreeHelper(t.Root, key)
	if !ok {
		var dummy T
		return dummy, ok
	}

	return t.successorHelper(node)
}

func (t *binaryTree[T]) searchTreeHelper(node *Node[T], key T) (*Node[T], bool) {
	if node == t.NIL {
		return node, false
	}

	if key == node.Key {
		return node, true
	}

	if key < node.Key {
		return t.searchTreeHelper(node.Left, key)
	}
	return t.searchTreeHelper(node.Right, key)
}

// The iterative inorder;
// The recursive way is similar to the preOrderRecursive
func (t *binaryTree[T]) inOrderHelper(node *Node[T]) []T {
	var stack []*Node[T]
	var ret []T

	for node != t.NIL || len(stack) > 0 {
		for node != t.NIL {
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

func (t *binaryTree[T]) preOrderRecursive(n *Node[T], traversal *[]T) {
	if n == t.NIL {
		return
	}

	*traversal = append(*traversal, n.Key)
	t.preOrderRecursive(n.Left, traversal)
	t.preOrderRecursive(n.Right, traversal)
}

func (t *binaryTree[T]) postOrderRecursive(n *Node[T], traversal *[]T) {
	if n == t.NIL {
		return
	}

	t.postOrderRecursive(n.Left, traversal)
	t.postOrderRecursive(n.Right, traversal)
	*traversal = append(*traversal, n.Key)
}

func (t *binaryTree[T]) calculateDepth(n *Node[T], depth int) int {
	if n == t.NIL {
		return depth
	}

	return max.Int(t.calculateDepth(n.Left, depth+1), t.calculateDepth(n.Right, depth+1))
}

// Returns the minimum value of node of the tree
func (t *binaryTree[T]) minimum(node *Node[T]) *Node[T] {
	if node == t.NIL {
		return node
	}

	for node.Left != t.NIL {
		node = node.Left
	}
	return node
}

// Returns the maximum value of node of the tree
func (t *binaryTree[T]) maximum(node *Node[T]) *Node[T] {
	if node == t.NIL {
		return node
	}

	for node.Right != t.NIL {
		node = node.Right
	}
	return node
}

func (t *binaryTree[T]) levelOrderHelper(root *Node[T], traversal *[]T) {
	var q []*Node[T] // queue
	var tmp *Node[T]

	q = append(q, root)

	for len(q) != 0 {
		tmp, q = q[0], q[1:]
		*traversal = append(*traversal, tmp.Key)
		if tmp.Left != t.NIL {
			q = append(q, tmp.Left)
		}

		if tmp.Right != t.NIL {
			q = append(q, tmp.Right)
		}
	}
}

// Reference: https://stackoverflow.com/a/51730733/15437172
func (t *binaryTree[T]) printHelper(root *Node[T], indent string, isLeft bool) {
	if root == t.NIL {
		return
	}

	fmt.Print(indent)
	if isLeft {
		fmt.Print("├──")
		indent += "│  "
	} else {
		fmt.Print("└──")
		indent += "   "
	}

	if t.isRB() {
		color := "Black"
		if root.Color == Red {
			color = "Red"
		}

		fmt.Println(root.Key, "(", color, ")")
	} else {
		fmt.Println(root.Key)
	}

	t.printHelper(root.Left, indent, true)
	t.printHelper(root.Right, indent, false)
}

// Determines the binary tree is RB, whose NIL != nil
func (t *binaryTree[T]) isRB() bool {
	return t.NIL != nil
}

func (t *binaryTree[T]) predecessorHelper(node *Node[T]) (T, bool) {
	if node.Left != t.NIL {
		return t.maximum(node.Left).Key, true
	}

	p := node.Parent
	for p != t.NIL && node == p.Left {
		node = p
		p = p.Parent
	}

	if p == t.NIL {
		var dummy T
		return dummy, false
	}
	return p.Key, true
}

func (t *binaryTree[T]) successorHelper(node *Node[T]) (T, bool) {
	if node.Right != t.NIL {
		return t.minimum(node.Right).Key, true
	}

	p := node.Parent
	for p != t.NIL && node == p.Right {
		node = p
		p = p.Parent
	}

	if p == t.NIL {
		var dummy T
		return dummy, false
	}
	return p.Key, true
}

func (t *binaryTree[T]) transplant(u, v *Node[T]) {
	switch {
	case u.Parent == t.NIL:
		t.Root = v
	case u == u.Parent.Left:
		u.Parent.Left = v
	default:
		u.Parent.Right = v
	}

	if t.isRB() || v != t.NIL {
		v.Parent = u.Parent
	}
}
