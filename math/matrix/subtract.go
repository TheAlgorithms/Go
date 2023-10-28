package matrix

import (
	"context"
	"errors"
	"sync"
)

// Subtract subtracts two matrices.
func (m1 Matrix[T]) Subtract(m2 Matrix[T]) (Matrix[T], error) {
	// Check if the matrices have the same dimensions.
	if !m1.MatchDimensions(m2) {
		return Matrix[T]{}, errors.New("matrices are not compatible for subtraction")
	}

	// Create a new matrix to store the result.
	var zeroVal T
	result := New(m1.Rows(), m1.Columns(), zeroVal)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // Make sure it's called to release resources even if no errors

	var wg sync.WaitGroup
	errCh := make(chan error, 1)

	for i := 0; i < m1.rows; i++ {
		i := i // Capture the loop variable for the goroutine
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < m1.columns; j++ {
				select {
				case <-ctx.Done():
					return // Context canceled; return without an error
				default:
				}

				diff := m1.elements[i][j] - m2.elements[i][j]
				err := result.Set(i, j, diff)
				if err != nil {
					cancel() // Cancel the context on error
					select {
					case errCh <- err:
					default:
					}
					return
				}
			}
		}()
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
