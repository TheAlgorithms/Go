package tree_test

import (
	"math/rand"
	"sort"
	"testing"
	"time"

	bt "github.com/TheAlgorithms/Go/structure/tree"
)

func TestRBTreePush(t *testing.T) {
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

	nums := []int{10, 8, 88, 888, 4, 1<<63 - 1, -(1 << 62), 188, -188, 4, 1 << 32}

	tree.Push(nums...)

	ret = tree.InOrder()

	if !sort.IntsAreSorted(ret) {
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

	if ret := tree.InOrder(); !ok || !sort.IntsAreSorted(ret) {
		t.Errorf("Error with Delete: %v", ret)
	}

	ok = tree.Delete(188)
	if ret := tree.InOrder(); ok || !sort.IntsAreSorted(ret) {
		t.Errorf("Error with Delete: %v", ret)
	}

	ok = tree.Delete(1<<63 - 1)
	if ret := tree.InOrder(); !ok || !sort.IntsAreSorted(ret) {
		t.Errorf("Error with Delete: %v", ret)
	}

	ok = tree.Delete(4)
	if ret := tree.InOrder(); !ok || !sort.IntsAreSorted(ret) {
		t.Errorf("Error with Delete: %v", ret)
	}

	if ret, ok := tree.Max(); !ok || ret != (1<<32) {
		t.Errorf("Error with Delete, max: %v, want: %v", ret, (1 << 32))
	}

	if ret, ok := tree.Min(); !ok || ret != -(1<<62) {
		t.Errorf("Error with Delete, min: %v, want: %v", ret, (1 << 32))
	}
}

func TestRBTree(t *testing.T) {
	testcases := []int{100, 200, 1000, 10000}
	for _, n := range testcases {
		rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
		tree := bt.NewRB[int]()
		nums := rnd.Perm(n)
		tree.Push(nums...)

		rets := tree.InOrder()
		if !sort.IntsAreSorted(rets) {
			t.Error("Error with Push")
		}

		if res, ok := tree.Min(); !ok || res != rets[0] {
			t.Errorf("Error with Min, get %d, want: %d", res, rets[0])
		}

		if res, ok := tree.Max(); !ok || res != rets[n-1] {
			t.Errorf("Error with Max, get %d, want: %d", res, rets[n-1])
		}

		for i := 0; i < n-1; i++ {
			if ret, ok := tree.Successor(rets[0]); ret != rets[1] || !ok {
				t.Error("Error with Successor")
			}
			if ret, ok := tree.Predecessor(rets[1]); ret != rets[0] || !ok {
				t.Error("Error with Predecessor")
			}

			ok := tree.Delete(nums[i])
			rets = tree.InOrder()
			if !ok || !sort.IntsAreSorted(rets) {
				t.Errorf("Error With Delete")
			}
		}
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
