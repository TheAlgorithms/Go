package set

// https://en.wikipedia.org/wiki/Disjoint-set_data_structure

type Value interface{}

type node struct {
	rank  int
	value Value
}

type UFS struct {
	nodes map[Value]*node
}

func NewUFS() *UFS {
	return &UFS{
		nodes: make(map[Value]*node),
	}
}

func (u *UFS) Find(v Value) Value {
	return u.find(v, false).value
}

func (u *UFS) find(value Value, init bool) *node {
	p, ok := u.nodes[value]
	if !ok {
		p = &node{value: value}
		if init {
			u.nodes[value] = p
		}
	}
	if p.value != value {
		u.nodes[value] = u.find(p.value, init)
	}
	return p
}

func (u *UFS) Connected(x, y Value) bool {
	return u.Find(x) == u.Find(y)
}

func (u *UFS) Union(x Value, ys ...Value) {
	for _, y := range ys {
		u.union(x, y)
	}
}

func (u *UFS) union(x, y Value) {
	xRoot, yRoot := u.find(x, true), u.find(y, true)
	if xRoot.value == yRoot.value {
		return
	}
	if xRoot.rank < yRoot.rank {
		xRoot, yRoot = yRoot, xRoot
	}
	yRoot.value = xRoot.value

	if xRoot.rank == yRoot.rank {
		xRoot.rank++
	}
}
