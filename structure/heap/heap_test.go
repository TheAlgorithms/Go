package heap

import (
	"testing"
)

type _int int

func (u _int) Less(o _int) bool {
	return u < o
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
			h := New[_int]()
			if got := h.Empty(); got != tt.want {
				t.Errorf("Empty() = %v, want %v", got, tt.want)
			}
		})
	}
}

type _opType int

const (
	_push  = 1
	_pop   = 2
	_test  = 3
	_empty = 4
)

type _op struct {
	typ     _opType
	x       _int
	isEmpty bool
}

func TestHeap(t *testing.T) {
	tests := []struct {
		name string
		ops  []_op
	}{
		{
			name: "example",
			ops: []_op{
				{typ: _empty, isEmpty: true},
				{typ: _push, x: 10},
				{typ: _empty, isEmpty: false},
				{typ: _test, x: 10},
				{typ: _pop},
				{typ: _empty, isEmpty: true},
				{typ: _push, x: 9},
				{typ: _push, x: 8},
				{typ: _pop},
				{typ: _push, x: 3},
				{typ: _test, x: 3},
				{typ: _push, x: 2},
				{typ: _test, x: 2},
				{typ: _push, x: 4},
				{typ: _push, x: 6},
				{typ: _push, x: 5},
				{typ: _test, x: 2},
				{typ: _pop},
				{typ: _test, x: 3},
				{typ: _pop},
				{typ: _pop},
				{typ: _test, x: 5},
				{typ: _empty, isEmpty: false},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := New[_int]()
			for i, op := range tt.ops {
				switch op.typ {
				case _push:
					h.Push(op.x)
				case _pop:
					h.Pop()
				case _test:
					if got := h.Top(); got != op.x {
						t.Errorf("op %d Test() = %v, want %v", i, got, op.x)
					}
				case _empty:
					if got := h.Empty(); got != op.isEmpty {
						t.Errorf("op %d Empty() = %v, want %v", i, got, op.isEmpty)
					}
				}
			}
		})
	}
}
