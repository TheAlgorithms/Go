package matrix

func MatrixEqual[T comparable](mat1, mat2 [][]T) bool {
	if len(mat1) != len(mat2) || len(mat1[0]) != len(mat2[0]) {
		return false
	}
	for i := range mat1 {
		for j := range mat1[i] {
			if mat1[i][j] != mat2[i][j] {
				return false
			}
		}
	}
	return true
}
