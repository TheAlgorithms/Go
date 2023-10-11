package matrix

import (
	"errors"
	"fmt"
	"sync"
)

// Subtract subtracts the elements of the matrix (m1) from another matrix (m2) element-wise and returns a new matrix.
// The two matrices must have the same dimensions.
// If the matrices have different dimensions or if an error occurs during subtraction, an error is returned.
func (m1 Matrix[T]) Subtract(m2 Matrix[T]) (Matrix[T], error) {
	// Check if the matrices have the same dimensions.
	if !m1.MatchDimensions(m2) {
		return Matrix[T]{}, errors.New("matrices are not compatible for addition")
	}

	var zeroVal T // zero value for default values
	// Create a new matrix to store the result.
	result := New(m1.rows, m1.columns, zeroVal)
	var wg sync.WaitGroup
	for i := 0; i < m1.rows; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			for j := 0; j < m1.columns; j++ {
				sum := m1.elements[i][j] - m2.elements[i][j]
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
