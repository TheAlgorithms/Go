package tree

import (
	"math/rand"
	"time"
)

func GetRand() *rand.Rand  {
	return rand.New(rand.NewSource(time.Now().UnixNano()))
}

func RandomSlice(start int, end int, count int) []int {
	if end < start || (end-start) < count {
		return nil
	}
	nums := make([]int, 0)
	for len(nums) < count {
		num := GetRand().Intn((end - start)) + start
		exist := false
		for _, v := range nums {
			if v == num {
				exist = true
				break
			}
		}
		if !exist {
			nums = append(nums, num)
		}
	}
	return nums
}