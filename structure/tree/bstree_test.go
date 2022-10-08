package tree

import (
	"reflect"
	"testing"
)

func TestInsert(t *testing.T) {
	BSTree := BSTree{
		Root: NewNode(90),
	}

	root := BSTree.Root

	Insert(root, 80)
	Insert(root, 100)

	if root.key != 90 {
		t.Errorf("root should have value = 90")
	}

	if root.left.key != 80 {
		t.Errorf("left child should have value = 80")
	}

	if root.right.key != 100 {
		t.Errorf("right child should have value = 100")
	}

	if BSTree.Depth() != 2 {
		t.Errorf("tree should have depth = 1")
	}
}

func TestDelete(t *testing.T) {
	t.Run("Delete a node with no child", func(t *testing.T) {
		BSTree := BSTree{
			Root: NewNode(90),
		}

		root := BSTree.Root

		Insert(root, 80)
		Insert(root, 100)

		BstDelete(root, 100)

		if root.key != 90 {
			t.Errorf("root should have value = 90")
		}

		if root.left.key != 80 {
			t.Errorf("left child should have value = 80")
		}

		if root.right != nil {
			t.Errorf("right child should have value = nil")
		}

		if BSTree.Depth() != 2 {
			t.Errorf("Depth should have value = 2")
		}
	})

	t.Run("Delete a node with one child", func(t *testing.T) {
		BSTree := BSTree{
			Root: NewNode(90),
		}

		root := BSTree.Root

		Insert(root, 80)
		Insert(root, 100)
		Insert(root, 70)

		BstDelete(root, 80)

		if root.key != 90 {
			t.Errorf("root should have value = 90")
		}

		if root.right.key != 100 {
			t.Errorf("right child should have value = 100")
		}

		if root.left.key != 70 {
			t.Errorf("left child should have value = 70")
		}

		if BSTree.Depth() != 2 {
			t.Errorf("Depth should have value = 2")
		}
	})

	t.Run("Delete a node with two children", func(t *testing.T) {
		BSTree := BSTree{
			Root: NewNode(90),
		}

		root := BSTree.Root

		Insert(root, 80)
		Insert(root, 100)
		Insert(root, 70)
		Insert(root, 85)

		BstDelete(root, 80)

		if root.key != 90 {
			t.Errorf("root should have value = 90")
		}

		if root.left.key != 85 {
			t.Errorf("left child should have value = 85")
		}

		if root.right.key != 100 {
			t.Errorf("right child should have value = 100")
		}

		if BSTree.Depth() != 3 {
			t.Errorf("Depth should have value = 3")
		}
	})
}

func TestInOrder(t *testing.T) {
	BSTree := BSTree{
		Root: NewNode(90),
	}

	root := BSTree.Root

	Insert(root, 80)
	Insert(root, 100)
	Insert(root, 70)
	Insert(root, 85)
	Insert(root, 95)
	Insert(root, 105)

	a := InOrder(root)
	b := []int{70, 80, 85, 90, 95, 100, 105}

	if !reflect.DeepEqual(a, b) {
		t.Errorf("Nodes should have value = [70 80 85 90 95 100 105]")
	}
}

func TestPreOrder(t *testing.T) {
	BSTree := BSTree{
		Root: NewNode(90),
	}

	root := BSTree.Root

	Insert(root, 80)
	Insert(root, 100)
	Insert(root, 70)
	Insert(root, 85)
	Insert(root, 95)
	Insert(root, 105)

	a := PreOrder(root)
	b := []int{90, 80, 70, 85, 100, 95, 105}

	if !reflect.DeepEqual(a, b) {
		t.Errorf("Nodes should have value = [90 80 70 85 100 95 105]")
	}
}

func TestPostOrder(t *testing.T) {
	BSTree := BSTree{
		Root: NewNode(90),
	}

	root := BSTree.Root

	Insert(root, 80)
	Insert(root, 100)
	Insert(root, 70)
	Insert(root, 85)
	Insert(root, 95)
	Insert(root, 105)

	a := PostOrder(root)
	b := []int{70, 85, 80, 95, 105, 100, 90}

	if !reflect.DeepEqual(a, b) {
		t.Errorf("Nodes should have value = [70 85 80 95 105 100 90]")
	}
}

func TestLevelOrder(t *testing.T) {
	BSTree := BSTree{
		Root: NewNode(90),
	}

	root := BSTree.Root

	Insert(root, 80)
	Insert(root, 100)
	Insert(root, 70)
	Insert(root, 85)
	Insert(root, 95)
	Insert(root, 105)

	a := LevelOrder(root)
	b := []int{90, 80, 100, 70, 85, 95, 105}

	if !reflect.DeepEqual(a, b) {
		t.Errorf("Nodes should have value = [90 80 100 70 85 95 105]")
	}
}

func TestAccessNodesByLayer(t *testing.T) {
	BSTree := BSTree{
		Root: NewNode(90),
	}

	root := BSTree.Root

	Insert(root, 80)
	Insert(root, 100)
	Insert(root, 70)
	Insert(root, 85)
	Insert(root, 95)
	Insert(root, 105)

	a := AccessNodesByLayer(root)
	b := [][]int{{90}, {80, 100}, {70, 85, 95, 105}}

	if !reflect.DeepEqual(a, b) {
		t.Errorf("Nodes should have value = [[90] [80 100] [70 85 95 105]]")
	}
}
