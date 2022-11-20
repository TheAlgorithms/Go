package tree_test

import (
	"math/rand"
	"reflect"
	"sort"
	"testing"

	bt "github.com/TheAlgorithms/Go/structure/tree"
)

func TestTreeGetOrHas(t *testing.T) {
	helper := func(tree TestTree[int], nums []int) {
		tree.Push(nums...)
		for _, num := range nums {
			if !tree.Has(num) {
				t.Errorf("Error with Has or Push method")
			}
		}

		min, _ := tree.Min()
		max, _ := tree.Max()

		if _, ok := tree.Get(min - 1); ok {
			t.Errorf("Error with Get method")
		}

		if _, ok := tree.Get(max + 1); ok {
			t.Errorf("Error with Get method")
		}
	}

	lens := []int{100, 1_000, 10_000, 100_000}
	for _, ll := range lens {
		nums := rand.Perm(ll)
		t.Run("Test Binary Search Tree", func(t *testing.T) {
			bsTree := bt.NewBinarySearch[int]()
			helper(bsTree, nums)
		})

		t.Run("Test Red-Black Tree", func(t *testing.T) {
			rbTree := bt.NewRB[int]()
			helper(rbTree, nums)
		})

		t.Run("Test AVL Tree", func(t *testing.T) {
			avlTree := bt.NewAVL[int]()
			helper(avlTree, nums)
		})
	}
}

func TestTreePreOrder(t *testing.T) {
	t.Run("Test for Binary-Search Tree", func(t *testing.T) {
		tests := []struct {
			input []int
			want  []int
		}{
			{[]int{90, 80, 100, 70, 85, 95, 105}, []int{90, 80, 70, 85, 100, 95, 105}},
			{[]int{90, 80, 100, 70, 85, 95, 105, 1, 21, 31, 41, 51, 61, 71}, []int{90, 80, 70, 1, 21, 31, 41, 51, 61, 71, 85, 100, 95, 105}},
			{[]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}, []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}},
		}
		for _, tt := range tests {
			tree := bt.NewBinarySearch[int]()
			tree.Push(tt.input...)
			if ret := tree.PreOrder(); !reflect.DeepEqual(ret, tt.want) {
				t.Errorf("Error with PreOrder")
			}
		}
	})

	t.Run("Test for AVL Tree", func(t *testing.T) {
		tests := []struct {
			input []int
			want  []int
		}{
			{[]int{90, 80, 100, 70, 85, 95, 105}, []int{90, 80, 70, 85, 100, 95, 105}},
			{[]int{90, 80, 100, 70, 85, 95, 105, 1, 21, 31, 41, 51, 61, 71}, []int{70, 41, 21, 1, 31, 51, 61, 90, 80, 71, 85, 100, 95, 105}},
			{[]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}, []int{7, 3, 2, 1, 5, 4, 6, 9, 8, 10}},
		}
		for _, tt := range tests {
			tree := bt.NewAVL[int]()
			tree.Push(tt.input...)
			if ret := tree.PreOrder(); !reflect.DeepEqual(ret, tt.want) {
				t.Errorf("Error with PreOrder")
			}
		}
	})

	t.Run("Test for Red-Black Tree", func(t *testing.T) {
		tests := []struct {
			input []int
			want  []int
		}{
			{[]int{90, 80, 100, 70, 85, 95, 105}, []int{90, 80, 70, 85, 100, 95, 105}},
			{[]int{90, 80, 100, 70, 85, 95, 105, 1, 21, 31, 41, 51, 61, 71}, []int{80, 41, 21, 1, 31, 61, 51, 70, 71, 90, 85, 100, 95, 105}},
			{[]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}, []int{7, 5, 3, 2, 1, 4, 6, 9, 8, 10}},
		}
		for _, tt := range tests {
			tree := bt.NewRB[int]()
			tree.Push(tt.input...)
			if ret := tree.PreOrder(); !reflect.DeepEqual(ret, tt.want) {
				t.Errorf("Error with PreOrder")
			}
		}
	})
}

