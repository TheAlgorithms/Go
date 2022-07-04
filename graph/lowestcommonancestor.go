package graph

type TreeEdge struct {
	from int
	to   int
}

type ITree interface {
	dfs(int, int)
	addEdge(int, int)
	BuildLCA()
	GetDepth(int) int
	GetLCA(int, int) int
}

type Tree struct {
	numbersVertex int
	root          int
	MAXLOG        int
	depth         []int
	jump          [][]int
	edges         [][]int
}

func (tree *Tree) addEdge(u, v int) {
	tree.edges[u] = append(tree.edges[u], v)
	tree.edges[v] = append(tree.edges[v], u)
}

func (tree *Tree) dfs(u, par int) {
	tree.jump[0][u] = par
	for _, v := range tree.edges[u] {
		if v != par {
			tree.depth[v] = tree.depth[u] + 1
			tree.dfs(v, u)
		}
	}
}

func (tree *Tree) BuildLCA() {
	tree.dfs(tree.root, tree.root)

	for j := 1; j < tree.MAXLOG; j++ {
		for u := 0; u < tree.numbersVertex; u++ {
			tree.jump[j][u] = tree.jump[j-1][tree.jump[j-1][u]]
		}
	}
}

func (tree *Tree) GetDepth(u int) int {
	return tree.depth[u]
}

func (tree *Tree) GetLCA(u, v int) int {
	if tree.GetDepth(u) < tree.GetDepth(v) {
		u, v = v, u
	}

	for j := tree.MAXLOG - 1; j >= 0; j-- {
		if tree.GetDepth(tree.jump[j][u]) >= tree.GetDepth(v) {
			u = tree.jump[j][u]
		}
	}

	if u == v {
		return u
	}

	for j := tree.MAXLOG - 1; j >= 0; j-- {
		if tree.jump[j][u] != tree.jump[j][v] {
			u = tree.jump[j][u]
			v = tree.jump[j][v]
		}
	}

	return tree.jump[0][u]
}

func NewTree(numbersVertex, root int, edges []TreeEdge) (tree *Tree) {
	tree = new(Tree)
	tree.numbersVertex, tree.root, tree.MAXLOG = numbersVertex, root, 0
	tree.depth = make([]int, numbersVertex)

	for (1 << tree.MAXLOG) <= numbersVertex {
		(tree.MAXLOG) += 1
	}
	(tree.MAXLOG) += 1

	tree.jump = make([][]int, tree.MAXLOG)
	for j := 0; j < tree.MAXLOG; j++ {
		tree.jump[j] = make([]int, numbersVertex)
	}

	tree.edges = make([][]int, numbersVertex)
	for _, e := range edges {
		tree.addEdge(e.from, e.to)
	}

	return tree
}
