/*
	Binary search implementation in Go
*/
package main

import "fmt"

func binarySearch(array []int, target int, lowIndex int, highIndex int) int {
	if highIndex < lowIndex {
		return -1
	}
	mid := int((lowIndex + highIndex) / 2)
	if array[mid] > target {
		return binarySearch(array, target, lowIndex, mid)
	} else if array[mid] < target {
		return binarySearch(array, target, mid+1, highIndex)
	} else {
		return mid
	}
}

func iterBinarySearch(array []int, target int, lowIndex int, highIndex int) int {
	startIndex := lowIndex
	endIndex := highIndex
	var mid int
	for startIndex < endIndex {
		mid = int((startIndex + endIndex) / 2)
		if array[mid] > target {
			endIndex = mid
		} else if array[mid] < target {
			startIndex = mid
		} else {
			return mid
		}
	}
	return -1
}
func main(){
	var n,num,ch int
	fmt.Print("Enter number of Elements in an array:")
	fmt.Scan(&n)
	arr:=make([]int, n)
	i:=0
	for(i<n){
		fmt.Print("Enter Element ",i+1,":")
		fmt.Scan(&arr[i])
		i=i+1
	}
	fmt.Print("Enter the element to be searched:")
	fmt.Scan(&num)
	fmt.Println("Enter your choice")
	fmt.Println("1. Iteration")
	fmt.Println("2. Recursion")
	fmt.Scan(&ch)
	switch ch {
	case 1:
		var loc int = iterBinarySearch(arr,num,0,n-1)
		fmt.Print("Element found at position:",loc+1)
	case 2:
		var loc int = binarySearch(arr,num,0,n-1)
		fmt.Print("Element found at position:",loc+1)
	}
}
