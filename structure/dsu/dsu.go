//Disjoint Set Union Data Structure. putting the root of the larger set as a parent for the root of the smaller set.
//Build: O(n)
//FindRoot: O(log(n))
//Merge: O(log(n))
//AreInTheSameSet: O(log(n))
package dsu

type dsu struct {
	n            int
	parent, size []int
}

type DSU interface {
	FindRoot(int) int
	Merge(int, int) bool
	AreInTheSameSet(int, int) bool
}

//finds the root of the set which v belongs to
func (d *dsu) FindRoot(v int) int {
	if d.parent[v] == -1 {
		return v
	}
	return d.FindRoot(d.parent[v])
}

//if v and u belong to the same set returns false, otherwise merges these sets and returns true
func (d *dsu) Merge(v, u int) bool {
	v, u = d.FindRoot(v), d.FindRoot(u)
	if v == u {
		return false
	}
	if d.size[v] < d.size[u] {
		v, u = u, v
	}
	d.size[v] += d.size[u]
	d.parent[u] = v
	return true
}

func (d *dsu) AreInTheSameSet(v, u int) bool {
	return d.FindRoot(v) == d.FindRoot(u)
}

func NewDsu(n int) *dsu {
	d := &dsu{
		n:      n,
		parent: make([]int, n),
		size:   make([]int, n),
	}
	for i := 0; i < n; i++ {
		d.size[i] = 1
		d.parent[i] = -1
	}
	return d
}
