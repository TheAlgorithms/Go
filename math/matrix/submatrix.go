package matrix

import (
	"context"
	"errors"

	"golang.org/x/sync/errgroup"
)

// SubMatrix extracts a submatrix from the current matrix.
func (m Matrix[T]) SubMatrix(rowStart, colStart, numRows, numCols int) (Matrix[T], error) {
	if rowStart < 0 || colStart < 0 || numRows < 0 || numCols < 0 {
		return Matrix[T]{}, errors.New("negative dimensions are not allowed")
	}

	if rowStart+numRows > m.rows || colStart+numCols > m.columns {
		return Matrix[T]{}, errors.New("submatrix dimensions exceed matrix bounds")
	}

	var zeroVal T
	if numRows == 0 || numCols == 0 {
		return New(numRows, numCols, zeroVal), nil // Return an empty matrix
	}

	subMatrix := New(numRows, numCols, zeroVal)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // Make sure it's called to release resources even if no errors

	var eg errgroup.Group
	var firstErr error

	for i := 0; i < numRows; i++ {
		i := i // Capture the loop variable for the goroutine
		eg.Go(func() error {
			for j := 0; j < numCols; j++ {
				select {
				case <-ctx.Done():
					return nil // Context canceled; return without an error
				default:
				}

				val, err := m.Get(rowStart+i, colStart+j)
				if err != nil {
					cancel()
					return err
				}

				err = subMatrix.Set(i, j, val)
				if err != nil {
					cancel()
					return err
				}
			}
			return nil
		})
	}

	if err := eg.Wait(); err != nil {
		return Matrix[T]{}, err
	}

	return subMatrix, firstErr
}
