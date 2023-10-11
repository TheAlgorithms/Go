package matrix

import (
	"errors"
	"fmt"
	"sync"
)

// Add adds the elements of the current matrix (m1) to another matrix (m2) element-wise and returns a new matrix.
// The two matrices must have the same dimensions.
// If the matrices have different dimensions or if an error occurs during addition, an error is returned.
func (m1 Matrix[T]) Add(m2 Matrix[T]) (Matrix[T], error) {
	// Check if the matrices have the same dimensions.
	if !m1.MatchDimensions(m2) {
		return Matrix[T]{}, errors.New("matrices are not compatible for addition")
	}

	// Create a new matrix to store the result.
	var zeroVal T
	result := New(m1.Rows(), m1.Columns(), zeroVal)
	var wg sync.WaitGroup
	for i := 0; i < m1.rows; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			for j := 0; j < m1.columns; j++ {
				sum := m1.elements[i][j] + m2.elements[i][j]
				err := result.Set(i, j, sum)
				if err != nil {
					fmt.Println("Error:", err)
					return
				}
			}
		}(i)
	}

	wg.Wait()
	return result, nil
}
