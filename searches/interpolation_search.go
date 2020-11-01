package searches

import (
	"fmt"
)

/*
Interpolation search is an improved variant of binary search.
This search algorithm works on the probing position of the required value.
For this algorithm to work properly, the data collection should be in a sorted form and equally distributed.
*/

/*
On average the interpolation search makes about log(log(n)) comparisons
(if the elements are uniformly distributed), where n is the number of elements to be searched.
In the worst case (for instance where the numerical values of the keys increase exponentially)
it can make up to O(n) comparisons.
*/

func interpolationSearch(arr []int, n int, x int) int {
	low := 0
	high := n - 1
	for low <= high && x >= arr[low] && x <= arr[high] {
		if low == high {
			if arr[low] == x {
				return low
			}
			return -1
		}

		pos := low + int((float64((high-low))/float64((arr[high]-arr[low])))*float64(x-arr[low]))

		if arr[pos] == x {
			return pos
		}
		if arr[pos] < x {
			low = pos + 1
		} else {
			high = pos - 1
		}
	}
	return -1
}

func main() {
	var num int = 1000000
	fmt.Println("Interpolation Search:")
	var myArr []int
	for i := 0; i < num; i++ {
		myArr = append(myArr, i)
	}
	x := interpolationSearch(myArr, num, 890)
	if x != -1 {
		fmt.Println("\tElement found at location: ", x)
	} else {
		fmt.Println("\tNot Found!")
	}
}
