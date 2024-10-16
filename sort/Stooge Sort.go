package main

import "fmt"

func stoogeSort(arr []int, l, h int) {
    if l >= h {
        return
    }

    if arr[l] > arr[h] {
        arr[l], arr[h] = arr[h], arr[l]
    }

    if h-l+1 > 2 {
        t := (h - l + 1) / 3
        stoogeSort(arr, l, h-t)
        stoogeSort(arr, l+t, h)
        stoogeSort(arr, l, h-t)
    }
}

func main() {
    arr := []int{34, 2, 78, 12, 45, 21, 67, 1}
    stoogeSort(arr, 0, len(arr)-1)
    fmt.Println("Stooge Sorted array:", arr)
}
