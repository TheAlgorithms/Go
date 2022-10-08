package avl

import (
	"testing"
)

func TestInsert(t *testing.T) {
	t.Run("LLRotation-Test", func(t *testing.T) {
		root := NewAVLTree()
		Insert(&root, 5)
		Insert(&root, 4)
		Insert(&root, 3)

		if root.key != 4 {
			t.Errorf("root should have value = 4")
		}
		if root.height != 2 {
			t.Errorf("height of root should be = 2")
		}

		if root.left.key != 3 {
			t.Errorf("left child should have value = 3")
		}
		if root.left.height != 1 {
			t.Errorf("height of left child should be 1")
		}

		if root.right.key != 5 {
			t.Errorf("right child should have value = 5")
		}
		if root.right.height != 1 {
			t.Errorf("height of right should be 1")
		}

	})
	t.Run("LRRotaion-Test", func(t *testing.T) {
		root := NewAVLTree()
		Insert(&root, 5)
		Insert(&root, 3)
		Insert(&root, 4)

		if root.key != 4 {
			t.Errorf("root should have value = 4")
		}
		if root.height != 2 {
			t.Errorf("height of root should be = 2")
		}

		if root.left.key != 3 {
			t.Errorf("left child should have value = 3")
		}
		if root.left.height != 1 {
			t.Errorf("height of left child should be 1")
		}

		if root.right.key != 5 {
			t.Errorf("right child should have value = 5")
		}
		if root.right.height != 1 {
			t.Errorf("height of right should be 1")
		}
	})

	t.Run("RRRotaion-Test", func(t *testing.T) {
		root := NewAVLTree()
		Insert(&root, 3)
		Insert(&root, 4)
		Insert(&root, 5)

		if root.key != 4 {
			t.Errorf("root should have value = 4")
		}
		if root.height != 2 {
			t.Errorf("height of root should be = 2")
		}

		if root.left.key != 3 {
			t.Errorf("left child should have value = 3")
		}
		if root.left.height != 1 {
			t.Errorf("height of left child should be 1")
		}

		if root.right.key != 5 {
			t.Errorf("right child should have value = 5")
		}
		if root.right.height != 1 {
			t.Errorf("height of right should be 1")
		}
	})
	t.Run("RLRotaion-Test", func(t *testing.T) {
		root := NewAVLTree()
		Insert(&root, 3)
		Insert(&root, 5)
		Insert(&root, 4)

		if root.key != 4 {
			t.Errorf("root should have value = 4")
		}
		if root.height != 2 {
			t.Errorf("height of root should be = 2")
		}

		if root.left.key != 3 {
			t.Errorf("left child should have value = 3")
		}
		if root.left.height != 1 {
			t.Errorf("height of left child should be 1")
		}

		if root.right.key != 5 {
			t.Errorf("right child should have value = 5")
		}
		if root.right.height != 1 {
			t.Errorf("height of right should be 1")
		}
	})
}

func TestDelete(t *testing.T) {
	t.Run("LLRotaion-Test", func(t *testing.T) {
		root := NewAVLTree()

		Insert(&root, 5)
		Insert(&root, 4)
		Insert(&root, 3)
		Insert(&root, 2)

		Delete(&root, 5)

		if root.key != 3 {
			t.Errorf("root should have value = 3")
		}
		if root.height != 2 {
			t.Errorf("height of root should be = 2")
		}

		if root.left.key != 2 {
			t.Errorf("left child should have value = 2")
		}
		if root.left.height != 1 {
			t.Errorf("height of left child should be 1")
		}

		if root.right.key != 4 {
			t.Errorf("right child should have value = 5")
		}
		if root.right.height != 1 {
			t.Errorf("height of right should be 1")
		}
	})

	t.Run("LRRotaion-Test", func(t *testing.T) {
		root := NewAVLTree()

		Insert(&root, 10)
		Insert(&root, 8)
		Insert(&root, 6)
		Insert(&root, 7)

		Delete(&root, 10)

		if root.key != 7 {
			t.Errorf("root should have value = 7")
		}
		if root.height != 2 {
			t.Errorf("height of root should be = 2")
		}

		if root.left.key != 6 {
			t.Errorf("left child should have value = 6")
		}
		if root.left.height != 1 {
			t.Errorf("height of left child should be 1")
		}

		if root.right.key != 8 {
			t.Errorf("right child should have value = 8")
		}
		if root.right.height != 1 {
			t.Errorf("height of right should be 1")
		}

	})

	t.Run("RRRotaion-Test", func(t *testing.T) {
		root := NewAVLTree()

		Insert(&root, 2)
		Insert(&root, 3)
		Insert(&root, 4)
		Insert(&root, 5)

		Delete(&root, 2)

		if root.key != 4 {
			t.Errorf("root should have value = 4")
		}
		if root.height != 2 {
			t.Errorf("height of root should be = 2")
		}

		if root.left.key != 3 {
			t.Errorf("left child should have value = 3")
		}
		if root.left.height != 1 {
			t.Errorf("height of left child should be 1")
		}

		if root.right.key != 5 {
			t.Errorf("right child should have value = 5")
		}
		if root.right.height != 1 {
			t.Errorf("height of right should be 1")
		}

	})

	t.Run("RLRotaion-Test", func(t *testing.T) {
		root := NewAVLTree()

		Insert(&root, 7)
		Insert(&root, 6)
		Insert(&root, 9)
		Insert(&root, 8)

		Delete(&root, 6)

		if root.key != 8 {
			t.Errorf("root should have value = 8")
		}
		if root.height != 2 {
			t.Errorf("height of root should be = 2")
		}

		if root.left.key != 7 {
			t.Errorf("left child should have value = 7")
		}
		if root.left.height != 1 {
			t.Errorf("height of left child should be 1")
		}

		if root.right.key != 9 {
			t.Errorf("right child should have value = 9")
		}
		if root.right.height != 1 {
			t.Errorf("height of right should be 1")
		}

	})
}

func TestGet(t *testing.T) {

	root := NewAVLTree()

	if Get(root, 5) != nil {
		t.Error("no item should exists in newly created AVL tree")
	}

	Insert(&root, 5)
	Insert(&root, 4)
	Insert(&root, 3)

	n := Get(root, 4)
	if n.key != 4 {
		t.Error("key should be 4")
	}
	if n.right.key != 5 {
		t.Error("right child should have value 5")
	}

	if n.left.key != 3 {
		t.Error("left child should have value 3")
	}

}
