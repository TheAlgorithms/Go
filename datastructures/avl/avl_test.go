package avl

import (
	"testing"
)

func TestInsert(t *testing.T) {
	t.Run("LLRotaion-Test", func(t *testing.T) {
		root := NewTree()
		Insert(&root, 5)
		Insert(&root, 4)
		Insert(&root, 3)

		if root.Key != 4 {
			t.Errorf("root should have value = 4")
		}
		if root.Height != 2 {
			t.Errorf("height of root should be = 2")
		}

		if root.Left.Key != 3 {
			t.Errorf("left child should have value = 3")
		}
		if root.Left.Height != 1 {
			t.Errorf("height of left child should be 1")
		}

		if root.Right.Key != 5 {
			t.Errorf("right child should have value = 5")
		}
		if root.Right.Height != 1 {
			t.Errorf("height of right should be 1")
		}

	})
	t.Run("LRRotaion-Test", func(t *testing.T) {
		root := NewTree()
		Insert(&root, 5)
		Insert(&root, 3)
		Insert(&root, 4)

		if root.Key != 4 {
			t.Errorf("root should have value = 4")
		}
		if root.Height != 2 {
			t.Errorf("height of root should be = 2")
		}

		if root.Left.Key != 3 {
			t.Errorf("left child should have value = 3")
		}
		if root.Left.Height != 1 {
			t.Errorf("height of left child should be 1")
		}

		if root.Right.Key != 5 {
			t.Errorf("right child should have value = 5")
		}
		if root.Right.Height != 1 {
			t.Errorf("height of right should be 1")
		}
	})

	t.Run("RRRotaion-Test", func(t *testing.T) {
		root := NewTree()
		Insert(&root, 3)
		Insert(&root, 4)
		Insert(&root, 5)

		if root.Key != 4 {
			t.Errorf("root should have value = 4")
		}
		if root.Height != 2 {
			t.Errorf("height of root should be = 2")
		}

		if root.Left.Key != 3 {
			t.Errorf("left child should have value = 3")
		}
		if root.Left.Height != 1 {
			t.Errorf("height of left child should be 1")
		}

		if root.Right.Key != 5 {
			t.Errorf("right child should have value = 5")
		}
		if root.Right.Height != 1 {
			t.Errorf("height of right should be 1")
		}
	})
	t.Run("RLRotaion-Test", func(t *testing.T) {
		root := NewTree()
		Insert(&root, 3)
		Insert(&root, 5)
		Insert(&root, 4)

		if root.Key != 4 {
			t.Errorf("root should have value = 4")
		}
		if root.Height != 2 {
			t.Errorf("height of root should be = 2")
		}

		if root.Left.Key != 3 {
			t.Errorf("left child should have value = 3")
		}
		if root.Left.Height != 1 {
			t.Errorf("height of left child should be 1")
		}

		if root.Right.Key != 5 {
			t.Errorf("right child should have value = 5")
		}
		if root.Right.Height != 1 {
			t.Errorf("height of right should be 1")
		}
	})
}

func TestDelete(t *testing.T) {
	t.Run("LLRotaion-Test", func(t *testing.T) {
		root := NewTree()

		Insert(&root, 5)
		Insert(&root, 4)
		Insert(&root, 3)
		Insert(&root, 2)

		delete(&root, 5)

		if root.Key != 3 {
			t.Errorf("root should have value = 3")
		}
		if root.Height != 2 {
			t.Errorf("height of root should be = 2")
		}

		if root.Left.Key != 2 {
			t.Errorf("left child should have value = 2")
		}
		if root.Left.Height != 1 {
			t.Errorf("height of left child should be 1")
		}

		if root.Right.Key != 4 {
			t.Errorf("right child should have value = 5")
		}
		if root.Right.Height != 1 {
			t.Errorf("height of right should be 1")
		}
	})

	t.Run("LRRotaion-Test", func(t *testing.T) {
		root := NewTree()

		Insert(&root, 10)
		Insert(&root, 8)
		Insert(&root, 6)
		Insert(&root, 7)

		delete(&root, 10)

		if root.Key != 7 {
			t.Errorf("root should have value = 7")
		}
		if root.Height != 2 {
			t.Errorf("height of root should be = 2")
		}

		if root.Left.Key != 6 {
			t.Errorf("left child should have value = 6")
		}
		if root.Left.Height != 1 {
			t.Errorf("height of left child should be 1")
		}

		if root.Right.Key != 8 {
			t.Errorf("right child should have value = 8")
		}
		if root.Right.Height != 1 {
			t.Errorf("height of right should be 1")
		}

	})

	t.Run("RRRotaion-Test", func(t *testing.T) {
		root := NewTree()

		Insert(&root, 2)
		Insert(&root, 3)
		Insert(&root, 4)
		Insert(&root, 5)

		delete(&root, 2)

		if root.Key != 4 {
			t.Errorf("root should have value = 4")
		}
		if root.Height != 2 {
			t.Errorf("height of root should be = 2")
		}

		if root.Left.Key != 3 {
			t.Errorf("left child should have value = 3")
		}
		if root.Left.Height != 1 {
			t.Errorf("height of left child should be 1")
		}

		if root.Right.Key != 5 {
			t.Errorf("right child should have value = 5")
		}
		if root.Right.Height != 1 {
			t.Errorf("height of right should be 1")
		}

	})

	t.Run("RLRotaion-Test", func(t *testing.T) {
		root := NewTree()

		Insert(&root, 7)
		Insert(&root, 6)
		Insert(&root, 9)
		Insert(&root, 8)

		delete(&root, 6)

		if root.Key != 8 {
			t.Errorf("root should have value = 8")
		}
		if root.Height != 2 {
			t.Errorf("height of root should be = 2")
		}

		if root.Left.Key != 7 {
			t.Errorf("left child should have value = 7")
		}
		if root.Left.Height != 1 {
			t.Errorf("height of left child should be 1")
		}

		if root.Right.Key != 9 {
			t.Errorf("right child should have value = 9")
		}
		if root.Right.Height != 1 {
			t.Errorf("height of right should be 1")
		}

	})
}

func TestGet(t *testing.T) {

	root := NewTree()

	if Get(root, 5) != nil {
		t.Error("no item should exists in newly created AVL tree")
	}

	Insert(&root, 5)
	Insert(&root, 4)
	Insert(&root, 3)

	n := Get(root, 4)
	if n.Key != 4 {
		t.Error("key should be 4")
	}
	if n.Right.Key != 5 {
		t.Error("right child should have value 5")
	}

	if n.Left.Key != 3 {
		t.Error("left child should have value 3")
	}

}
