package main

import "fmt"

func bitonicMerge(arr []int, low, cnt int, dir bool) {
    if cnt > 1 {
        k := cnt / 2
        for i := low; i < low+k; i++ {
            if (arr[i] > arr[i+k]) == dir {
                arr[i], arr[i+k] = arr[i+k], arr[i]
            }
        }
        bitonicMerge(arr, low, k, dir)
        bitonicMerge(arr, low+k, k, dir)
    }
}

func bitonicSort(arr []int, low, cnt int, dir bool) {
    if cnt > 1 {
        k := cnt / 2
        bitonicSort(arr, low, k, true)
        bitonicSort(arr, low+k, k, false)
        bitonicMerge(arr, low, cnt, dir)
    }
}

func main() {
    arr := []int{34, 2, 78, 12, 45, 21, 67, 1}
    bitonicSort(arr, 0, len(arr), true)
    fmt.Println("Bitonic Sorted array:", arr)
}
