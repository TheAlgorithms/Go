// Print Pyramid of Numbers
package main
import "fmt"
 
func main() {
    var rows,k,temp,temp1 int
    fmt.Print("Enter number of rows :")
    fmt.Scan(&rows)
 
    for i := 1; i <= rows; i++ {
         
        for j := 1; j <= rows-i; j++ {
            fmt.Print("  ")
            temp++
        }
        for{
            if( temp <= rows-1){
                fmt.Printf(" %d",i+k)
                temp++
            }else{
                temp1++
                fmt.Printf(" %d",(i+k-2*temp1))
            }
            k++
 
            if(k == 2*i-1){             
                break
            }
 
        }
        temp = 0
        temp1 = 0
        k = 0
        fmt.Println("")
    }
 
}
 
// Output
// Enter number of rows : 5
//          1 
//        2 3 2
//      3 4 5 4 3
//    4 5 6 7 6 5 4
//  5 6 7 8 9 8 7 6 5
