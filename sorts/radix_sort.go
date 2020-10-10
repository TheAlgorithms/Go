// Reference: https://medium.com/@houzier.saurav/radix-sort-golang-14cdb995f702

package main
import (
 "fmt"
 "strconv"
 "math"
 "strings"
)
//BigNumPlaceCount to find the biggest number in an array 
//and find the length of this biggest number, for example if 
//the largest element in a target array is 4 digits long 
//this function will return 4
func BigNumPlaceCount(arr []int) int {
 for i := 1; i < len(arr); i++{
  if arr[i] < arr[i-1]{
   rplc := arr[i-1]
   arr[i-1] = arr[i]
   arr[i] = rplc
  }
 }
 biggest := arr[len(arr)-1]
 return len(strconv.Itoa(biggest))
}
func Loop(divisor int, intArr []int, count *[10]int){
for _, value := range intArr {
  rem := (value/divisor)%10
  (*count)[rem] += 1 
 }
 (*count)[0] = (*count)[0] -1
 
 for i := 1; i <len(*count); i ++{
  (*count)[i] = (*count)[i] + (*count)[i-1]
 }
}
//AuxArray generate a new array for a place value divisor
// The new array will be sorted according to the place value in
//numbers
func AuxArray(divisor int, intArr []int, count *[10]int) []int {
 //Start from the end 
 aux := make([]int, len(intArr))
 for i := len(intArr) - 1; i >= 0; i-- {
  //find the target significant digit, if divisor is 10, 
  //find the 10th place digit in the number.
  k := (intArr[i] / divisor) % 10
//find the value corrsponding to this index in the count array
  index := (*count)[k]
//Now in aux array, put this number at the index 
  aux[index] = intArr[i]
//Decrement the count array at the kth index.
  (*count)[k]--
  //fmt.Printf("Count %v -- Aux %v --- IntArr %v\n", *count, intArr, aux)
 }
return aux
}
func RadixSort(intArr []int) []int{
 tmp := make([]int, len(intArr))
 copy(tmp, intArr)
 places := BigNumPlaceCount(tmp)
 
 
 for index, _ := range make([]int, places){
 
  place := int(math.Pow(float64(10), float64(index)))
 
  count := [10]int{}
  
  Loop(place, intArr, &count)
  intArr = AuxArray(place, intArr, &count)
  printString(intArr)
 }
 
 
 return intArr
}
func printString(arr []int){
 z := []string{}
 
 for _, value := range arr{
 
  z = append(z, strconv.Itoa(value))
 }
 fmt.Println(strings.Join(z, " "))
}
func main() {
 a := []int{170, 45, 75, 90, 802, 998, 1234,  24, 2, 66, 121, 223}
 RadixSort(a)
 
}
