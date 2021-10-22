package dynamic

import (
	"log"

	"github.com/TheAlgorithms/Go/math/max"
)

// LIS returns the longest increasing subsequence (LIS)
// where all elements of the subsequence are sorted in increasing order
func LIS(elements []int) int {
	n := len(elements)
	lis := make([]int, n)
	for i := range lis {
		lis[i] = 1
	}
	for i := range lis {
		for j := 0; j < i; j++ {
			log.Print(elements[i], elements[j], lis[i], lis[j]+1)
			if elements[i] > elements[j] && lis[i] < lis[j]+1 {
				lis[i] = lis[j] + 1
			}
		}
	}
	log.Println(elements)
	log.Println(lis)
	res := 0
	for i := range lis {
		res = max.Int(res, lis[i])
	}
	return res
}
