// multiply.go
// description: Implementation of matrix multiplication
// time complexity: O(n^3) where n is the number of rows in the first matrix
// space complexity: O(n^2) where n is the number of rows in the first matrix

package matrix

import (
	"context"
	"errors"
	"sync"
)

// Multiply multiplies the current matrix (m1) with another matrix (m2) and returns the result as a new matrix.
func (m1 Matrix[T]) Multiply(m2 Matrix[T]) (Matrix[T], error) {
	// Check if the matrices can be multiplied.
	if m1.Columns() != m2.Rows() {
		return Matrix[T]{}, errors.New("matrices cannot be multiplied: column count of the first matrix must match row count of the second matrix")
	}

	// Create a new matrix to store the result.
	var zeroVal T
	result := New(m1.Rows(), m2.Columns(), zeroVal)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // Make sure it's called to release resources even if no errors

	var wg sync.WaitGroup
	errCh := make(chan error, 1)

	for i := 0; i < m1.Rows(); i++ {
		for j := 0; j < m2.Columns(); j++ {
			i, j := i, j // Capture the loop variable for the goroutine
			wg.Add(1)
			go func() {
				defer wg.Done()
				// Compute the dot product of the row from the first matrix and the column from the second matrix.
				dotProduct := zeroVal
				for k := 0; k < m1.Columns(); k++ {
					select {
					case <-ctx.Done():
						return // Context canceled; return without an error
					default:
					}

					val1, err := m1.Get(i, k)
					if err != nil {
						cancel()
						select {
						case errCh <- err:
						default:
						}
						return
					}
					val2, err := m2.Get(k, j)
					if err != nil {
						cancel()
						select {
						case errCh <- err:
						default:
						}
						return
					}
					dotProduct += val1 * val2
				}
				err := result.Set(i, j, dotProduct)
				if err != nil {
					cancel()
					select {
					case errCh <- err:
					default:
					}
					return
				}
			}()
		}
	}

	// Wait for all goroutines to finish
	go func() {
		wg.Wait()
		close(errCh)
	}()

	// Check for any errors
	if err := <-errCh; err != nil {
		return Matrix[T]{}, err
	}

	return result, nil
}
