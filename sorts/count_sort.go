package main

import 	"fmt"
		

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
  	var max int 					//finds the maximum from an array
 	max = arr[0]
 	i=0
 	for(i<n){
 		if arr[i] > max{
 			max=arr[i]

 		}
 		i=i+1
 	}
 	count:=make([]int, max-0+1) 	//initializing count array to zero
 	i=0
 	for(i<max){
 		count[i]=0
 		i++;
 	}

 	i=0
 	for(i<n){
 		count[arr[i]]++
 		i++
 	}
 	i=1
 	for(i<max+1){
 		count[i]+=count[i-1]
 		i=i+1;
 	}
 	out:=make([]int, n)				
 	i=0
 	for(i<n){						//sorting the array
 		out[count[arr[i]]-1] = arr[i]
 		count[arr[i]]--
 		i++
 	}
  fmt.Println("Sorted Array is : ",out)

}
