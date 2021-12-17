// Algorithm to identify articulation points in a connected graph
// You can read more about articulation points here: https://en.wikipedia.org/wiki/Biconnected_component
// and the implementation and explanation here: https://cptalks.quora.com/Cut-Vertex-Articulation-point

package graph

import "github.com/TheAlgorithms/Go/math/min"

func ArticulationPoint(graph Graph) []bool {
	//initialize all the variables
	time := 0                                         // time variable to keep track of the time of discovery_time of a vertex
	is_ap := make([]bool, graph.vertices)             // boolean array to mark articulation points
	visited := make([]bool, graph.vertices)           // boolean array to mark visited vertices
	child_cnt := make([]int, graph.vertices)          // array to store the number of children of a vertex
	discovery_time := make([]int, graph.vertices)     // array to store the time of discovery of a vertex
	earliest_discovery := make([]int, graph.vertices) // array to store the earliest discovered vertex reachable from a vertex

	articulationPointHelper(graph, 0, -1, &time, is_ap, visited, child_cnt, discovery_time, earliest_discovery)

	if child_cnt[0] == 1 {
		is_ap[0] = false // if the root has only one child, it is not an articulation point
	}

	return is_ap
}

func articulationPointHelper(graph Graph, vertex int, parent int, time *int, is_ap []bool, visited []bool, child_cnt []int, discovery_time []int, earliest_discovery []int) {
	// A recursive function to identify articulation points in a graph
	// based on the depth first search transversal of the graph, however modified to keep track
	// and update the child_cnt, discovery_time and earliest_discovery arrays defined above

	visited[vertex] = true

	discovery_time[vertex] = *time                      // Mark the time of discovery of a vertex
	earliest_discovery[vertex] = discovery_time[vertex] // set the earliest discovery time to the discovered time
	*time++                                             // increment the time

	for next_vertex := range graph.edges[vertex] {
		if next_vertex == parent {
			continue
		}
		if !visited[next_vertex] {
			child_cnt[vertex]++
			articulationPointHelper(graph, next_vertex, vertex, time, is_ap, visited, child_cnt, discovery_time, earliest_discovery)
			earliest_discovery[vertex] = min.Int(earliest_discovery[vertex], earliest_discovery[next_vertex])
			if earliest_discovery[next_vertex] >= discovery_time[vertex] {
				is_ap[vertex] = true
			}
		} else {
			earliest_discovery[vertex] = min.Int(earliest_discovery[vertex], discovery_time[next_vertex])
		}
	}
}
