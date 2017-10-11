package main

import "fmt"

func main() {
	arr:= [9]int{2,1,4,3,5,9,7,6,8}

	for out:=1; out<len(arr); out++{
		temp:=arr[out]
		in := out

		for;in>0 && arr[in-1]>=temp;in--{
			arr[in] = arr[in-1]
		}
		arr[in] = temp
	}

	for sortedvals:= range arr{
		fmt.Println(sortedvals)
	}
}
