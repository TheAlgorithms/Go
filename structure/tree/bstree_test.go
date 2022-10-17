package tree_test

import (
	"testing"

	bt "github.com/TheAlgorithms/Go/structure/tree"
)

func TestInsert(t *testing.T) {
	bst := bt.NewBinarySearch[int]()

	bst.Push(90)
	bst.Push(80)
	bst.Push(100)

	if bst.Root.Key != 90 {
		t.Errorf("Root should have value = 90")
	}

	if bst.Root.Left.Key != 80 {
		t.Errorf("Left child should have value = 80")
	}

	if bst.Root.Right.Key != 100 {
		t.Errorf("Right child should have value = 100")
	}

	if bst.Depth() != 2 {
		t.Errorf("tree should have depth = 1")
	}
}

func TestDelete(t *testing.T) {
	t.Run("Delete a node with no child", func(t *testing.T) {
		bst := bt.NewBinarySearch[int]()

		bst.Push(90)
		bst.Push(80)
		bst.Push(100)

		bst.Delete(100)

		root := bst.Root
		if root.Key != 90 {
			t.Errorf("Root should have value = 90")
		}

		if root.Left.Key != 80 {
			t.Errorf("Left child should have value = 80")
		}

		if root.Right != nil {
			t.Errorf("Right child should have value = nil")
		}

		if bst.Depth() != 2 {
			t.Errorf("Depth should have value = 2")
		}

		bst.Delete(80)

		if root.Key != 90 {
			t.Errorf("Root should have value = 90")
		}

		if root.Left != nil {
			t.Errorf("Left child should have value = nil")
		}

		if bst.Depth() != 1 {
			t.Errorf("Depth should have value = 1")
		}
	})

	t.Run("Delete a node with one child", func(t *testing.T) {
		bst := bt.NewBinarySearch[int]()

		bst.Push(90)
		bst.Push(80)
		bst.Push(100)
		bst.Push(70)

		bst.Delete(80)

		root := bst.Root
		if root.Key != 90 {
			t.Errorf("Root should have value = 90")
		}

		if root.Right.Key != 100 {
			t.Errorf("Right child should have value = 100")
		}

		if root.Left.Key != 70 {
			t.Errorf("Left child should have value = 70")
		}

		if bst.Depth() != 2 {
			t.Errorf("Depth should have value = 2")
		}
	})

	t.Run("Delete a node with two children", func(t *testing.T) {
		bst := bt.NewBinarySearch[int]()

		bst.Push(90)
		bst.Push(80)
		bst.Push(100)
		bst.Push(70)
		bst.Push(85)

		bst.Delete(80)

		root := bst.Root
		if root.Key != 90 {
			t.Errorf("Root should have value = 90")
		}

		if root.Left.Key != 85 {
			t.Errorf("Left child should have value = 85")
		}

		if root.Right.Key != 100 {
			t.Errorf("Right child should have value = 100")
		}

		if bst.Depth() != 3 {
			t.Errorf("Depth should have value = 3")
		}
	})
}

// func TestAccessNodesByLayer(t *testing.T) {
// 	bst := bt.NewBinarySearch[int]()
//
// 	bst.Push(90)
// 	bst.Push(80)
// 	bst.Push(100)
// 	bst.Push(70)
// 	bst.Push(85)
// 	bst.Push(95)
// 	bst.Push(105)
//
// 	a := bst.AccessNodesByLayer()
// 	b := [][]int{{90}, {80, 100}, {70, 85, 95, 105}}
//
// 	if !reflect.DeepEqual(a, b) {
// 		t.Errorf("Nodes should have value = [[90] [80 100] [70 85 95 105]]")
// 	}
// }
