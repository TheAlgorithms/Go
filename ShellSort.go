package main


import "fmt"
import "math/rand"
func main() {
	arr := RandomArray(10)
    fmt.Println("Initial array is:", arr)
    fmt.Println("")

    for d := int(len(arr)/2); d > 0; d /= 2 {
    	for i := d; i < len(arr); i++ {
    		for j := i; j >= d && arr[j-d] > arr[j]; j -= d {
    			arr[j], arr[j-d] = arr[j-d], arr[j]
    		}
    	}
    }

    fmt.Println("Sorted array is: ", arr) 	
}
func RandomArray(n int) []int {
    arr := make([]int, n)
    for i := 0; i <= n - 1; i++ {
        arr[i] = rand.Intn(n)
    }
    return arr
}