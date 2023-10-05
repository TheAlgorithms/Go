package matrix

import (
	"errors"
	"fmt"
)

// Submatrix extracts a submatrix from the current matrix.
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

	for i := 0; i < numRows; i++ {
		for j := 0; j < numCols; j++ {
			val, err := m.Get(rowStart+i, colStart+j)
			if err != nil {
				return Matrix[T]{}, err
			}
			err1 := subMatrix.Set(i, j, val)
			if err1 != nil {
				return Matrix[T]{}, fmt.Errorf("invalid matrices: %v", err1)
			}
		}
	}

	return subMatrix, nil
}
