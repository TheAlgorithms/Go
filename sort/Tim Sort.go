package main

import "fmt"

const run = 32

func insertionSort(arr []int, left, right int) {
    for i := left + 1; i <= right; i++ {
        temp := arr[i]
        j := i - 1
        for j >= left && arr[j] > temp {
            arr[j+1] = arr[j]
            j--
        }
        arr[j+1] = temp
    }
}

func mergeTim(arr []int, l, m, r int) {
    left := append([]int(nil), arr[l:m+1]...)
    right := append([]int(nil), arr[m+1:r+1]...)

    i, j, k := 0, 0, l
    for i < len(left) && j < len(right) {
        if left[i] <= right[j] {
            arr[k] = left[i]
            i++
        } else {
            arr[k] = right[j]
            j++
        }
        k++
    }

    for i < len(left) {
        arr[k] = left[i]
        i++
        k++
    }

    for j < len(right) {
        arr[k] = right[j]
        j++
        k++
    }
}

func timSort(arr []int, n int) {
    for i := 0; i < n; i += run {
        insertionSort(arr, i, min((i+run-1), (n-1)))
    }

    size := run
    for size < n {
        for left := 0; left < n; left += 2 * size {
            mid := min(left+size-1, n-1)
            right := min((left + 2*size - 1), (n-1))

            if mid < right {
                mergeTim(arr, left, mid, right)
            }
        }
        size *= 2
    }
}

func min(x, y int) int {
    if x < y {
        return x
    }
    return y
}

func main() {
    arr := []int{34, 2, 78, 12, 45, 21, 67, 1}
    timSort(arr, len(arr))
    fmt.Println("Tim Sorted array:", arr)
}
