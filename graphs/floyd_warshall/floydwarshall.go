// Floyd-Warshall algorithm
// https://en.wikipedia.org/wiki/Floyd%E2%80%93Warshall_algorithm

package floyd_warshall

// Matrix Defining matrix to use 2d array easier
type Matrix [][]float64

// Defining maximum value. If two vertices share this value, it means they are not connected
// var maxValue = math.Inf(1) // This is not being used in the code??

// FloydWarshall Returns all pair's shortest path using Floyd Warshall algorithm
func FloydWarshall(graph Matrix) Matrix {
	// If graph is empty or width != height, returns nil
	if len(graph) == 0 || len(graph) != len(graph[0]) {
		return nil
	}

	numVertecies := len(graph)

	// Initializing result matrix and filling it up with same values as given graph
	result := make(Matrix, numVertecies)

	for i := 0; i < numVertecies; i++ {
		result[i] = make([]float64, numVertecies)
		for j := 0; j < numVertecies; j++ {
			result[i][j] = graph[i][j]
		}
	}

	// Running over the result matrix and following the algorithm
	for k := 0; k < numVertecies; k++ {
		for i := 0; i < numVertecies; i++ {
			for j := 0; j < numVertecies; j++ {
				// If there is a less costly path from i to j node, remembering it
				if result[i][j] > result[i][k]+result[k][j] {
					result[i][j] = result[i][k] + result[k][j]
				}
			}
		}
	}

	return result
}
