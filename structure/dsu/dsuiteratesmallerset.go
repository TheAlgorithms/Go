//Disjoint Set Union Data Structure. in this implementation, each set is a tree with a height at most 1
//when merging two trees we assign the root of the larger tree as a new parent for each node in the smaller tree by iterating on the smaller set.
//Build:           O(n)
//FindRoot:        O(1)
//Merge:           amortized O(log(n))
//AreInTheSameSet: O(1)
package dsu

type dsuIterateSmallerSet struct {
	n      int
	parent []int
	childs [][]int
}

//finds the root of the set which v belongs to
func (d *dsuIterateSmallerSet) FindRoot(v int) int {
	if d.parent[v] == -1 {
		return v
	}
	return d.parent[v]
}

//during merging two sets we use this method to iterate the smaller set and assign the new root as a parent for each node
func (d *dsuIterateSmallerSet) setNewParent(v, newParent int) {
	d.parent[v] = newParent
	for _, u := range d.childs[v] {
		d.parent[u] = newParent
	}
}

//if v and u belong to the same set returns false, otherwise merges these sets and returns true
func (d *dsuIterateSmallerSet) Merge(v, u int) bool {
	v, u = d.FindRoot(v), d.FindRoot(u)
	if v == u {
		return false
	}
	if len(d.childs[v]) < len(d.childs[u]) {
		v, u = u, v
	}
	d.setNewParent(u, v)
	d.childs[v] = append(d.childs[v], u)
	d.childs[v] = append(d.childs[v], d.childs[u]...)
	d.childs[u] = make([]int, 0)
	return true
}

func (d *dsuIterateSmallerSet) AreInTheSameSet(v, u int) bool {
	return d.FindRoot(v) == d.FindRoot(u)
}

func NewDsuIterateSmallerSet(n int) *dsuIterateSmallerSet {
	d := &dsuIterateSmallerSet{
		n:      n,
		parent: make([]int, n),
		childs: make([][]int, n),
	}
	for i := 0; i < n; i++ {
		d.childs[i] = make([]int, 0)
		d.parent[i] = -1
	}
	return d
}