func TestTreeInOrder(t *testing.T) {
	lens := []int{100, 1_000, 10_000, 100_000}
	for _, ll := range lens {
		nums := rand.Perm(ll)

		t.Run("Test Binary Search Tree", func(t *testing.T) {
			bsTree := bt.NewBinarySearch[int]()
			bsTree.Push(nums...)
			if ret := bsTree.InOrder(); !sort.IntsAreSorted(ret) {
				t.Errorf("Error with InOrder")
			}
		})

		t.Run("Test Red-Black Tree", func(t *testing.T) {
			rbTree := bt.NewRB[int]()
			rbTree.Push(nums...)
			if ret := rbTree.InOrder(); !sort.IntsAreSorted(ret) {
				t.Errorf("Error with InOrder")
			}
		})

		t.Run("Test AVL Tree", func(t *testing.T) {
			avlTree := bt.NewAVL[int]()
			avlTree.Push(nums...)
			if ret := avlTree.InOrder(); !sort.IntsAreSorted(ret) {
				t.Errorf("Error with InOrder")
			}
		})
	}
}

func TestTreePostOrder(t *testing.T) {
	t.Run("Test for Binary-Search Tree", func(t *testing.T) {
		tests := []struct {
			input []int
			want  []int
		}{
			{[]int{90, 80, 100, 70, 85, 95, 105}, []int{70, 85, 80, 95, 105, 100, 90}},
			{[]int{90, 80, 100, 70, 85, 95, 105, 1, 21, 31, 41, 51, 61, 71},
				[]int{61, 51, 41, 31, 21, 1, 71, 70, 85, 80, 95, 105, 100, 90}},
			{[]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
		}
		for i, tt := range tests {
			tree := bt.NewBinarySearch[int]()
			tree.Push(tt.input...)
			if ret := tree.PostOrder(); !reflect.DeepEqual(ret, tt.want) {
				t.Errorf("#%d Error with Post", i)
			}
		}
	})

	t.Run("Test for AVL Tree", func(t *testing.T) {
		tests := []struct {
			input []int
			want  []int
		}{
			{[]int{90, 80, 100, 70, 85, 95, 105}, []int{70, 85, 80, 95, 105, 100, 90}},
			{[]int{90, 80, 100, 70, 85, 95, 105, 1, 21, 31, 41, 51, 61, 71},
				[]int{1, 31, 21, 61, 51, 41, 71, 85, 80, 95, 105, 100, 90, 70}},
			{[]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}, []int{1, 2, 4, 6, 5, 3, 8, 10, 9, 7}},
		}
		for i, tt := range tests {
			tree := bt.NewAVL[int]()
			tree.Push(tt.input...)
			if ret := tree.PostOrder(); !reflect.DeepEqual(ret, tt.want) {
				t.Errorf("#%d Error with PostOrder", i)
			}
		}
	})

	t.Run("Test for Red-Black Tree", func(t *testing.T) {
		tests := []struct {
			input []int
			want  []int
		}{
			{[]int{90, 80, 100, 70, 85, 95, 105}, []int{70, 85, 80, 95, 105, 100, 90}},
			{[]int{90, 80, 100, 70, 85, 95, 105, 1, 21, 31, 41, 51, 61, 71},
				[]int{1, 31, 21, 51, 71, 70, 61, 41, 85, 95, 105, 100, 90, 80}},
			{[]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}, []int{1, 2, 4, 3, 6, 5, 8, 10, 9, 7}},
		}
		for i, tt := range tests {
			tree := bt.NewRB[int]()
			tree.Push(tt.input...)
			if ret := tree.PostOrder(); !reflect.DeepEqual(ret, tt.want) {
				t.Errorf("#%d Error with PostOrder", i)
			}
		}
	})
}

