package set

import (
	"testing"
)

func TestConnected(t *testing.T) {
	s := NewUFS()

	group1 := []Value{1, 3, 5, 7, 9}
	group2 := []Value{"a", "b", "c", "d", "e"}

	s.Union(group1[0], group1[1:]...)
	s.Union(group2[0], group2[1:]...)

	testConnected(t, s, group1, group1, true)
	testConnected(t, s, group2, group2, true)
	testConnected(t, s, group1, group2, false)

	s.Union(1, "b")

	testConnected(t, s, group1, group2, true)
}

func testConnected(t *testing.T, s *UFS, xs, ys []Value, want bool) {
	for _, x := range xs {
		for _, y := range ys {
			if got := s.Connected(x, y); got != want {
				t.Errorf("test Conneted error: x: %v, y:%v got:%v want:%v", x, y, got, want)
			}
		}
	}
}
