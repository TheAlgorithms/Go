package main

import "fmt"

func strassenMatrixMult(A, B [][]int) [][]int {
	n := len(A)
	if n <= 2 {
		return standardMatrixMult(A, B)
	}

	halfN := n / 2

	A11 := make([][]int, halfN)
	A12 := make([][]int, halfN)
	A21 := make([][]int, halfN)
	A22 := make([][]int, halfN)

	B11 := make([][]int, halfN)
	B12 := make([][]int, halfN)
	B21 := make([][]int, halfN)
	B22 := make([][]int, halfN)

	for i := 0; i < halfN; i++ {
		A11[i] = A[i][:halfN]
		A12[i] = A[i][halfN:]
		A21[i] = A[i+halfN][:halfN]
		A22[i] = A[i+halfN][halfN:]

		B11[i] = B[i][:halfN]
		B12[i] = B[i][halfN:]
		B21[i] = B[i+halfN][:halfN]
		B22[i] = B[i+halfN][halfN:]
	}

	P1 := strassenMatrixMult(A11, subtractMatrices(B12, B22))
	P2 := strassenMatrixMult(addMatrices(A11, A12), B22)
	P3 := strassenMatrixMult(addMatrices(A21, A22), B11)
	P4 := strassenMatrixMult(A22, subtractMatrices(B21, B11))
	P5 := strassenMatrixMult(addMatrices(A11, A22), addMatrices(B11, B22))
	P6 := strassenMatrixMult(subtractMatrices(A12, A22), addMatrices(B21, B22))
	P7 := strassenMatrixMult(subtractMatrices(A11, A21), addMatrices(B11, B12))

	C11 := subtractMatrices(addMatrices(addMatrices(P5, P4), P6), P2)
	C12 := addMatrices(P1, P2)
	C21 := addMatrices(P3, P4)
	C22 := subtractMatrices(subtractMatrices(P5, P3), P7)

	C := make([][]int, n)
	for i := 0; i < n; i++ {
		C[i] = make([]int, n)
		copy(C[i][:halfN], C11[i])
		copy(C[i][halfN:], C12[i])
	}
	for i := 0; i < n; i++ {
		copy(C[i+halfN][:halfN], C21[i])
		copy(C[i+halfN][halfN:], C22[i])
	}

	return C
}

func standardMatrixMult(A, B [][]int) [][]int {
	n := len(A)
	C := make([][]int, n)
	for i := 0; i < n; i++ {
		C[i] = make([]int, n)
		for j := 0; j < n; j++ {
			for k := 0; k < n; k++ {
				C[i][j] += A[i][k] * B[k][j]
			}
		}
	}
	return C
}

func addMatrices(A, B [][]int) [][]int {
	n := len(A)
	C := make([][]int, n)
	for i := 0; i < n; i++ {
		C[i] = make([]int, n)
		for j := 0; j < n; j++ {
			C[i][j] = A[i][j] + B[i][j]
		}
	}
	return C
}

func subtractMatrices(A, B [][]int) [][]int {
	n := len(A)
	C := make([][]int, n)
	for i := 0; i < n; i++ {
		C[i] = make([]int, n)
		for j := 0; j < n; j++ {
			C[i][j] = A[i][j] - B[i][j]
		}
	}
	return C
}

func main() {
	A := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	}

	B := [][]int{
		{17, 18, 19, 20},
		{21, 22, 23, 24},
		{25, 26, 27, 28},
		{29, 30, 31, 32},
	}

	result := strassenMatrixMult(A, B)

	for _, row := range result {
		fmt.Println(row)
	}
}
