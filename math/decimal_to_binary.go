package math
import "fmt"

func decimal_to_binary(num int) {
	var binary []int

   for num !=0 {
      binary = append(binary, num%2)
      num = num / 2
   }
   if len(binary)==0{
      fmt.Printf("%d\n", 0)
   } else {
      for i:=len(binary)-1; i>=0; i--{
         fmt.Printf("%d", binary[i])
      }
      fmt.Println()
   }
}

