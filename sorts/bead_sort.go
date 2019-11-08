package main
 
import (
    "fmt"
    "sync"
)
 
var a = []int{170, 45, 75, 90, 802, 24, 2, 66}
var aMax = 1000
 
const bead = 'o'
 
func main() {
    fmt.Println("before:", a)
    beadSort()
    fmt.Println("after: ", a)
}
 
func beadSort() {
    // All space in the abacus = aMax poles x len(a) rows.
    all := make([]byte, aMax*len(a))
    // Slice up space by pole.  (The space could be sliced by row instead,
    // but slicing by pole seemed a more intuitive model of a physical abacus.)
    abacus := make([][]byte, aMax)
    for pole, space := 0, all; pole < aMax; pole++ {
        abacus[pole] = space[:len(a)]
        space = space[len(a):]
    }
    // Use a sync.Waitgroup as the checkpoint mechanism.
    var wg sync.WaitGroup
    // Place beads for each number concurrently. (Presumably beads can be
    // "snapped on" to the middle of a pole without disturbing neighboring
    // beads.)  Also note 'row' here is a row of the abacus.
    wg.Add(len(a))
    for row, n := range a {
        go func(row, n int) {
            for pole := 0; pole < n; pole++ {
                abacus[pole][row] = bead
            }
            wg.Done()
        }(row, n)
    }
    wg.Wait()
    // Now tip the abacus, letting beads fall on each pole concurrently.
    wg.Add(aMax)
    for _, pole := range abacus {
        go func(pole []byte) {
            // Track the top of the stack of beads that have already fallen.
            top := 0
            for row, space := range pole {
                if space == bead {
                    // Move each bead individually, but move it from its
                    // starting row to the top of stack in a single operation.
                    // (More physical simulation such as discovering the top
                    // of stack by inspection, or modeling gravity, are
                    // possible, but didn't seem called for by the task.
                    pole[row] = 0
                    pole[top] = bead
                    top++
                }
            }
            wg.Done()
        }(pole)
    }
    wg.Wait()
    // Read out sorted numbers by row.
    for row := range a {
        x := 0
        for pole := 0; pole < aMax && abacus[pole][row] == bead; pole++ {
            x++
        }
        a[len(a)-1-row] = x
    }
}
