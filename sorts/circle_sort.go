package main
 
import "fmt"
 
func circle_sort(arr[] int, start int,end int, swaps int)int{
	s := start
	e := end
	mid := (end - start)/2
	if start == end {
		return swaps
	}

	for(start < end){
		if arr[start] > arr[end] {
			arr[start],arr[end] = arr[end],arr[start]
			swaps = swaps + 1
		}
		start++
		end--
	}
	if start == end {
		if arr[start] > arr[end+1] {
			arr[start],arr[end+1] = arr[end+1],arr[start]
			swaps++
		}
	}
	swaps = circle_sort(arr,s,s+mid,swaps)
	swaps = circle_sort(arr,s+mid+1,e,swaps)
	return swaps
}
 
func main() {
    var n int
    fmt.Print("Enter the number of elements in an Array:")
    fmt.Scan(&n)
    a:=make([]int, n)
    fmt.Println("Enter array elements")
    i:=0
    for(i<n){
      fmt.Print("Enter element ",i+1,":")
      fmt.Scan(&a[i])
      i = i + 1
    }
    for circle_sort(a, 0, n-1, 0) != 0 {
    }
    fmt.Printf("Sorted  : %v\n\n", a)
}
