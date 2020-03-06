package main

import (
	"container/list"
	"fmt"
	"math"
)

func main() {
	i:=0
	fmt.Print("Enter Number of Elements: ")
  	var n int
  	fmt.Scan(&n)
  	arr:=make([]int, n)
  	for(i<n){
  	fmt.Print("Enter Element ",i+1,": ")
    fmt.Scan(&arr[i])
    i=i+1
  }
	radixSort(arr)
  fmt.Print("Sorted array: ")
	fmt.Print(arr)
}
func radixSort(arr []int) {

	var bucketList [10]*list.List
	for i := 0; i < 10; i++ {
		bucketList[i] = list.New()
	}

	max := max(arr)
	maxdigit := 0
	for max > 0 {
		max /= 10
		maxdigit++
	}
	for i := 0; i < maxdigit; i++ {
		p := int(math.Pow10(i + 1))
		q := int(math.Pow10(i))

		for j := 0; j < len(arr); j++ {
			x := arr[j] % p / q
			bucketList[x].PushBack(arr[j])
		}

		count := 0
		for k := 0; k < 10; k++ {
			for elem := bucketList[k].Front(); elem != nil; elem = elem.Next() {
				arr[count] = elem.Value.(int)
				count++
			}
			bucketList[k].Init()
		}
	}
}

func max(nums []int) int {
	max := nums[0]
	for _, elem := range nums {
		if max < elem {
			max = elem
		}
	}
	return max
}
