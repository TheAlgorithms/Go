package matrix

import (
	"context"
	"errors"

	"golang.org/x/sync/errgroup"
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

	var eg errgroup.Group

	for i := 0; i < m1.rows; i++ {
		i := i // Capture the loop variable for the goroutine
		eg.Go(func() error {
			for j := 0; j < m1.columns; j++ {
				select {
				case <-ctx.Done():
					return nil // Context canceled; return without an error
				default:
				}

				diff := m1.elements[i][j] - m2.elements[i][j]
				err := result.Set(i, j, diff)
				if err != nil {
					cancel() // Cancel the context on error
					return err
				}
			}
			return nil
		})
	}

	if err := eg.Wait(); err != nil {
		return Matrix[T]{}, err
	}

	return result, nil
}
