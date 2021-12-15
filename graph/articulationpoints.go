// Algorithm to identify articulation points in a connected graph
// You can read more about articulation points here: https://en.wikipedia.org/wiki/Biconnected_component
// and the implementation and explanation here: https://cptalks.quora.com/Cut-Vertex-Articulation-point

package graph

import "github.com/TheAlgorithms/Go/math/min"

var (
	time               int    // time variable to keep track of the time of discovery_time of a vertex
	is_ap              []bool // boolean array to mark articulation points
	visited            []bool // boolean array to mark visited vertices
	child_cnt          []int  // array to store the number of children of a vertex
	discovery_time     []int  // array to store the time of discovery of a vertex
	earliest_discovery []int  // array to store the earliest discovered vertex reachable from a vertex
)

func ArticulationPoint(graph Graph) []bool {
	//initialize all the variables
	time = 0
	is_ap = make([]bool, graph.vertices)
	visited = make([]bool, graph.vertices)
	child_cnt = make([]int, graph.vertices)
	discovery_time = make([]int, graph.vertices)
	earliest_discovery = make([]int, graph.vertices)

	articulationPointHelper(graph, 0, -1)

	if child_cnt[0] == 1 {
		is_ap[0] = false // if the root has only one child, it is not an articulation point
	}

	return is_ap
}

func articulationPointHelper(graph Graph, u int, parent int) {
	// A recursive function to identify articulation points in a graph
	// based on the depth first search transversal of the graph, however modified to keep track
	// and update the child_cnt, discovery_time and earliest_discovery arrays defined above

	visited[u] = true

	discovery_time[u] = time                  // Mark the time of discovery of a vertex
	earliest_discovery[u] = discovery_time[u] // set the earliest discovery time to the discovered time
	time++                                    // increment the time

	for v := range graph.edges[u] {
		if v == parent {
			continue
		}
		if !visited[v] {
			child_cnt[u]++
			ArticulationPointHelper(graph, v, u)
			earliest_discovery[u] = min.Int(earliest_discovery[u], earliest_discovery[v])
			if earliest_discovery[v] >= discovery_time[u] {
				is_ap[u] = true
			}
		} else {
			earliest_discovery[u] = min.Int(earliest_discovery[u], discovery_time[v])
		}
	}
}
