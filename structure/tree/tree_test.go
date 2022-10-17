package tree_test

import (
	"math/rand"
	"reflect"
	"sort"
	"testing"

	bt "github.com/TheAlgorithms/Go/structure/tree"
)

func TestTreeGetOrHas(t *testing.T) {
	lens := []int{100, 1_000, 10_000, 100_000}
	for _, ll := range lens {
		nums := rand.Perm(ll)
		t.Run("Test Binary Search Tree", func(t *testing.T) {
			bsTree := bt.NewBinarySearch[int]()
			bsTree.Push(nums...)
			for _, num := range nums {
				if !bsTree.Has(num) {
					t.Errorf("Error with Has or Push method")
				}
			}
			min, _ := bsTree.Min()
			max, _ := bsTree.Max()

			if _, ok := bsTree.Get(min - 1); ok {
				t.Errorf("Error with Get method")
			}

			if _, ok := bsTree.Get(max + 1); ok {
				t.Errorf("Error with Get method")
			}
		})

		t.Run("Test Red-Black Tree", func(t *testing.T) {
			rbTree := bt.NewRB[int]()
			rbTree.Push(nums...)
			for _, num := range nums {
				if !rbTree.Has(num) {
					t.Errorf("Error with Has or Push method")
				}
			}
		})

		t.Run("Test AVL Tree", func(t *testing.T) {
			avlTree := bt.NewAVL[int]()
			avlTree.Push(nums...)
			for _, num := range nums {
				if !avlTree.Has(num) {
					t.Errorf("Error with Has or Push method")
				}
			}
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
	lens := []int{500, 1_000, 10_000}
	for _, ll := range lens {
		nums := rand.Perm(ll)
		sort.Ints(nums)

		t.Run("Test Binary Search Tree", func(t *testing.T) {
			bsTree := bt.NewBinarySearch[int]()
			bsTree.Push(nums...)
			if min, ok := bsTree.Min(); !ok || min != nums[0] {
				t.Errorf("Error with Min method.")
			}
			if max, ok := bsTree.Max(); !ok || max != nums[ll-1] {
				t.Errorf("Error with Max method.")
			}
		})

		t.Run("Test Red-Black Tree", func(t *testing.T) {
			rbTree := bt.NewRB[int]()
			rbTree.Push(nums...)
			if min, ok := rbTree.Min(); !ok || min != nums[0] {
				t.Errorf("Error with Min method.")
			}
			if max, ok := rbTree.Max(); !ok || max != nums[ll-1] {
				t.Errorf("Error with Max method.")
			}
		})

		t.Run("Test AVL Tree", func(t *testing.T) {
			avlTree := bt.NewAVL[int]()
			avlTree.Push(nums...)
			if min, ok := avlTree.Min(); !ok || min != nums[0] {
				t.Errorf("Error with Min method.")
			}

			if max, ok := avlTree.Max(); !ok || max != nums[ll-1] {
				t.Errorf("Error with Max method.")
			}
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

func TestTreePrint(t *testing.T) {
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
		for _, tt := range tests {
			tree := bt.NewBinarySearch[int]()
			t.Log(reflect.TypeOf(tree).String())
			tree.Push(tt.input...)
			tree.Print()
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
		for _, tt := range tests {
			tree := bt.NewAVL[int]()
			t.Log(reflect.TypeOf(tree).String())
			tree.Push(tt.input...)
			tree.Print()
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
		for _, tt := range tests {
			tree := bt.NewRB[int]()
			t.Log(reflect.TypeOf(tree).String() == "*tree.RB[int]")
			tree.Push(tt.input...)
			tree.Print()
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
