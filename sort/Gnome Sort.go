package main

import "fmt"

func gnomeSort(arr []int) {
    i := 0
    for i < len(arr) {
        if i == 0 || arr[i] >= arr[i-1] {
            i++
        } else {
            arr[i], arr[i-1] = arr[i-1], arr[i]
            i--
        }
    }
}

func main() {
    arr := []int{34, 2, 78, 12, 45, 21, 67, 1}
    gnomeSort(arr)
    fmt.Println("Gnome Sorted array:", arr)
}
