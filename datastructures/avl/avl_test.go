package avl

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsert(t *testing.T) {
	is := assert.New(t)
	t.Run("LLRotaion-Test", func(t *testing.T) {
		var root *Node
		Insert(&root, 5)
		Insert(&root, 4)
		Insert(&root, 3)

		is.Equal(4, root.Key)
		is.Equal(2, root.Height)

		is.Equal(3, root.Left.Key)
		is.Equal(1, root.Left.Height)

		is.Equal(5, root.Right.Key)
		is.Equal(1, root.Right.Height)
	})
	t.Run("LRRotaion-Test", func(t *testing.T) {
		var root *Node
		Insert(&root, 5)
		Insert(&root, 3)
		Insert(&root, 4)

		is.Equal(4, root.Key)
		is.Equal(2, root.Height)

		is.Equal(3, root.Left.Key)
		is.Equal(1, root.Left.Height)

		is.Equal(5, root.Right.Key)
		is.Equal(1, root.Right.Height)
	})

	t.Run("RRRotaion-Test", func(t *testing.T) {
		var root *Node
		Insert(&root, 3)
		Insert(&root, 4)
		Insert(&root, 5)

		is.Equal(4, root.Key)
		is.Equal(2, root.Height)

		is.Equal(3, root.Left.Key)
		is.Equal(1, root.Left.Height)

		is.Equal(5, root.Right.Key)
		is.Equal(1, root.Right.Height)
	})
	t.Run("RLRotaion-Test", func(t *testing.T) {
		var root *Node
		Insert(&root, 3)
		Insert(&root, 5)
		Insert(&root, 4)

		is.Equal(4, root.Key)
		is.Equal(2, root.Height)

		is.Equal(3, root.Left.Key)
		is.Equal(1, root.Left.Height)

		is.Equal(5, root.Right.Key)
		is.Equal(1, root.Right.Height)
	})
}

func TestDelete(t *testing.T) {
	is := assert.New(t)
	t.Run("LLRotaion-Test", func(t *testing.T) {
		var root *Node

		Insert(&root, 5)
		Insert(&root, 4)
		Insert(&root, 3)
		Insert(&root, 2)

		delete(&root, 5)

		is.Equal(3, root.Key)
		is.Equal(2, root.Height)

		is.Equal(2, root.Left.Key)
		is.Equal(1, root.Left.Height)

		is.Equal(4, root.Right.Key)
		is.Equal(1, root.Right.Height)
	})

	t.Run("LRRotaion-Test", func(t *testing.T) {
		var root *Node

		Insert(&root, 10)
		Insert(&root, 8)
		Insert(&root, 6)
		Insert(&root, 7)

		delete(&root, 10)

		is.Equal(7, root.Key)
		is.Equal(2, root.Height)

		is.Equal(6, root.Left.Key)
		is.Equal(1, root.Left.Height)

		is.Equal(8, root.Right.Key)
		is.Equal(1, root.Right.Height)
	})

	t.Run("RRRotaion-Test", func(t *testing.T) {
		var root *Node

		Insert(&root, 2)
		Insert(&root, 3)
		Insert(&root, 4)
		Insert(&root, 5)

		delete(&root, 2)

		is.Equal(4, root.Key)
		is.Equal(2, root.Height)

		is.Equal(3, root.Left.Key)
		is.Equal(1, root.Left.Height)

		is.Equal(5, root.Right.Key)
		is.Equal(1, root.Right.Height)
	})

	t.Run("RLRotaion-Test", func(t *testing.T) {
		var root *Node

		Insert(&root, 7)
		Insert(&root, 6)
		Insert(&root, 9)
		Insert(&root, 8)

		delete(&root, 6)

		is.Equal(8, root.Key)
		is.Equal(2, root.Height)

		is.Equal(7, root.Left.Key)
		is.Equal(1, root.Left.Height)

		is.Equal(9, root.Right.Key)
		is.Equal(1, root.Right.Height)
	})
}

func TestGet(t *testing.T) {
	is := assert.New(t)

	var root *Node

	is.Nil(Get(root, 5))

	Insert(&root, 5)
	Insert(&root, 4)
	Insert(&root, 3)

	n := Get(root, 4)
	is.Equal(4, n.Key)
	is.Equal(5, n.Right.Key)
	is.Equal(3, n.Left.Key)
}
