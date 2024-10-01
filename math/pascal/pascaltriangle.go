// pascaltriangle.go
// description: Pascal's triangle
// details:
// Pascal's triangle is a triangular array of the binomial coefficients that arises in probability theory, combinatorics, and algebra. - [Pascal's triangle](https://en.wikipedia.org/wiki/Pascal%27s_triangle)
// example:
//1
//1 1
//1 2 1
//1 3 3 1
//1 4 6 4 1
//1 5 10 10 5 1
//1 6 15 20 15 6 1
//1 7 21 35 35 21 7 1
//1 8 28 56 70 56 28 8 1
//1 9 36 84 126 126 84 36 9 1
//1 10 45 120 210 252 210 120 45 10 1
//...
// author(s) [red_byte](https://github.com/i-redbyte)
// time complexity: O(n^2)
// space complexity: O(n^2)
// see pascaltriangle_test.go

package pascal

// GenerateTriangle This function generates a Pascal's triangle of n lines
func GenerateTriangle(n int) [][]int {
	var triangle = make([][]int, n)
	for i := 0; i < n; i++ {
		triangle[i] = make([]int, i+1)
		triangle[i][0], triangle[i][i] = 1, 1
		for j := 1; j < i; j++ {
			triangle[i][j] = triangle[i-1][j] + triangle[i-1][j-1]
		}
	}
	return triangle
}
