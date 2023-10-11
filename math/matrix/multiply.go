package matrix

import (
	"context"
	"errors"

	"golang.org/x/sync/errgroup"
)

// Multiply multiplies the current matrix (m1) with another matrix (m2) and returns the result as a new matrix.
// The two matrices must have compatible dimensions for multiplication, i.e., the number of columns in the first matrix
// must match the number of rows in the second matrix.

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

	g, ctx := errgroup.WithContext(ctx)

	for i := 0; i < m1.Rows(); i++ {
		for j := 0; j < m2.Columns(); j++ {
			i, j := i, j // Capture the loop variable for the goroutine
			g.Go(func() error {
				// Compute the dot product of the row from the first matrix and the column from the second matrix.
				dotProduct := zeroVal
				for k := 0; k < m1.Columns(); k++ {
					select {
					case <-ctx.Done():
						return nil // Context canceled; return without an error
					default:
					}

					val1, err := m1.Get(i, k)
					if err != nil {
						cancel()
						return err
					}
					val2, err := m2.Get(k, j)
					if err != nil {
						cancel()
						return err
					}
					dotProduct += val1 * val2
				}
				err := result.Set(i, j, dotProduct)
				if err != nil {
					cancel()
					return err
				}
				return nil
			})
		}
	}

	if err := g.Wait(); err != nil {
		return Matrix[T]{}, err
	}

	return result, nil
}