func TestTreeLevelOrder(t *testing.T) {
	t.Run("Test for Binary-Search Tree", func(t *testing.T) {
		tests := []struct {
			input []int
			want  []int
		}{
			{[]int{90, 80, 100, 70, 85, 95, 105}, []int{90, 80, 100, 70, 85, 95, 105}},
			{[]int{90, 80, 100, 70, 85, 95, 105, 1, 21, 31, 41, 51, 61, 71},
				[]int{90, 80, 100, 70, 85, 95, 105, 1, 71, 21, 31, 41, 51, 61}},
			{[]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}, []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}},
		}
		for i, tt := range tests {
			tree := bt.NewBinarySearch[int]()
			tree.Push(tt.input...)
			if ret := tree.LevelOrder(); !reflect.DeepEqual(ret, tt.want) {
				t.Errorf("#%d Error with LevelOrder", i)
			}
		}
	})

	t.Run("Test for AVL Tree", func(t *testing.T) {
		tests := []struct {
			input []int
			want  []int
		}{
			{[]int{90, 80, 100, 70, 85, 95, 105}, []int{90, 80, 100, 70, 85, 95, 105}},
			{[]int{90, 80, 100, 70, 85, 95, 105, 1, 21, 31, 41, 51, 61, 71},
				[]int{70, 41, 90, 21, 51, 80, 100, 1, 31, 61, 71, 85, 95, 105}},
			{[]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}, []int{7, 3, 9, 2, 5, 8, 10, 1, 4, 6}},
		}
		for i, tt := range tests {
			tree := bt.NewAVL[int]()
			tree.Push(tt.input...)
			if ret := tree.LevelOrder(); !reflect.DeepEqual(ret, tt.want) {
				t.Errorf("#%d Error with LevelOrder", i)
			}
		}
	})

	t.Run("Test for Red-Black Tree", func(t *testing.T) {
		tests := []struct {
			input []int
			want  []int
		}{
			{[]int{90, 80, 100, 70, 85, 95, 105}, []int{90, 80, 100, 70, 85, 95, 105}},
			{[]int{90, 80, 100, 70, 85, 95, 105, 1, 21, 31, 41, 51, 61, 71},
				[]int{80, 41, 90, 21, 61, 85, 100, 1, 31, 51, 70, 95, 105, 71}},
			{[]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}, []int{7, 5, 9, 3, 6, 8, 10, 2, 4, 1}},
		}
		for i, tt := range tests {
			tree := bt.NewRB[int]()
			tree.Push(tt.input...)
			if ret := tree.LevelOrder(); !reflect.DeepEqual(ret, tt.want) {
				t.Errorf("#%d Error with LevelOrder %v", i, ret)
			}
		}
	})
}

func TestTreeMinAndMax(t *testing.T) {
	helper := func(tree TestTree[int], nums []int) {
		ll := len(nums)
		if _, ok := tree.Min(); ok {
			t.Errorf("Error with Min method.")
		}
		if _, ok := tree.Max(); ok {
			t.Errorf("Error with Max method.")
		}
		tree.Push(nums...)
		if min, ok := tree.Min(); !ok || min != nums[0] {
			t.Errorf("Error with Min method.")
		}
		if max, ok := tree.Max(); !ok || max != nums[ll-1] {
			t.Errorf("Error with Max method.")
		}
	}

	lens := []int{500, 1_000, 10_000}
	for _, ll := range lens {
		nums := rand.Perm(ll)
		sort.Ints(nums)

		t.Run("Test Binary Search Tree", func(t *testing.T) {
			helper(bt.NewBinarySearch[int](), nums)
		})

		t.Run("Test Red-Black Tree", func(t *testing.T) {
			helper(bt.NewRB[int](), nums)
		})

		t.Run("Test AVL Tree", func(t *testing.T) {
			helper(bt.NewAVL[int](), nums)
		})
	}
}

func TestTreeDepth(t *testing.T) {
	t.Run("Test for Binary-Search Tree", func(t *testing.T) {
		tests := []struct {
			input []int
			want  int
		}{
			{[]int{}, 0},
			{[]int{90, 80, 100, 70, 85, 95, 105}, 3},
			{[]int{90, 80, 100, 70, 85, 95, 105, 1, 21, 31, 41, 51, 61, 71}, 9},
			{[]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}, 10},
		}
		for _, tt := range tests {
			tree := bt.NewBinarySearch[int]()
			tree.Push(tt.input...)
			if ret := tree.Depth(); ret != tt.want {
				t.Errorf("Error with Depth")
			}
		}
	})

	t.Run("Test for AVL Tree", func(t *testing.T) {
		tests := []struct {
			input []int
			want  int
		}{
			{[]int{}, 0},
			{[]int{90, 80, 100, 70, 85, 95, 105}, 3},
			{[]int{90, 80, 100, 70, 85, 95, 105, 1, 21, 31, 41, 51, 61, 71}, 4},
			{[]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}, 4},
		}
		for i, tt := range tests {
			tree := bt.NewAVL[int]()
			tree.Push(tt.input...)
			if ret := tree.Depth(); ret != tt.want {
				t.Errorf("#%d Error with Depth", i)
			}
		}
	})

	t.Run("Test for Red-Black Tree", func(t *testing.T) {
		tests := []struct {
			input []int
			want  int
		}{
			{[]int{}, 0},
			{[]int{90, 80, 100, 70, 85, 95, 105}, 3},
			{[]int{90, 80, 100, 70, 85, 95, 105, 1, 21, 31, 41, 51, 61, 71}, 5},
			{[]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}, 5},
		}
		for i, tt := range tests {
			tree := bt.NewRB[int]()
			tree.Push(tt.input...)
			if ret := tree.Depth(); ret != tt.want {
				t.Errorf("#%d Error with Depth", i)
			}
		}
	})
}

