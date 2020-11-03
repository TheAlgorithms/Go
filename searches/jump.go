package main

import (
	"fmt"
	"math"
)

/*
Jump Search is a searching algorithm for sorted arrays.
The basic idea is to check fewer elements (than linear search) by jumping ahead by fixed steps or
skipping some elements in place of searching all elements.
*/

/*
In the worst case, we have to do n/m jumps and if the last checked value
is greater than the element to be searched for, we perform m-1 comparisons more for linear search.
Therefore the total number of comparisons in the worst case will be ((n/m) + m-1).
The value of the function ((n/m) + m-1) will be minimum when m = √n.
Therefore, the best step size is m = √n.
*/

/*
The time complexity of Jump Search is O(√ n)
It is between Linear Search ( ( O(n) ) and Binary Search ( O (Log n) )
*/

func jumpSearch(arr []int, n int, x int) int {
	step := int(math.Sqrt(float64(n)))
	fstep := step
	prev := 0
	for arr[int(math.Min(float64(step), float64(n))-1)] < x {
		prev = step
		step += fstep
		if prev >= n {
			return -1
		}
	}
	m := int(math.Min(float64(step), float64(n)))
	for arr[prev] < x {
		prev++
		if prev == m {
			return -1
		}
	}
	if arr[prev] == x {
		return prev
	}
	return -1
}

func main() {
	fmt.Println("Welcome to jump search!")
	var num int = 1000000
	var myArr []int
	for i := 0; i < num; i++ {
		myArr = append(myArr, i+10000)
	}
	x := jumpSearch(myArr, num, 56789)
	if x != -1 {
		fmt.Println("\tElement found at location: ", x)
	} else {
		fmt.Println("\tElement Not Found!")
	}
}
