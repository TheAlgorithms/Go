package main
import (
  "fmt"
)
// Part of Cosmos by OpenGenus Foundation
func ternarySearch(data []int, left int, right int, value int) int {
  if right >= left {
    mid1 := left + (right-left)/3
    mid2 := right - (right-left)/3
    if data[mid1] == value {
      return mid1
    }
    if data[mid2] == value {
      return mid2
    }
    if value < data[mid1] {
      return ternarySearch(data, left, mid1-1, value)
    } else if value > data[mid2] {
      return ternarySearch(data, mid2+1, right, value)
    } else {
      return ternarySearch(data, mid1+1, mid2-1, value)
    }
  }
  return -1
}
func main() {
  values := []int{1, 2, 3, 4, 5, 12, 30, 35, 46, 84}
  fmt.Println(ternarySearch(values, 0, len(values)-1, 5))
  fmt.Println(ternarySearch(values, 0, len(values)-1, 7))
}
  