func TestTreeAccessNodesByLayer(t *testing.T) {
	t.Run("Test for Binary-Search Tree", func(t *testing.T) {
		tests := []struct {
			input []int
			want  [][]int
		}{
			{[]int{90, 80, 100, 70, 85, 95, 105}, [][]int{{90}, {80, 100}, {70, 85, 95, 105}}},
			{[]int{90, 80, 100, 70, 85, 95, 105, 1, 21, 31, 41, 51, 61, 71},
				[][]int{{90}, {80, 100}, {70, 85, 95, 105}, {1, 71}, {21}, {31}, {41}, {51}, {61}}},
			{[]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}, [][]int{{10}, {9}, {8}, {7}, {6}, {5}, {4}, {3}, {2}, {1}}},
			{[]int{}, [][]int{}},
		}
		for i, tt := range tests {
			tree := bt.NewBinarySearch[int]()
			tree.Push(tt.input...)
			if ret := tree.AccessNodesByLayer(); !reflect.DeepEqual(ret, tt.want) {
				t.Errorf("#%d Error with AccessNoedsByLayer", i)
			}
		}
	})

	t.Run("Test for AVL Tree", func(t *testing.T) {
		tests := []struct {
			input []int
			want  [][]int
		}{
			{[]int{90, 80, 100, 70, 85, 95, 105}, [][]int{{90}, {80, 100}, {70, 85, 95, 105}}},
			{[]int{90, 80, 100, 70, 85, 95, 105, 1, 21, 31, 41, 51, 61, 71},
				[][]int{{70}, {41, 90}, {21, 51, 80, 100}, {1, 31, 61, 71, 85, 95, 105}}},
			{[]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}, [][]int{{7}, {3, 9}, {2, 5, 8, 10}, {1, 4, 6}}},
			{[]int{}, [][]int{}},
		}
		for i, tt := range tests {
			tree := bt.NewAVL[int]()
			tree.Push(tt.input...)
			if ret := tree.AccessNodesByLayer(); !reflect.DeepEqual(ret, tt.want) {
				t.Errorf("#%d Error with AccessNoedsByLayer", i)
			}
		}
	})

	t.Run("Test for Red-Black Tree", func(t *testing.T) {
		tests := []struct {
			input []int
			want  [][]int
		}{
			{[]int{90, 80, 100, 70, 85, 95, 105}, [][]int{{90}, {80, 100}, {70, 85, 95, 105}}},
			{[]int{90, 80, 100, 70, 85, 95, 105, 1, 21, 31, 41, 51, 61, 71},
				[][]int{{80}, {41, 90}, {21, 61, 85, 100}, {1, 31, 51, 70, 95, 105}, {71}}},
			{[]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}, [][]int{{7}, {5, 9}, {3, 6, 8, 10}, {2, 4}, {1}}},
			{[]int{}, [][]int{}},
		}
		for i, tt := range tests {
			tree := bt.NewRB[int]()
			t.Log(reflect.TypeOf(tree).String() == "*tree.RB[int]")
			tree.Push(tt.input...)
			if ret := tree.AccessNodesByLayer(); !reflect.DeepEqual(ret, tt.want) {
				t.Errorf("#%d Error with AccessNoedsByLayer, %v", i, ret)
			}
		}
	})
}

