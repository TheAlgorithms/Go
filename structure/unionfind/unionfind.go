package unionfind

// https://en.wikipedia.org/wiki/Disjoint-set_data_structure

type Value interface{}

type node struct {
	rank  int
	value Value
}

type Set struct {
	nodes map[Value]*node
}

func New() *Set {
	return &Set{
		nodes: make(map[Value]*node),
	}
}

func (s *Set) Find(v Value) Value {
	return s.find(v, false).value
}

func (s *Set) find(value Value, init bool) *node {
	p, ok := s.nodes[value]
	if !ok {
		p = &node{value: value}
		if init {
			s.nodes[value] = p
		}
	}
	if p.value != value {
		s.nodes[value] = s.find(p.value, init)
	}
	return p
}

func (s *Set) Connected(x, y Value) bool {
	return s.Find(x) == s.Find(y)
}

func (s *Set) Union(x Value, ys ...Value) {
	for _, y := range ys {
		s.union(x, y)
	}
}

func (s *Set) union(x, y Value) {
	xRoot, yRoot := s.find(x, true), s.find(y, true)
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
