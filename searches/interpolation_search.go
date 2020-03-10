
// Interpolation Search in Golang
package main
import "fmt"
 
func interpolationSearch(array []int, key int) int {
 
    min, max := array[0], array[len(array)-1]
 
    start, end := 0, len(array)-1
 
    for {
        if key < min {
            return start
        }
 
        if key > max {
            return end + 1
        }
 
        // make a guess of the location
        var guess int
        if end == start {
            guess = end
        } else {
            size := end - start
            pos := int(float64(size-1) * (float64(key-min) / float64(max-min)))
            guess = start + pos
        }
 
        // maybe we found it?
        if array[guess] == key {
            // scan backwards for start of value range
            for guess > 0 && array[guess-1] == key {
                guess--
            }
            return guess
        }
 
        // if we guessed to high, guess lower or vice versa
        if array[guess] > key {
            end = guess - 1
            max = array[end]
        } else {
            start = guess + 1
            min = array[start]
        }
    }
}


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
  var p int=interpolationSearch(arr,num)
  fmt.Print("Element found at position ",p+1)
}  
