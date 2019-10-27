// Breadth-first search implementation
package bfs

import "fmt"

type Vertex struct {
	Value    int
	Adjacent []*Vertex
}

type BFS struct {
	visited map[*Vertex]bool
	parents map[*Vertex]*Vertex
}

func (b *BFS) BFS(start *Vertex, goal *Vertex) *Vertex {
	b.visited = make(map[*Vertex]bool)
	b.visited[start] = true

	b.parents = make(map[*Vertex]*Vertex)

	queue := []*Vertex{start}
	for len(queue) > 0 {
		v := queue[0]
		fmt.Println("Visiting", v.Value)
		queue = queue[1:]
		if v.Value == goal.Value {
			fmt.Println("Found goal")
			return v
		}

		for _, adjacent := range v.Adjacent {
			if !b.visited[adjacent] {
				b.visited[adjacent] = true
				b.parents[adjacent] = v
				queue = append(queue, adjacent)
			}
		}
	}

	return nil
}

func (b *BFS) PrintPath(destination *Vertex) {
	if destination == nil {
		fmt.Println("No path")
		return
	}

	current := destination
	for current != nil {
		fmt.Printf("%d", current.Value)

		parent := b.parents[current]
		if parent == nil {
			break
		}

		fmt.Printf(" <- ")

		current = parent
	}

	fmt.Println()
}

func main() {
	a := Vertex{1, nil}
	b := Vertex{2, nil}
	c := Vertex{3, nil}
	d := Vertex{4, nil}
	e := Vertex{5, nil}
	f := Vertex{6, nil}
	g := Vertex{7, nil}
	h := Vertex{8, nil}

	a.Adjacent = append(a.Adjacent, []*Vertex{&b, &c}...)
	b.Adjacent = append(b.Adjacent, []*Vertex{&a, &d, &e}...)
	c.Adjacent = append(c.Adjacent, []*Vertex{&a, &f, &g}...)
	d.Adjacent = append(d.Adjacent, []*Vertex{&b}...)
	e.Adjacent = append(e.Adjacent, []*Vertex{&b, &h}...)
	f.Adjacent = append(f.Adjacent, []*Vertex{&c}...)
	g.Adjacent = append(g.Adjacent, []*Vertex{&c}...)
	h.Adjacent = append(h.Adjacent, []*Vertex{&e}...)

	bfs := BFS{}
	found := bfs.BFS(&a, &h)
	if found == nil {
		fmt.Println("Vertex not found")
		return
	}

	bfs.PrintPath(found) // 8 <- 5 <- 2 <- 1
}
