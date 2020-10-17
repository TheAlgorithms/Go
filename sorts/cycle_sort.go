package main
 
import (
   "fmt" 
)
 
func cyclesort(ints []int) int {
   writes := 0
 
   for cyclestart := 0; cyclestart < len(ints)-1; cyclestart++ {
      item := ints[cyclestart]
 
      pos := cyclestart
 
      for i := cyclestart + 1; i < len(ints); i++ {
         if ints[i] < item {
            pos++
         }
      }
 
      if pos == cyclestart {
         continue
      }
 
      for item == ints[pos] {
         pos++
      }
 
      ints[pos], item = item, ints[pos]
 
      writes++
 
      for pos != cyclestart {
         pos = cyclestart
         for i := cyclestart + 1; i < len(ints); i++ {
            if ints[i] < item {
               pos++
            }
         }
 
         for item == ints[pos] {
            pos++
         }
 
         ints[pos], item = item, ints[pos]
         writes++
      }
   }
 
   return writes
}
 
func main() {
   list := []int{91,28,73,46,55,64,37,82,19}
   fmt.Println("before:", list)
   writes := cyclesort(list)
   fmt.Println("Writes:",writes)
   fmt.Println("after: ", list)
}
