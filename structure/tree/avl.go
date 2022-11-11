package tree

import (
	"github.com/TheAlgorithms/Go/constraints"
	"github.com/TheAlgorithms/Go/math/max"
)

type AVL[T constraints.Ordered] struct {
	*binaryTree[T]
}

// Verify Interface Compliance for AVL
var _ Tree[int] = (*AVL[int])(nil)

// NewAVL create a novel AVL tree
func NewAVL[T constraints.Ordered]() *AVL[T] {
	return &AVL[T]{
		binaryTree: &binaryTree[T]{
			Root: nil,
			NIL:  nil,
		},
	}
}

// Push a chain of Node's into the AVL Tree
func (avl *AVL[T]) Push(keys ...T) {
	for _, k := range keys {
		avl.Root = avl.pushHelper(avl.Root, k)
	}
}

// Delete a Node from the AVL Tree
func (avl *AVL[T]) Delete(key T) bool {
	if !avl.Has(key) {
		return false
	}

	avl.Root = avl.deleteHelper(avl.Root, key)
	return true
}

func (avl *AVL[T]) pushHelper(root *Node[T], key T) *Node[T] {
	if root == avl.NIL {
		return &Node[T]{
			Key:    key,
			Left:   avl.NIL,
			Right:  avl.NIL,
			Parent: avl.NIL,
			Height: 1,
		}
	}

	switch {
	case key < root.Key:
		tmp := avl.pushHelper(root.Left, key)
		tmp.Parent = root
		root.Left = tmp
	case key > root.Key:
		tmp := avl.pushHelper(root.Right, key)
		tmp.Parent = root
		root.Right = tmp
	default:
		return root
	}

	// balance the tree
	root.Height = avl.height(root)
	bFactor := avl.balanceFactor(root)
	if bFactor > 1 {
		switch {
		case key < root.Left.Key:
			return avl.rightRotate(root)
		case key > root.Left.Key:
			root.Left = avl.leftRotate(root.Left)
			return avl.rightRotate(root)
		}
	}

	if bFactor < -1 {
		switch {
		case key > root.Right.Key:
			return avl.leftRotate(root)
		case key < root.Right.Key:
			root.Right = avl.rightRotate(root.Right)
			return avl.leftRotate(root)
		}
	}

	return root
}

func (avl *AVL[T]) deleteHelper(root *Node[T], key T) *Node[T] {
	if root == avl.NIL {
		return root
	}

	switch {
	case key < root.Key:
		tmp := avl.deleteHelper(root.Left, key)
		root.Left = tmp
		if tmp != avl.NIL {
			tmp.Parent = root
		}
	case key > root.Key:
		tmp := avl.deleteHelper(root.Right, key)
		root.Right = tmp
		if tmp != avl.NIL {
			tmp.Parent = root
		}
	default:
		if root.Left == avl.NIL || root.Right == avl.NIL {
			tmp := root.Left
			if root.Right != avl.NIL {
				tmp = root.Right
			}

			if tmp == avl.NIL {
				root = avl.NIL
			} else {
				tmp.Parent = root.Parent
				root = tmp
			}
		} else {
			tmp := avl.minimum(root.Right)
			root.Key = tmp.Key
			del := avl.deleteHelper(root.Right, tmp.Key)
			root.Right = del
			if del != avl.NIL {
				del.Parent = root
			}
		}
	}

	if root == avl.NIL {
		return root
	}

	// balance the tree
	root.Height = avl.height(root)
	bFactor := avl.balanceFactor(root)
	switch {
	case bFactor > 1:
		switch {
		case avl.balanceFactor(root.Left) >= 0:
			return avl.rightRotate(root)
		default:
			root.Left = avl.leftRotate(root.Left)
			return avl.rightRotate(root)
		}
	case bFactor < -1:
		switch {
		case avl.balanceFactor(root.Right) <= 0:
			return avl.leftRotate(root)
		default:
			root.Right = avl.rightRotate(root.Right)
			return avl.leftRotate(root)
		}
	}

	return root
}

func (avl *AVL[T]) height(root *Node[T]) int {
	if root == avl.NIL {
		return 0
	}

	var leftHeight, rightHeight int
	if root.Left != avl.NIL {
		leftHeight = root.Left.Height
	}
	if root.Right != avl.NIL {
		rightHeight = root.Right.Height
	}
	return 1 + max.Int(leftHeight, rightHeight)
}

// balanceFactor : negative balance factor means subtree Root is heavy toward Left
// and positive balance factor means subtree Root is heavy toward Right side
func (avl *AVL[T]) balanceFactor(root *Node[T]) int {
	var leftHeight, rightHeight int
	if root.Left != avl.NIL {
		leftHeight = root.Left.Height
	}
	if root.Right != avl.NIL {
		rightHeight = root.Right.Height
	}
	return leftHeight - rightHeight
}

func (avl *AVL[T]) leftRotate(x *Node[T]) *Node[T] {
	y := x.Right
	yl := y.Left
	y.Left = x
	x.Right = yl

	if yl != avl.NIL {
		yl.Parent = x
	}

	y.Parent = x.Parent
	x.Parent = y

	x.Height = avl.height(x)
	y.Height = avl.height(y)
	return y
}

func (avl *AVL[T]) rightRotate(x *Node[T]) *Node[T] {
	y := x.Left
	yr := y.Right
	y.Right = x
	x.Left = yr

	if yr != avl.NIL {
		yr.Parent = x
	}

	y.Parent = x.Parent
	x.Parent = y

	x.Height = avl.height(x)
	y.Height = avl.height(y)
	return y
}
