package backtracking

import "fmt"

// Board | Sudoku Board
type Board struct {
	Value [][]int
	Size  int
}

// NewBoard returns a new sudoku board of size n
func NewBoard(n int) *Board {
	b := &Board{
		Value: make([][]int, n),
		Size:  n,
	}
	for i := range b.Value {
		b.Value[i] = make([]int, n)
	}
	return b
}

// IsPossible checks the possibility of writing num at [i,j]
// for that it should not present in ith row and jth column
// and also should not be present in sub-grid
func (b *Board) IsPossible(i, j, num int) bool {
	for x := 0; x < b.Size; x++ {
		if b.Value[x][j] == num || b.Value[i][x] == num {
			return false
		}
	}
	i = (i / 3) * 3
	j = (j / 3) * 3
	for x := i; x < i+3; x++ {
		for y := j; y < j+3; y++ {
			if b.Value[x][y] == num {
				return false
			}
		}
	}
	return true
}

// PrintMatrix prints the matrix
func (b *Board) PrintMatrix() {
	for i, row := range b.Value {
		for j, ele := range row {
			fmt.Printf("%v ", ele)
			if j%3 == 2 {
				fmt.Printf("\t")
			}
		}
		if i%3 == 2 {
			fmt.Printf("\n")
		}
		fmt.Printf("\n")
	}
}

// Solve uses backtracking to solve sudoku
func (b *Board) Solve(i, j int) bool {
	if i == b.Size {
		return true
	}
	if j == b.Size {
		return b.Solve(i+1, 0)
	}
	if b.Value[i][j] != 0 { // Already filled
		return b.Solve(i, j+1)
	}
	for num := 1; num <= b.Size; num++ {
		if b.IsPossible(i, j, num) {
			b.Value[i][j] = num
			rest := b.Solve(i, j+1)
			if rest {
				return rest
			}
		}
	}
	b.Value[i][j] = 0
	return false
}

// Runner | run this in main to check the working
func Runner() {
	board := NewBoard(9)
	board.Value = [][]int{
		{5, 3, 0, 0, 7, 0, 0, 0, 0},
		{6, 0, 0, 1, 9, 5, 0, 0, 0},
		{0, 9, 8, 0, 0, 0, 0, 6, 0},
		{8, 0, 0, 0, 6, 0, 0, 0, 3},
		{4, 0, 0, 8, 0, 3, 0, 0, 1},
		{7, 0, 0, 0, 2, 0, 0, 0, 6},
		{0, 6, 0, 0, 0, 0, 2, 8, 0},
		{0, 0, 0, 4, 1, 9, 0, 0, 5},
		{0, 0, 0, 0, 8, 0, 0, 7, 9},
	}
	board.PrintMatrix()
	fmt.Println("Solution: ", board.Solve(0, 0))
	board.PrintMatrix()
	return
}
