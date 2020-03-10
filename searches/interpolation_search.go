package main

import ("fmt")

func main(){
  var n,num int
	fmt.Print("Enter the Number of Array: ")
	fmt.Scan(&n)
	fmt.Println("Enter the elements of array")
	arr:=make([]int, n)
	i:=0
	for(i<n){
    fmt.Print("Enter element ",i+1,":")
		fmt.Scan(&arr[i])
		i++
	}
	fmt.Print("Enter number to be searched: ")
    fmt.Scan(&num)
  var p int=interpolation_search(arr,n,num)
  fmt.Print("Element found at position ",p)
}  
func interpolation_search(arr[] int ,n int,num int)int{
    var start,end int
    var pos float64
    start= 0
    end= n-1
    for(start <= end && num >=arr[start] && num <= arr[end]){
      if start == end {
        if arr[start] == num {
          return start
        }
        return -1
      }
      pos = float64(start) + ((float64(end-start)/float64(arr[end]-arr[start])) * float64(num-arr[start]))
      if(arr[int(pos)] == num){
        return int(pos)+1
      } else if(arr[int(pos)] < num){
        start = int(pos) - 1
      } else {
        end = int(pos) - 1
     }
   }
   return -1
}

