package tree_test

import (
	"math/rand"
	"sort"
	"testing"
	"time"

	bt "github.com/TheAlgorithms/Go/structure/tree"
)

func TestRBTreeInsert(t *testing.T) {
	tree := bt.NewRB[int]()

	ret := tree.InOrder()

	if !sort.IntsAreSorted(ret) || len(ret) != 0 {
		t.Errorf("Error with Push: %v", ret)
	}

	if r, ok := tree.Min(); ok {
		t.Errorf("Error with Min: %v", r)
	}

	if r, ok := tree.Max(); ok {
		t.Errorf("Error with Max: %v", r)
	}

	nums := []int{10, 8, 88, 888, 4, 1<<63 - 1, -(1 << 62), 188, -188, 4, 88, 1 << 32}

	tree.Push(nums...)

	ret = tree.InOrder()

	if !sort.IntsAreSorted(ret) || len(ret) != len(nums) {
		t.Errorf("Error with Push: %v", ret)
	}

	if r, ok := tree.Min(); !ok || ret[0] != r {
		t.Errorf("Error with Min: %v", r)
	}

	if r, ok := tree.Max(); !ok || ret[len(ret)-1] != r {
		t.Errorf("Error with Max: %v", r)
	}
}

func TestRBTreeDelete(t *testing.T) {
	tree := bt.NewRB[int]()
	var ok bool

	nums := []int{10, 8, 88, 888, 4, 1<<63 - 1, -(1 << 62), 188, -188, 4, 88, 1 << 32}
	tree.Push(nums...)

	ok = tree.Delete(188)

	if ret := tree.InOrder(); !ok || !sort.IntsAreSorted(ret) || len(ret) != len(nums)-1 {
		t.Errorf("Error with Delete: %v", ret)
	}

	ok = tree.Delete(188)
	if ret := tree.InOrder(); ok || !sort.IntsAreSorted(ret) || len(ret) != len(nums)-1 {
		t.Errorf("Error with Delete: %v", ret)
	}

	ok = tree.Delete(1<<63 - 1)
	if ret := tree.InOrder(); !ok || !sort.IntsAreSorted(ret) || len(ret) != len(nums)-2 {
		t.Errorf("Error with Delete: %v", ret)
	}

	ok = tree.Delete(4)
	if ret := tree.InOrder(); !ok || !sort.IntsAreSorted(ret) || len(ret) != len(nums)-3 {
		t.Errorf("Error with Delete: %v", ret)
	}

	ok = tree.Delete(4)
	if ret := tree.InOrder(); !ok || !sort.IntsAreSorted(ret) || len(ret) != len(nums)-4 {
		t.Errorf("Error with Delete: %v", ret)
	}

	if ret, ok := tree.Max(); !ok || ret != (1<<32) {
		t.Errorf("Error with Delete, max: %v, want: %v", ret, (1 << 32))
	}

	if ret, ok := tree.Min(); !ok || ret != -(1<<62) {
		t.Errorf("Error with Delete, min: %v, want: %v", ret, (1 << 32))
	}
}

func FuzzRBTree(f *testing.F) {
	testcases := []int{100, 200, 1000, 10000}
	for _, tc := range testcases {
		f.Add(tc)
	}

	f.Fuzz(func(t *testing.T, a int) {
		rand.Seed(time.Now().Unix())
		tree := bt.NewRB[int]()
		nums := rand.Perm(a)
		tree.Push(nums...)

		rets := tree.InOrder()
		if !sort.IntsAreSorted(rets) {
			t.Error("Error with Push")
		}

		if res, ok := tree.Min(); !ok || res != rets[0] {
			t.Errorf("Error with Min, get %d, want: %d", res, rets[0])
		}

		if res, ok := tree.Max(); !ok || res != rets[a-1] {
			t.Errorf("Error with Max, get %d, want: %d", res, rets[a-1])
		}

		for i := 0; i < a-1; i++ {
			ok := tree.Delete(nums[i])
			rets = tree.InOrder()
			if !ok || !sort.IntsAreSorted(rets) {
				t.Errorf("Error With Delete")
			}
		}
	})
}

func TestRBTreePredecessorAndSuccessor(t *testing.T) {
	tree := bt.NewRB[int]()

	nums := []int{10, 8, 88, 888, 4, -1, 100}
	tree.Push(nums...)

	if ret, ok := tree.Predecessor(100); !ok && ret == 88 {
		t.Errorf("Error with Predecessor")
	}

	if _, ok := tree.Predecessor(-1); ok {
		t.Errorf("Error with Predecessor")
	}

	if _, ok := tree.Predecessor(-12); ok {
		t.Errorf("Error with Predecessor")
	}

	if ret, ok := tree.Predecessor(4); !ok && ret == -1 {
		t.Errorf("Error with Predecessor")
	}

	if ret, ok := tree.Successor(4); !ok && ret == 8 {
		t.Errorf("Error with Successor")
	}

	if ret, ok := tree.Successor(8); !ok && ret == 88 {
		t.Errorf("Error with Successor")
	}

	if ret, ok := tree.Successor(88); !ok && ret == 100 {
		t.Errorf("Error with Successor")
	}

	if ret, ok := tree.Successor(100); !ok && ret == 888 {
		t.Errorf("Error with Successor")
	}

	if ret, ok := tree.Successor(-1); !ok && ret == 4 {
		t.Errorf("Error with Successor")
	}

	if _, ok := tree.Successor(888); ok {
		t.Errorf("Error with Successor")
	}

	if _, ok := tree.Successor(188); ok {
		t.Errorf("Error with Successor")
	}
}

func TestRBTreeString(t *testing.T) {
	tree := bt.NewRB[string]()

	if tree.Has("Golang") {
		t.Errorf("Error with Has when T is string.")
	}

	strs := []string{"Hello", "World", "Golang", "Python", "Rust", "C", "JavaScript", "Haskell", "Pascal", "ZZ"}
	for _, str := range strs {
		tree.Push(str)
	}

	if !tree.Has("Golang") {
		t.Errorf("Error with Has when T is string.")
	}

	if tree.Has("Pasc") {
		t.Errorf("Error with Has when T is string.")
	}

	if ok := tree.Delete("Hello"); !ok {
		t.Errorf("Error with Delete when T is string.")
	}

	if ok := tree.Delete("Pasc"); ok {
		t.Errorf("Error with Delete when T is string.")
	}

	if tree.Has("Hello") {
		t.Errorf("Error with Has when T is string.")
	}

	ret := tree.InOrder()
	if !sort.StringsAreSorted(ret) {
		t.Errorf("Error with Push when T is string")
	}

	if ret, ok := tree.Min(); !ok || ret != "C" {
		t.Errorf("Error with Min when T is string")
	}

	if ret, ok := tree.Max(); !ok || ret != "ZZ" {
		t.Errorf("Error with Max when T is string")
	}
}
