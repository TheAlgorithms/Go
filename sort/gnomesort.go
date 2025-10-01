// gnomesort.go
// description: Implementation of gnome sort algorithm
// worst-case time complexity: O(n^2)
// average-case time complexity: O(n^2)
// best-case time complexity: O(n) (if nearly sorted)
// space complexity: O(1)

package sort

import "github.com/TheAlgorithms/Go/constraints"

// GnomeSort sorts the slice using the gnome sort algorithm
func Gnome[T constraints.Ordered](arr []T) []T {
    i := 1
    n := len(arr)
    for i < n {
        // If at start or current element is in correct order relative to previous
        if i == 0 || arr[i] >= arr[i-1] {
            i++  // move forward
        } else {
            // swap arr[i] and arr[i-1]
            arr[i], arr[i-1] = arr[i-1], arr[i]
            // move back one position
            i--
        }
    }
    return arr
}