package graph

import "fmt"

func addto(adj map[int][]int, a int, b int) {
    _, isin := adj[a]
    if !isin {
        adj[a] = []int{}
    }
    adj[a] = append(adj[a], b)
}

func IsBipartite(N int, edges [][]int) bool {
    adj := map[int][]int{}
    for _,e := range edges {
        addto(adj, e[0],e[1])
        addto(adj, e[1],e[0])
    }

    stack := []int{0}
    visited := map[int]bool{}
    for len(stack)>0 {
    }
    fmt.Println(adj)

    // 0 is uninit, 1/2 is colors
    colors := make([]int, N)
    return len(colors)>N
}
