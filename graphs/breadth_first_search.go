package graph

import (
	"strconv"
	"testing"
)

// BFS traverses g in breadth-first order starting at v.
// When the algorithm follows an edge (v, w) and finds a previously
// unvisited vertex w, it calls do(v, w, c) with c equal to
// the cost of the edge (v, w).
func BFS(g Iterator, v int, do func(v, w int, c int64)) {
	visited := make([]bool, g.Order())
	visited[v] = true
	for queue := []int{v}; len(queue) > 0; {
		v := queue[0]
		queue = queue[1:]
		g.Visit(v, func(w int, c int64) (skip bool) {
			if visited[w] {
				return
			}
			do(v, w, c)
			visited[w] = true
			queue = append(queue, w)
			return
		})
	}
}


//Driver Function
func TestBFS(t *testing.T) {
	g := New(10)
	for _, e := range []struct {
		v, w int
	}{
		{0, 1}, {0, 4}, {0, 7}, {0, 9},
		{4, 2}, {7, 5}, {7, 8},
		{2, 3}, {5, 6},
		{3, 6}, {8, 9}, {4, 4},
	} {
		g.AddBoth(e.v, e.w)
	}
	exp := "0147925836"
	res := "0"
	BFS(Sort(g), 0, func(v, w int, c int64) {
		res += strconv.Itoa(w)
	})
	if mess, diff := diff(res, exp); diff {
		t.Errorf("BFS: %s", mess)
	}
}
