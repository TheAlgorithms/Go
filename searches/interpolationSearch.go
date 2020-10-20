package main
import "fmt"
 
func interpolationSearch(array []int, key int) int {
 
    min, max := array[0], array[len(array)-1]
 
    low, high := 0, len(array)-1
 
    for {
        if key < min {
            return low
        }
 
        if key > max {
            return high + 1
        }
 
        // make a guess of the location
        var guess int
        if high == low {
            guess = high
        } else {
            size := high - low
            offset := int(float64(size-1) * (float64(key-min) / float64(max-min)))
            guess = low + offset
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
            high = guess - 1
            max = array[high]
        } else {
            low = guess + 1
            min = array[low]
        }
    }
}
 
 
func main(){
    items := []int{1,2, 9, 20, 31, 45, 63, 70, 100}
    fmt.Println(interpolationSearch(items,63))
}