// Algorithm to identify articulation points in a connected graph
// You can read more about articulation points here: https://en.wikipedia.org/wiki/Biconnected_component
// and the implementation and explanation here: https://cptalks.quora.com/Cut-Vertex-Articulation-point

package graph

import "github.com/TheAlgorithms/Go/math/min"

type apHelper struct {
	is_ap              []bool
	visited            []bool
	child_cnt          []int
	discovery_time     []int
	earliest_discovery []int
}

func ArticulationPoint(graph Graph) []bool {
	time := 0 // time variable to keep track of the time of discovery_time of a vertex

	//initialize all the variables
	apHelperInstance := &apHelper{
		is_ap:              make([]bool, graph.vertices), // boolean array to mark articulation points
		visited:            make([]bool, graph.vertices), // boolean array to mark visited vertices
		child_cnt:          make([]int, graph.vertices),  // array to store the number of children of a vertex
		discovery_time:     make([]int, graph.vertices),  // array to store the time of discovery of a vertex
		earliest_discovery: make([]int, graph.vertices),  // array to store the earliest discovered vertex reachable from a vertex
	}
	articulationPointHelper(apHelperInstance, 0, -1, &time, graph)

	if apHelperInstance.child_cnt[0] == 1 {
		apHelperInstance.is_ap[0] = false // if the root has only one child, it is not an articulation point
	}

	return apHelperInstance.is_ap
}

func articulationPointHelper(apHelperInstance *apHelper, vertex int, parent int, time *int, graph Graph) {
	// A recursive function to identify articulation points in a graph
	// based on the depth first search transversal of the graph, however modified to keep track
	// and update the child_cnt, discovery_time and earliest_discovery arrays defined above

	apHelperInstance.visited[vertex] = true

	apHelperInstance.discovery_time[vertex] = *time                                       // Mark the time of discovery of a vertex
	apHelperInstance.earliest_discovery[vertex] = apHelperInstance.discovery_time[vertex] // set the earliest discovery time to the discovered time
	*time++                                                                               // increment the time

	for next_vertex := range graph.edges[vertex] {
		if next_vertex == parent {
			continue
		}

		if apHelperInstance.visited[next_vertex] { // if the vertex is already visited, then it is a back edge and manipulate earliest discovery time accordingly
			apHelperInstance.earliest_discovery[vertex] = min.Int(apHelperInstance.earliest_discovery[vertex], apHelperInstance.discovery_time[next_vertex])
		} else {
			apHelperInstance.child_cnt[vertex]++
			articulationPointHelper(apHelperInstance, next_vertex, vertex, time, graph)
			apHelperInstance.earliest_discovery[vertex] = min.Int(apHelperInstance.earliest_discovery[vertex], apHelperInstance.earliest_discovery[next_vertex])
			if apHelperInstance.earliest_discovery[next_vertex] >= apHelperInstance.discovery_time[vertex] {
				apHelperInstance.is_ap[vertex] = true
			}
		}
	}
}
