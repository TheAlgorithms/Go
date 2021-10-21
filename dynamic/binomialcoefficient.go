package dynamic

import "github.com/TheAlgorithms/Go/math/min"

// func main() {
// 	myArrayOfK := [4]int{5, 6, 7, 8}
// 	var x int

// 	fmt.Println("\nBinomial Coefficient Using Dynamic Programming:", bin2(50, 5))
// 	for _, element := range myArrayOfK {
// 		start := time.Now()
// 		x = bin2(50, element)
// 		elapsed := time.Since(start)
// 		fmt.Println("bin2 (50,", element, ") = ", x, "    took ", elapsed)

// 	}

// }

// Bin2 function
func Bin2(n int, k int) int {
	var i, j int
	B := make([][]int, (n + 1))
	for i := range B {
		B[i] = make([]int, k+1)
	}

	for i = 0; i <= n; i++ {
		for j = 0; j <= min.Int(i, k); j++ {
			if j == 0 || j == i {
				B[i][j] = 1
			} else {
				B[i][j] = B[i-1][j-1] + B[i-1][j]
			}
		}
	}
	return B[n][k]
}
