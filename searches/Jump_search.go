package main

import ("fmt"
		"math")

func main(){
  i:=0
	fmt.Print("Enter the Number of Array: ")
	var n int
  	fmt.Scan(&n)
	fmt.Println("Enter the elements of array:")
	arr:=make([]int, n)
	for(i<n){
		fmt.Scan(&arr[i])			//element should be entered in sorted order
		i++
	}
  var jump float64
  var num int
	fmt.Print("Enter the Number to be Searched: ")
	fmt.Scan(&num)
	jump=math.Sqrt(float64(n))			//initialising value of Jump by find square_root of number of elements in an array
	i=0
	for(i<n){
		if arr[i] == num {
			fmt.Print("Element found at Location ",i+1)
			break
		} else if arr[i] > num {
			for(arr[i] != num){
				i=i-1
			}
      		  fmt.Print("Element found at Location ",i+1)	
			break
		}
		i=i+int(jump)
	}
	fmt.Print("Element not found")
}
