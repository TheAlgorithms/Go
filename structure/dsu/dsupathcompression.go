//Disjoint Set Union Data Structure with path compression. putting the root of the set with larger rank as a parent for the root of the set with smaller rank.
//Build:           O(n)
//FindRoot:        amortized O(α(n))
//Merge:           amortized O(α(n))
//AreInTheSameSet: amortized O(α(n))
package dsu

type dsuPathCompression struct {
	n            int
	parent, rank []int
}

//finds the root of the set which v belongs to
func (d *dsuPathCompression) FindRoot(v int) int {
	if d.parent[v] == -1 {
		return v
	}
	d.parent[v] = d.FindRoot(d.parent[v])
	return d.parent[v]
}

//if v and u belong to the same set returns false, otherwise merges these sets and returns true
func (d *dsuPathCompression) Merge(v, u int) bool {
	v, u = d.FindRoot(v), d.FindRoot(u)
	if v == u {
		return false
	}

	if d.rank[v] > d.rank[u] {
		d.parent[u] = v
	} else if d.rank[v] < d.rank[u] {
		d.parent[v] = u
	} else {
		d.parent[u] = v
		d.rank[v]++
	}
	return true
}

func (d *dsuPathCompression) AreInTheSameSet(v, u int) bool {
	return d.FindRoot(v) == d.FindRoot(u)
}

func NewDsuPathCompression(n int) *dsuPathCompression {
	d := &dsuPathCompression{
		n:      n,
		parent: make([]int, n),
		rank:   make([]int, n),
	}
	for i := 0; i < n; i++ {
		d.rank[i] = 1
		d.parent[i] = -1
	}
	return d
}