func TestTreePredecessorAndSuccessor(t *testing.T) {
	helper := func(tree TestTree[int]) {
		nums := []int{10, 8, 88, 888, 4, -1, 100}
		tree.Push(nums...)
		if ret, ok := tree.Predecessor(100); !ok || ret != 88 {
			t.Error("Error with Predecessor")
		}

		if _, ok := tree.Predecessor(-1); ok {
			t.Error("Error with Predecessor")
		}

		tree.Push(-100)
		if ret, ok := tree.Predecessor(-1); !ok || ret != -100 {
			t.Error("Error with Predecessor")
		}

		if _, ok := tree.Predecessor(-12); ok {
			t.Error("Error with Predecessor")
		}

		if ret, ok := tree.Predecessor(4); !ok || ret != -1 {
			t.Error("Error with Predecessor")
		}

		if ret, ok := tree.Successor(4); !ok || ret != 8 {
			t.Error("Error with Successor")
		}

		if ret, ok := tree.Successor(8); !ok || ret != 10 {
			t.Error("Error with Successor")
		}

		if ret, ok := tree.Successor(88); !ok || ret != 100 {
			t.Error("Error with Successor")
		}

		if ret, ok := tree.Successor(100); !ok || ret != 888 {
			t.Error("Error with Successor")
		}

		tree.Delete(888)
		if _, ok := tree.Successor(100); ok {
			t.Error("Error with Successor")
		}

		if ret, ok := tree.Successor(-1); !ok || ret != 4 {
			t.Error("Error with Successor")
		}

		if _, ok := tree.Successor(888); ok {
			t.Error("Error with Successor")
		}

		if _, ok := tree.Successor(188); ok {
			t.Error("Error with Successor")
		}
	}

	t.Run("Test for Binary Search Tree", func(t *testing.T) {
		tree := bt.NewBinarySearch[int]()
		helper(tree)
	})

	t.Run("Test for Red-Black Tree", func(t *testing.T) {
		tree := bt.NewRB[int]()
		helper(tree)
	})

	t.Run("Test for AVL Tree", func(t *testing.T) {
		tree := bt.NewAVL[int]()
		helper(tree)
	})
}

// Benchmark the comparisons between BST, AVL and RB Tree
const testNum = 10_000

func BenchmarkBSTree_Insert(b *testing.B) {
	helper := func() {
		tree := bt.NewBinarySearch[int]()
		for i := 1; i <= testNum; i++ {
			tree.Push(i)
		}
	}

	for i := 0; i < b.N; i++ {
		helper()
	}
}

func BenchmarkBSTree_Has(b *testing.B) {
	helper := func() {
		tree := bt.NewBinarySearch[int]()
		for i := 1; i <= testNum; i++ {
			tree.Push(i)
		}

		for i := 1; i <= testNum; i++ {
			tree.Has(i)
		}
	}

	for i := 0; i < b.N; i++ {
		helper()
	}
}

func BenchmarkBSTree_Delete(b *testing.B) {
	helper := func() {
		tree := bt.NewBinarySearch[int]()
		for i := 1; i <= testNum; i++ {
			tree.Push(i)
		}

		for i := 1; i <= testNum; i++ {
			tree.Delete(i)
		}
	}

	for i := 0; i < b.N; i++ {
		helper()
	}
}

func BenchmarkRBTree_Insert(b *testing.B) {
	helper := func() {
		tree := bt.NewRB[int]()
		for i := 1; i <= testNum; i++ {
			tree.Push(i)
		}
	}

	for i := 0; i < b.N; i++ {
		helper()
	}
}

func BenchmarkRBTree_Has(b *testing.B) {
	helper := func() {
		tree := bt.NewRB[int]()
		for i := 1; i <= testNum; i++ {
			tree.Push(i)
		}

		for i := 1; i <= testNum; i++ {
			tree.Has(i)
		}
	}

	for i := 0; i < b.N; i++ {
		helper()
	}
}

func BenchmarkRBTree_Delete(b *testing.B) {
	helper := func() {
		tree := bt.NewRB[int]()
		for i := 1; i <= testNum; i++ {
			tree.Push(i)
		}

		for i := 1; i <= testNum; i++ {
			tree.Delete(i)
		}
	}

	for i := 0; i < b.N; i++ {
		helper()
	}
}

func BenchmarkAVLTree_Insert(b *testing.B) {
	helper := func() {
		tree := bt.NewAVL[int]()
		for i := 1; i <= testNum; i++ {
			tree.Push(i)
		}
	}

	for i := 0; i < b.N; i++ {
		helper()
	}

}

func BenchmarkAVLTree_Has(b *testing.B) {
	helper := func() {
		tree := bt.NewAVL[int]()
		for i := 1; i <= testNum; i++ {
			tree.Push(i)
		}

		for i := 1; i <= testNum; i++ {
			tree.Has(i)
		}
	}

	for i := 0; i < b.N; i++ {
		helper()
	}
}

func BenchmarkAVLTree_Delete(b *testing.B) {
	helper := func() {
		tree := bt.NewAVL[int]()
		for i := 1; i <= testNum; i++ {
			tree.Push(i)
		}

		for i := 1; i <= testNum; i++ {
			tree.Delete(i)
		}
	}

	for i := 0; i < b.N; i++ {
		helper()
	}
}
