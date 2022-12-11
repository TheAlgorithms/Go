package heap_test

import (
	"github.com/TheAlgorithms/Go/structure/heap"
	"reflect"
	"testing"
)

type testInt int

func (u testInt) Less(o testInt) bool {
	return u < o
}

type testStudent struct {
	Name  string
	Score int64
}

func (u testStudent) Less(o testStudent) bool {
	if u.Score == o.Score {
		return u.Name < o.Name
	}
	return u.Score > o.Score
}

func TestHeap_Empty(t *testing.T) {
	tests := []struct {
		name string
		want bool
	}{
		{name: "empty", want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := heap.New[testInt]()
			if got := h.Empty(); got != tt.want {
				t.Errorf("Empty() = %v, want %v", got, tt.want)
			}
		})
	}
}

type testOpType int

const (
	testPush  = 1
	testPop   = 2
	testTop   = 3
	testEmpty = 4
)

type testOp[T any] struct {
	typ     testOpType
	x       T
	isEmpty bool
}

type testStruct[T any] struct {
	name string
	ops  []testOp[T]
}

func TestHeapExample1(t *testing.T) {
	tests1 := []testStruct[testInt]{
		{
			name: "example 1",
			ops: []testOp[testInt]{
				{typ: testEmpty, isEmpty: true},
				{typ: testPush, x: 10},
				{typ: testEmpty, isEmpty: false},
				{typ: testTop, x: 10},
				{typ: testPop},
				{typ: testEmpty, isEmpty: true},
				{typ: testPush, x: 9},
				{typ: testPush, x: 8},
				{typ: testPop},
				{typ: testPush, x: 3},
				{typ: testTop, x: 3},
				{typ: testPush, x: 2},
				{typ: testTop, x: 2},
				{typ: testPush, x: 4},
				{typ: testPush, x: 6},
				{typ: testPush, x: 5},
				{typ: testTop, x: 2},
				{typ: testPop},
				{typ: testTop, x: 3},
				{typ: testPop},
				{typ: testPop},
				{typ: testTop, x: 5},
				{typ: testEmpty, isEmpty: false},
			},
		},
	}
	testFunc(t, tests1, testInt.Less)
}

func TestHeapExample2(t *testing.T) {
	tests1 := []testStruct[testStudent]{
		{
			name: "example 2",
			ops: []testOp[testStudent]{
				{typ: testPush, x: testStudent{Name: "Alan", Score: 87}},
				{typ: testPush, x: testStudent{Name: "Bob", Score: 98}},
				{typ: testTop, x: testStudent{Name: "Bob", Score: 98}},
				{typ: testPop},
				{typ: testPush, x: testStudent{Name: "Carl", Score: 70}},
				{typ: testTop, x: testStudent{Name: "Alan", Score: 87}},
			},
		},
	}
	testFunc(t, tests1, testStudent.Less)
}

func testFunc[T any](t *testing.T, tests []testStruct[T], less func(a, b T) bool) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h, err := heap.NewAny[T](less)
			if err != nil {
				t.Errorf("New Heap err %v", err)
			}
			for i, op := range tt.ops {
				switch op.typ {
				case testPush:
					oldSize := h.Size()
					h.Push(op.x)
					newSize := h.Size()
					if oldSize+1 != newSize {
						t.Errorf("op %d testPush %v failed", i, op.x)
					}
				case testPop:
					oldSize := h.Size()
					h.Pop()
					newSize := h.Size()
					if oldSize-1 != newSize {
						t.Errorf("op %d testPop %v failed", i, op.x)
					}
				case testTop:
					if got := h.Top(); !reflect.DeepEqual(got, op.x) {
						t.Errorf("op %d testTop %v, want %v", i, got, op.x)
					}
				case testEmpty:
					if got := h.Empty(); got != op.isEmpty {
						t.Errorf("op %d Empty() = %v, want %v", i, got, op.isEmpty)
					}
				}
			}
		})
	}
}
